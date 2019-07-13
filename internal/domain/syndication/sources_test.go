package syndication

import (
	"testing"
)

func TestShiftOnQueueSource(t *testing.T) {
	s1 := &Source{}
	s2 := &Source{}
	var a = []*Source{s1, s2}

	s := QueueSources(a)
	e := 2
	r := len(s)

	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}

	e = 1
	ss1 := s.Shift()
	r = len(s)
	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
	if ss1 != s1 {
		t.Errorf("Wanted [%v], got [%v]", s1, ss1)
	}

	e = 0
	ss2 := s.Shift()
	r = len(s)
	if r != e {
		t.Errorf("Wanted [%d], got [%d]", e, r)
	}
	if ss2 != s2 {
		t.Errorf("Wanted [%v], got [%v]", s2, ss2)
	}

	ss3 := s.Shift()
	if ss3 != nil {
		t.Errorf("Wanted nil, got [%v]", ss3)
	}
}
