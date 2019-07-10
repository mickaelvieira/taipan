package image

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"testing"
)

func TestGetExtensionFromContentType(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"foo", ""},
		{"image/jpg", "jpeg"},
		{"image/jpeg", "jpeg"},
		{"image/png", "png"},
		{"image/gif", "gif"},
		{"image/webp", "webp"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Image extension [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var r = GetExtensionFromContentType(tc.i)
			if r != tc.o {
				t.Errorf("Incorrect extension: Wanted [%s]; got [%s]", tc.o, r)
			}
		})
	}
}

func TestGetContentTypeFromExtension(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"foo", ""},
		{"jpg", "image/jpeg"},
		{"jpeg", "image/jpeg"},
		{"png", "image/png"},
		{"gif", "image/gif"},
		{"webp", "image/webp"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Image content type [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var r = GetContentTypeFromExtension(tc.i)
			if r != tc.o {
				t.Errorf("Incorrect content type: Wanted [%s]; got [%s]", tc.o, r)
			}
		})
	}
}

func TestGetName(t *testing.T) {
	var cs = checksum.FromBytes([]byte("foo"))
	var c = cs.String()
	var testcase = []struct {
		i string
		o string
	}{
		{"", c},
		{"foo", c},
		{"image/jpg", c + ".jpeg"},
		{"image/jpeg", c + ".jpeg"},
		{"image/png", c + ".png"},
		{"image/gif", c + ".gif"},
		{"image/webp", c + ".webp"},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Image extension [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var r = GetName(cs, tc.i)
			if r != tc.o {
				t.Errorf("Incorrect extension: Wanted [%s]; got [%s]", tc.o, r)
			}
		})
	}
}
