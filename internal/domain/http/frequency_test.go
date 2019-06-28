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

func TestcalculateIntervalInSeconds(t *testing.T) {
	var testcase = []struct {
		f string
		t string
		e float64
	}{
		{"2019-06-01T00:00:00Z", "2019-06-01T00:00:01Z", 1},      // one second
		{"2019-06-01T00:00:00Z", "2019-06-01T00:01:00Z", 60},     // one minute
		{"2019-06-01T00:00:00Z", "2019-06-01T01:00:00Z", 3600},   // one hour
		{"2019-06-01T00:00:00Z", "2019-06-02T00:00:00Z", 86400},  // one day
		{"2019-06-01T00:00:00Z", "2019-06-08T00:00:00Z", 604800}, // one week
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test interval in seconds [%d]", idx)
		t.Run(name, func(t *testing.T) {
			from := getTime(tc.f)
			to := getTime(tc.t)
			i := calculateIntervalInSeconds(from, to)
			if i != tc.e {
				t.Errorf("Got [%f], wanted [%f]", i, tc.e)
			}
		})
	}
}

func TestcalculateIntervalFrequency(t *testing.T) {
	var testcase = []struct {
		f string
		t string
		e Frequency
	}{
		{"2019-06-01T00:00:00Z", "2019-06-01T00:00:01Z", Hour}, // interval of one second
		{"2019-06-01T00:00:00Z", "2019-06-01T00:01:00Z", Hour}, // interval of one minute
		{"2019-06-01T00:00:00Z", "2019-06-01T01:00:00Z", Hour}, // interval of one hour
		{"2019-06-01T00:00:00Z", "2019-06-01T02:00:00Z", Hour}, // interval of 12 hours
		{"2019-06-01T00:00:00Z", "2019-06-01T23:00:00Z", Hour}, // interval of 23 hours
		{"2019-06-01T00:00:00Z", "2019-06-02T00:00:00Z", Day},  // interval of one day
		{"2019-06-01T00:00:00Z", "2019-06-07T00:00:00Z", Day},  // interval of 6 days
		{"2019-06-01T00:00:00Z", "2019-06-08T00:00:00Z", Week}, // interval of one week
		{"2019-06-01T00:00:00Z", "2019-06-15T00:00:00Z", Week}, // interval of one fortnight
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test interval frequency [%d]", idx)
		t.Run(name, func(t *testing.T) {
			from := getTime(tc.f)
			to := getTime(tc.t)
			i := calculateIntervalFrequency(from, to)
			if i != tc.e {
				t.Errorf("Got [%d], wanted [%d]", i, tc.e)
			}
		})
	}
}

func TestCalculateMode(t *testing.T) {
	var testcase = []struct {
		a []Frequency
		e Frequency
	}{
		{[]Frequency{}, Hour},
		{[]Frequency{Day}, Day},
		{[]Frequency{Week, Day}, Day},
		{[]Frequency{Day, Week}, Day},
		{[]Frequency{Day, Hour}, Hour},
		{[]Frequency{Day, Hour, Week}, Hour},
		{[]Frequency{Hour, Day, Week}, Hour},
		{[]Frequency{Hour, Day, Week}, Hour},
		{[]Frequency{Hour, Hour, Day, Week}, Hour},
		{[]Frequency{Day, Day, Hour, Week}, Day},
		{[]Frequency{Hour, Week, Week, Day}, Week},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test mode calculation [%d]", idx)
		t.Run(name, func(t *testing.T) {
			i := CalculateMode(tc.a)
			if i != tc.e {
				t.Errorf("Got [%d], wanted [%d]", i, tc.e)
			}
		})
	}
}

func TestCalculateFrequencyDefaultValue(t *testing.T) {
	e := Hour
	f := CalculateFrequency(make([]*Result, 0))
	if f != e {
		t.Errorf("Got [%d], wanted [%d]", f, e)
	}
}

func TestCalculateFrequencyNotEnoughResults(t *testing.T) {
	e := Hour
	results := getResults([]entry{
		{"baz", "2019-06-01T00:00:00Z", 404},
		{"bar", "2019-06-02T00:00:00Z", 500},
		{"foo", "2019-06-10T00:00:00Z", 200},
		{"foo", "2019-06-12T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%d], wanted [%d]", f, e)
	}
}

func TestCalculateFrequencyDayly(t *testing.T) {
	e := Day
	results := getResults([]entry{
		{"foo", "2019-06-26T00:00:00Z", 200},
		{"bar", "2019-06-27T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%d], wanted [%d]", f, e)
	}
}

func TestCalculateFrequencyWeekly(t *testing.T) {
	e := Week
	results := getResults([]entry{
		{"foo", "2019-06-01T00:00:00Z", 200},
		{"bar", "2019-06-08T00:00:00Z", 200},
	})
	f := CalculateFrequency(results)
	if f != e {
		t.Errorf("Got [%d], wanted [%d]", f, e)
	}
}

func TestfilterSuccessfulResults(t *testing.T) {
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
