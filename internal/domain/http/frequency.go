package http

import (
	"math"
	"sort"
	"time"
)

// List of frequencies a string
const (
	Hourly string = "hourly"
	Daily  string = "daily"
	Weekly string = "weekly"
)

// Frequency of change
type Frequency int

// Possible values of frequencies
const (
	Hour Frequency = 3600
	Day  Frequency = 86400
	Week Frequency = 604800
)

var m = map[Frequency]string{
	Hour: Hourly,
	Day:  Daily,
	Week: Weekly,
}

func (f Frequency) String() string {
	return m[f]
}

func calculateIntervalInSeconds(s time.Time, e time.Time) float64 {
	elapsed := e.Sub(s)
	return elapsed.Seconds()
}

func calculateIntervalFrequency(s time.Time, e time.Time) Frequency {
	i := calculateIntervalInSeconds(s, e)
	r := int(math.Round(i))
	if r >= int(Week) {
		return Week
	}
	if r >= int(Day) {
		return Day
	}
	return Hour
}

// @TODO I need to test this, especially the order
func filterDuplicateUnchanged(in []*Result) (out []*Result) {
	var p *Result
	for _, r := range in {
		if r.IsContentDifferent(p) {
			out = append(out, r)
			p = r
		}
	}
	return
}

func filterSuccessfulResults(in []*Result) (out []*Result) {
	for _, r := range in {
		if r.RespStatusCode == 200 {
			out = append(out, r)
		}
	}
	return
}

// CalculateMode calculates the mode of a list of frequencies
// if no mode is found, it returns the lowest fequency of the set
func CalculateMode(in []Frequency) Frequency {
	// Returns the default frequency
	if len(in) == 0 {
		return Hour
	}

	sort.Slice(in, func(i, j int) bool {
		return int(in[i]) < int(in[j])
	})

	// Deduplication frequencies
	var dedup = make([]Frequency, 0)
	var prev Frequency
	for i := 0; i < len(in); i++ {
		if prev == in[i] {
			continue
		}
		prev = in[i]
		dedup = append(dedup, in[i])
	}

	// Count frequencies
	totals := make(map[Frequency]int, len(dedup))
	for _, f := range in {
		totals[f] = totals[f] + 1
	}

	// Find mode
	var mode = dedup[0]
	var max = 0
	for _, f := range dedup {
		c := totals[f]
		if c > max {
			mode = f
			max = c
		}
	}
	return mode
}

// CalculateFrequency calculates how frequently the resource has changed
func CalculateFrequency(in []*Result) Frequency {
	in = filterSuccessfulResults(in)

	// @TODO I need to double check the sorting
	sort.Sort(ByCreationDate(in))

	// @TODO we want to be smarter here:
	// - What if a resource has many successful results but has never changed? For instance:
	//   - before filterDuplicateUnchanged: len(in) = 10
	//   - after filterDuplicateUnchanged: len(in) = 1
	// - What if a resource has only a few successful results:
	//   - do we want to calculate an early frequency?
	//   - or we do want to keep fetching in order to have a relevant set of results?
	in = filterDuplicateUnchanged(in)

	out := make([]Frequency, 0)
	if len(in) >= 2 {
		for i := 1; i < len(in); i++ {
			p := in[i-1]
			c := in[i]
			f := calculateIntervalFrequency(p.CreatedAt, c.CreatedAt)
			out = append(out, f)
		}
	}

	return CalculateMode(out)
}
