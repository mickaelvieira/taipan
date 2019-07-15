package syndication

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/url"
	"testing"
)

func getURL(s string) *url.URL {
	u, err := url.FromRawURL(s)
	if err != nil {
		panic(err)
	}
	return u
}

func getURLs(s []string) []*url.URL {
	o := make([]*url.URL, len(s))
	for i, u := range s {
		o[i] = getURL(u)
	}
	return o
}

func TestMakeQueue(t *testing.T) {
	u := getURLs([]string{
		"http://example1.local",
		"http://example2.local",
		"http://example3.local",
		"http://example4.local",
		"http://example5.local",
	})

	r := makeQueue(u)

	for idx, v := range r {
		var e = u[idx].String()
		if v != e {
			t.Errorf("Wanted [%s], got [%s]", e, v)
		}
	}
}

func TestShiftOnQueueURLs(t *testing.T) {
	s1 := "http://example1.local"
	s2 := "http://example2.local"
	var a = []string{s1, s2}

	q := queue(a)
	e := 2
	r := len(q)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	e = 1
	ss1 := q.shift()
	r = len(q)
	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
	if ss1 != s1 {
		t.Errorf("Wanted [%s], got [%s]", s1, ss1)
	}

	e = 0
	ss2 := q.shift()
	r = len(q)
	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
	if ss2 != s2 {
		t.Errorf("Wanted [%s], got [%s]", s2, ss2)
	}

	ss3 := q.shift()
	if ss3 != "" {
		t.Errorf("Wanted an empty string, got [%s]", ss3)
	}
}

func TestMixerWithZeroQueue(t *testing.T) {
	m := MakeMixer(0)
	o := m.Mixup()
	e := 0
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
}

func TestMixerWithOneQueue(t *testing.T) {
	u := []string{
		"http://foo1.local",
		"http://bar1.local",
		"http://baz1.local",
	}

	m := MakeMixer(1)
	m.Push(getURLs(u))
	o := m.Mixup()

	for idx := range o {
		if o[idx] != u[idx] {
			t.Errorf("Wanted [%s], got [%s]", o[idx], u[idx])
		}
	}
}

func TestMixerWithTwoEqualQueues(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
		"http://bar1.local",
		"http://baz1.local",
	}
	u2 := []string{
		"http://foo2.local",
		"http://bar2.local",
		"http://baz2.local",
	}

	m := MakeMixer(2)
	m.Push(getURLs(u1))
	m.Push(getURLs(u2))

	o := m.Mixup()

	e := len(u1) + len(u2)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
	w := []string{
		"http://foo1.local",
		"http://foo2.local",
		"http://bar1.local",
		"http://bar2.local",
		"http://baz1.local",
		"http://baz2.local",
	}
	for idx := range o {
		if o[idx] != w[idx] {
			t.Errorf("Wanted [%s], got [%s]", w[idx], o[idx])
		}
	}
}

func TestMixerWithThreeDifferentQueues(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
		"http://bar1.local",
		"http://baz1.local",
	}
	u2 := []string{
		"http://baz2.local",
	}
	u3 := []string{
		"http://foo3.local",
		"http://bar3.local",
	}

	m := MakeMixer(3)
	m.Push(getURLs(u1))
	m.Push(getURLs(u2))
	m.Push(getURLs(u3))

	o := m.Mixup()

	e := len(u1) + len(u2) + len(u3)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	w := []string{
		"http://foo1.local",
		"http://baz2.local",
		"http://foo3.local",
		"http://bar1.local",
		"http://bar3.local",
		"http://baz1.local",
	}
	for idx := range o {
		if o[idx] != w[idx] {
			t.Errorf("Wanted [%s], got [%s]", w[idx], o[idx])
		}
	}
}

func TestMixerWithFourDifferentQueues(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
	}
	u2 := []string{}
	u3 := []string{
		"http://foo3.local",
	}
	u4 := []string{
		"http://foo4.local",
		"http://bar4.local",
		"http://baz4.local",
	}

	m := MakeMixer(4)
	m.Push(getURLs(u1))
	m.Push(getURLs(u2))
	m.Push(getURLs(u3))
	m.Push(getURLs(u4))

	o := m.Mixup()

	e := len(u1) + len(u2) + len(u3) + len(u4)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	w := []string{
		"http://foo1.local",
		"http://foo3.local",
		"http://foo4.local",
		"http://bar4.local",
		"http://baz4.local",
	}
	for idx := range o {
		if o[idx] != w[idx] {
			t.Errorf("Wanted [%s], got [%s]", w[idx], o[idx])
		}
	}
}

func TestMixerIsFullWithZeroQueue(t *testing.T) {
	m := MakeMixer(0)
	var r = m.IsFull()
	var e = true
	if r != e {
		t.Errorf("Wanted [%t]; got [%t]", e, r)
	}
}

func TestMixerIsFullWithMultipleQueue(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
	}
	u2 := []string{}
	u3 := []string{
		"http://foo3.local",
	}
	u4 := []string{
		"http://foo4.local",
		"http://bar4.local",
		"http://baz4.local",
	}

	m := MakeMixer(4)
	var testcase = []struct {
		i []*url.URL
		o bool
	}{
		{nil, false},
		{getURLs(u1), false},
		{getURLs(u2), false},
		{getURLs(u3), false},
		{getURLs(u4), true},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Mixer is full [%d]", idx)
		t.Run(name, func(t *testing.T) {
			if tc.i != nil {
				m.Push(tc.i)
			}

			var r = m.IsFull()
			if r != tc.o {
				t.Errorf("Wanted [%t]; got [%t]", tc.o, r)
			}
		})
	}
}

func TestMixerCount(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
	}
	u2 := []string{}
	u3 := []string{
		"http://foo3.local",
	}
	u4 := []string{
		"http://foo4.local",
		"http://bar4.local",
		"http://baz4.local",
	}

	m := MakeMixer(4)
	m.Push(getURLs(u1))
	m.Push(getURLs(u2))
	m.Push(getURLs(u3))
	m.Push(getURLs(u4))

	e := len(u1) + len(u2) + len(u3) + len(u4)
	r := m.Count()

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
}

func TestMixerIsEmptyAfterMixing(t *testing.T) {
	u1 := []string{
		"http://foo1.local",
	}
	u2 := []string{}
	u3 := []string{
		"http://foo3.local",
	}
	u4 := []string{
		"http://foo4.local",
		"http://bar4.local",
		"http://baz4.local",
	}

	m := MakeMixer(4)
	m.Push(getURLs(u1))
	m.Push(getURLs(u2))
	m.Push(getURLs(u3))
	m.Push(getURLs(u4))

	e := len(u1) + len(u2) + len(u3) + len(u4)

	b := m.Count()
	o := m.Mixup()
	a := m.Count()

	if b != e {
		t.Errorf("Wanted [%d], got [%d]", e, b)
	}
	if len(o) != e {
		t.Errorf("Wanted [%d], got [%d]", e, len(o))
	}
	if a != 0 {
		t.Errorf("Wanted [%d], got [%d]", 0, a)
	}
}
