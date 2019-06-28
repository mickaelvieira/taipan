package http

import (
	"fmt"
	"sort"
	"time"
)

// Frequency of change
type Frequency string

// List of frequencies a string
const (
	Hourly Frequency = "hourly"
	Daily  Frequency = "daily"
	Weekly Frequency = "weekly"
)

// Possible values of frequencies
const (
	Hour time.Duration = time.Hour
	Day                = time.Hour * 24
	Week               = time.Hour * 24 * 7
)

var timeIntervals = map[Frequency]time.Duration{
	Hourly: Hour,
	Daily:  Day,
	Weekly: Week,
}

var sqlIntervals = map[Frequency]string{
	Hourly: "1 HOUR",
	Daily:  "1 DAY",
	Weekly: "1 WEEK",
}

// Duration returns the corresponding time interval
func (f Frequency) Duration() time.Duration {
	return timeIntervals[f]
}

// SQLInterval returns the corresponding SQL interval
func (f Frequency) SQLInterval() string {
	return sqlIntervals[f]
}

func inferIntervalFrequency(p time.Time, n time.Time) Frequency {
	i := n.Sub(p)
	if i >= Week {
		return Weekly
	}
	if i >= Day {
		return Daily
	}
	return Hourly
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
		return Hourly
	}

	sort.Slice(in, func(i, j int) bool {
		return in[i].Duration() < in[j].Duration()
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
			n := in[i]
			if p.CreatedAt.After(n.CreatedAt) {
				panic(fmt.Sprintf("Previous date [%s] cannot be after next date [%s]", p.CreatedAt, n.CreatedAt))
			}
			f := inferIntervalFrequency(p.CreatedAt, n.CreatedAt)
			out = append(out, f)
		}
	}

	return CalculateMode(out)
}
