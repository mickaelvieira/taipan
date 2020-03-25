package repository

import (
	"fmt"
	"testing"
)

func TestGetMultiInsertPlacements(t *testing.T) {
	var testcase = []struct {
		t int
		n int
		e string
	}{
		{0, 0, ""},
		{1, 1, "(?)"},
		{1, 3, "(?, ?, ?)"},
		{3, 3, "(?, ?, ?), (?, ?, ?), (?, ?, ?)"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test multiple inserted values [%d]", idx)
		t.Run(name, func(t *testing.T) {
			r := getMultiInsertPlacements(tc.t, tc.n)
			if r != tc.e {
				t.Errorf("Got [%s], wanted [%s]", r, tc.e)
			}
		})
	}
}
