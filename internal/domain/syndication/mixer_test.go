package syndication

import (
	"github/mickaelvieira/taipan/internal/domain/url"
	"testing"
)

func getURL(s string) (u *url.URL) {
	u, _ = url.FromRawURL(s)
	return
}

func TestMakeStack(t *testing.T) {
	var u []*url.URL
	u = append(u, getURL("http://example1.local"))
	u = append(u, getURL("http://example2.local"))
	u = append(u, getURL("http://example3.local"))
	u = append(u, getURL("http://example4.local"))
	u = append(u, getURL("http://example5.local"))

	r := MakeMixerStack(u)

	for idx, v := range r {
		var e = u[idx].String()
		if v != e {
			t.Errorf("Wanted [%s], got [%s]", e, v)
		}
	}
}

func TestMixerZeroStack(t *testing.T) {
	i := make([]Stack, 0)

	o := Mixup(i)
	e := 0
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
}

func TestMixerOneStack(t *testing.T) {
	s := Stack{"foo", "bar", "baz"}
	i := Stacks{s}

	o := Mixup(i)

	for idx := range s {
		if o[idx] != s[idx] {
			t.Errorf("Wanted [%s], got [%s]", o[idx], s[idx])
		}
	}
}

func TestMixerTwoEqualStacks(t *testing.T) {
	s1 := Stack{"foo1", "bar1", "baz1"}
	s2 := Stack{"foo2", "bar2", "baz2"}

	i := Stacks{s1, s2}
	o := Mixup(i)

	e := len(s1) + len(s2)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	c1 := 0
	c2 := 0
	for idx, v := range o {
		var w string
		if idx%2 == 0 {
			w = s1[c1]
			c1 = c1 + 1
		} else {
			w = s2[c2]
			c2 = c2 + 1
		}

		if v != w {
			t.Errorf("Wanted [%s], got [%s]", w, v)
		}
	}
}

func TestMixerThreeDifferentStacks(t *testing.T) {
	s1 := Stack{"foo1", "bar1", "baz1"}
	s2 := Stack{"foo2"}
	s3 := Stack{"foo3", "bar3"}

	i := Stacks{s1, s2, s3}
	o := Mixup(i)

	e := len(s1) + len(s2) + len(s3)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	w := Stack{"foo1", "foo2", "foo3", "bar1", "bar3", "baz1"}
	for idx := range o {
		if o[idx] != w[idx] {
			t.Errorf("Wanted [%s], got [%s]", w[idx], o[idx])
		}
	}
}

func TestMixerFourDifferentStacks(t *testing.T) {
	s1 := Stack{"foo1"}
	s2 := Stack{}
	s3 := Stack{"foo3"}
	s4 := Stack{"foo4", "bar4", "baz4"}

	i := Stacks{s1, s2, s3, s4}
	o := Mixup(i)

	e := len(s1) + len(s2) + len(s3) + len(s4)
	r := len(o)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	w := Stack{"foo1", "foo3", "foo4", "bar4", "baz4"}
	for idx := range o {
		if o[idx] != w[idx] {
			t.Errorf("Wanted [%s], got [%s]", w[idx], o[idx])
		}
	}
}

func TestStacksCount(t *testing.T) {
	s1 := Stack{"foo1"}
	s2 := Stack{}
	s3 := Stack{"foo3"}
	s4 := Stack{"foo4", "bar4", "baz4"}

	s := Stacks{s1, s2, s3, s4}

	e := len(s1) + len(s2) + len(s3) + len(s4)
	r := s.count()

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
}

func TestStacksCountAfterMixing(t *testing.T) {
	s1 := Stack{"foo1"}
	s2 := Stack{}
	s3 := Stack{"foo3"}
	s4 := Stack{"foo4", "bar4", "baz4"}

	s := Stacks{s1, s2, s3, s4}
	e := len(s1) + len(s2) + len(s3) + len(s4)

	b := s.count()
	o := Mixup(s)
	a := s.count()

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
