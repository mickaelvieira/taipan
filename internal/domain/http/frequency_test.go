package http

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"log"
	"testing"
	"time"
)

func getChecksum(val string) checksum.Checksum {
	return checksum.FromBytes([]byte(val))
}

func getTime(val string) time.Time {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

type entry struct {
	c string
	t string
	s int
}

func getResults(in []entry) (out []*Result) {
	for _, i := range in {
		out = append(out, &Result{
			Checksum:       getChecksum(i.c),
			CreatedAt:      getTime(i.t),
			RespStatusCode: i.s,
		})
	}
	return
}

func TestCalculateIntervalFrequency(t *testing.T) {
	var testcase = []struct {
		f string
		t string
		e Frequency
	}{
		{"2019-06-01T00:00:00Z", "2019-06-01T00:00:01Z", Hourly}, // interval of one second
		{"2019-06-01T00:00:00Z", "2019-06-01T00:01:00Z", Hourly}, // interval of one minute
		{"2019-06-01T00:00:00Z", "2019-06-01T01:00:00Z", Hourly}, // interval of one hour
		{"2019-06-01T00:00:00Z", "2019-06-01T02:00:00Z", Hourly}, // interval of 12 hours
		{"2019-06-01T00:00:00Z", "2019-06-01T23:00:00Z", Hourly}, // interval of 23 hours
		{"2019-06-01T00:00:00Z", "2019-06-02T00:00:00Z", Daily},  // interval of one day
		{"2019-06-01T00:00:00Z", "2019-06-07T00:00:00Z", Daily},  // interval of 6 days
		{"2019-06-01T00:00:00Z", "2019-06-08T00:00:00Z", Weekly}, // interval of one week
		{"2019-06-01T00:00:00Z", "2019-06-15T00:00:00Z", Weekly}, // interval of one fortnight
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test interval frequency [%d]", idx)
		t.Run(name, func(t *testing.T) {
			from := getTime(tc.f)
			to := getTime(tc.t)
			i := inferIntervalFrequency(from, to)
			if i != tc.e {
				t.Errorf("Got [%s], wanted [%s]", i, tc.e)
			}
		})
	}
}

func TestCalculateMode(t *testing.T) {
	var testcase = []struct {
		a []Frequency
		e Frequency
	}{
		{[]Frequency{}, Hourly},
		{[]Frequency{Daily}, Daily},
		{[]Frequency{Weekly, Daily}, Daily},
		{[]Frequency{Daily, Weekly}, Daily},
		{[]Frequency{Daily, Hourly}, Hourly},
		{[]Frequency{Daily, Hourly, Weekly}, Hourly},
		{[]Frequency{Hourly, Daily, Weekly}, Hourly},
		{[]Frequency{Hourly, Daily, Weekly}, Hourly},
		{[]Frequency{Hourly, Hourly, Daily, Weekly}, Hourly},
		{[]Frequency{Daily, Daily, Hourly, Weekly}, Daily},
		{[]Frequency{Hourly, Weekly, Weekly, Daily}, Weekly},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test mode calculation [%d]", idx)
		t.Run(name, func(t *testing.T) {
			i := CalculateMode(tc.a)
			if i != tc.e {
				t.Errorf("Got [%s], wanted [%s]", i, tc.e)
			}
		})
	}
}

func TestCalculateFrequencyDefaultValue(t *testing.T) {
	e := Hourly
	f := CalculateFrequency(make([]*Result, 0))
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}

func TestCalculateFrequencyNotEnoughResults(t *testing.T) {
	e := Hourly
	results := getResults([]entry{
		{"baz", "2019-06-01T00:00:00Z", 404},
		{"bar", "2019-06-02T00:00:00Z", 500},
		{"foo", "2019-06-10T00:00:00Z", 200},
		{"foo", "2019-06-12T00:00:00Z", 400},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}

func TestCalculateFrequencyDaily(t *testing.T) {
	e := Daily
	results := getResults([]entry{
		{"foo", "2019-06-27T00:00:00Z", 200},
		{"bar", "2019-06-26T00:00:00Z", 200},
		{"bar", "2019-06-25T00:00:00Z", 200},
		{"bar", "2019-06-28T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}

func TestCalculateFrequencyWeekly(t *testing.T) {
	e := Weekly
	results := getResults([]entry{
		{"foo", "2019-06-01T00:00:00Z", 200},
		{"bar", "2019-06-08T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}

func TestCalculateFrequencyHasNeverChangedInAWeek(t *testing.T) {
	e := Weekly
	results := getResults([]entry{
		{"bar", "2019-06-25T00:00:00Z", 200},
		{"bar", "2019-06-26T00:00:00Z", 200},
		{"bar", "2019-06-27T00:00:00Z", 200},
		{"bar", "2019-06-28T00:00:00Z", 200},
		{"bar", "2019-06-29T00:00:00Z", 200},
		{"bar", "2019-06-30T00:00:00Z", 200},
		{"bar", "2019-07-01T00:00:00Z", 200},
		{"bar", "2019-07-02T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}
func TestCalculateFrequencyHasRecentlyChanged(t *testing.T) {
	e := Daily
	results := getResults([]entry{
		{"bar", "2019-06-01T00:00:00Z", 200},
		{"bar", "2019-06-08T00:00:00Z", 200},
		{"bar", "2019-06-15T00:00:00Z", 200},
		{"foo", "2019-06-16T00:00:00Z", 200},
		{"baz", "2019-06-17T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%s], wanted [%s]", f, e)
	}
}

func TestFilterSuccessfulResults(t *testing.T) {
	e := 2
	results := getResults([]entry{
		{"foo", "2019-06-26T19:55:22Z", 200},
		{"bar", "2019-06-26T19:55:22Z", 400},
		{"bar", "2019-06-26T19:55:22Z", 200},
		{"bar", "2019-06-26T19:55:22Z", 500},
	})
	r := filterSuccessfulResults(results)
	if len(r) != e {
		t.Errorf("Got [%d], wanted [%d]", len(r), e)
	}
}
