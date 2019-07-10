package image

import (
	"fmt"
	"testing"
)

func TestGetBase64ContentType(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", ""},
		{"foo", ""},
		{"data:image/jpeg", ""},
		{"image/jpeg;base64,bar", "image/jpeg"},
		{"data:image/jpeg,bar", "image/jpeg"},
		{"data:image/png;base64,bar", "image/png"},
		{"data:image/gif;base64,bar", "image/gif"},
		{"data:image/webp;base64,bar", "image/webp"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Image content type [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var r = GetContentType(tc.i)
			if r != tc.o {
				t.Errorf("Incorrect content type: Wanted [%s]; got [%s]", tc.o, r)
			}
		})
	}
}

func TestGetBase64Data(t *testing.T) {
	i := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPo"
	o := GetBase64Data(i)
	e := "iVBORw0KGgoAAAANSUhEUgAAAPo"

	if o != e {
		t.Errorf("Incorrect data: Wanted [%s] got [%s]", e, o)
	}

	var testcase = []struct {
		i string
		o string
	}{
		{"", ""},
		{"foo", ""},
		{"image/jpeg;base64,bar", "bar"},
		{"data:image/jpeg;base64,", ""},
		{"data:image/jpeg,bar", "bar"},
		{"data:image/jpeg;base64,bar", "bar"},
		{"data:image/png,bar", "bar"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Image data [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var r = GetBase64Data(tc.i)
			if r != tc.o {
				t.Errorf("Incorrect data: Wanted [%s]; got [%s]", tc.o, r)
			}
		})
	}
}
