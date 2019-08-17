package bookmark

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"testing"
)

func TestBookmarkHasImage(t *testing.T) {
	var u, _ = url.FromRawURL("http://foo.example")
	var testcase = []struct {
		i *document.Image
		e bool
	}{
		{nil, false},
		{&document.Image{}, false},
		{&document.Image{Name: "foo"}, false},
		{&document.Image{Name: "foo", URL: u}, true},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test bookmark image [%d]", idx)
		t.Run(name, func(t *testing.T) {
			b := Bookmark{}
			b.Image = tc.i
			r := b.HasImage()
			if r != tc.e {
				t.Errorf("Got [%t], wanted [%t]", r, tc.e)
			}
		})
	}
}
