package user

import (
	"fmt"
	"testing"
)

func TestUserHasImage(t *testing.T) {
	var testcase = []struct {
		i *Image
		e bool
	}{
		{nil, false},
		{&Image{}, true},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test user image [%d]", idx)
		t.Run(name, func(t *testing.T) {
			u := User{}
			u.Image = tc.i
			r := u.HasImage()
			if r != tc.e {
				t.Errorf("Got [%t], wanted [%t]", r, tc.e)
			}
		})
	}
}
