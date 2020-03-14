package repository

import (
	"github.com/mickaelvieira/taipan/internal/domain/newsfeed"
	"testing"
	"time"
)

func TestMutlipleEntryParameters(t *testing.T) {
	u1 := "foo1"
	u2 := "foo2"
	u3 := "foo3"

	d1 := "bar1"
	d2 := "bar2"
	d3 := "bar3"

	dt1 := time.Now()
	dt2 := time.Now()
	dt3 := time.Now()

	e1 := &newsfeed.Entry{UserID: u1, DocumentID: d1, CreatedAt: dt1}
	e2 := &newsfeed.Entry{UserID: u2, DocumentID: d2, CreatedAt: dt2}
	e3 := &newsfeed.Entry{UserID: u3, DocumentID: d3, CreatedAt: dt3}

	entries := []*newsfeed.Entry{e1, e2, e3}
	args := getFeedEntriesParameters(entries)

	w := []interface{}{u1, d1, dt1, u2, d2, dt2, u3, d3, dt3}

	for i := range w {
		if args[i] != w[i] {
			t.Errorf("Wanted [%v], got [%v]", w[i], args[i])
		}
	}
}
