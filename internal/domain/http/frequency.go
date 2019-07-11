package http

import (
	"fmt"
	"log"
	"net/http"
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
	Hour = time.Hour
	Day  = time.Hour * 24
	Week = time.Hour * 24 * 7
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

// @TODO I need to test this properly
func filterDuplicateUnchanged(in []*Result) (out []*Result) {
	if len(in) <= 2 {
		return in
	}

	p := in[0]
	// Preserve the first one
	out = append(out, p)

	// remove duplicate between the first and last one
	for i := 1; i < len(in)-1; i++ {
		n := in[i]
		if n.IsContentDifferent(p) {
			out = append(out, n)
			p = n
		}
	}

	// Preserve the last one
	out = append(out, in[len(in)-1])

	return
}

func filterSuccessfulResults(in []*Result) (out []*Result) {
	for _, r := range in {
		if r.RespStatusCode == http.StatusOK {
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
	// Let's make sure we work only with successful results
	in = filterSuccessfulResults(in)

	// Not enough results
	if len(in) < 2 {
		return Hourly
	}

	// Sort http results by ascending creation date
	sort.Sort(ByCreatedAt(in))

	// @TODO we want to be smarter here:
	// - What if a resource has many successful results but has never changed? For instance:
	//   - before filterDuplicateUnchanged: len(in) = 10
	//   - after filterDuplicateUnchanged: len(in) = 1
	// - What if a resource has only a few successful results:
	//   - do we want to calculate an early frequency?
	//   - or we do want to keep fetching in order to have a relevant set of results?
	in = filterDuplicateUnchanged(in)

	out := make([]Frequency, 0)
	for i := 1; i < len(in); i++ {
		p := in[i-1]
		n := in[i]
		if p.CreatedAt.After(n.CreatedAt) {
			panic(fmt.Sprintf("Previous date [%s] cannot be after next date [%s]", p.CreatedAt, n.CreatedAt))
		}
		f := inferIntervalFrequency(
			p.CreatedAt,
			n.CreatedAt,
		)
		out = append(out, f)
	}

	if len(out) < 1 {
		panic("Logic error: We should have at least one frequency")
	}

	log.Printf("%v", out)

	return out[len(out)-1]
}
