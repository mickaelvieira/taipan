package document

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/url"
	"testing"
)

func TestDocumentHasImage(t *testing.T) {
	var u, _ = url.FromRawURL("http://foo.example")
	var testcase = []struct {
		i *Image
		e bool
	}{
		{nil, false},
		{&Image{}, false},
		{&Image{Name: "foo"}, false},
		{&Image{URL: u}, true},
		{&Image{Name: "foo", URL: u}, true},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test document image [%d]", idx)
		t.Run(name, func(t *testing.T) {
			b := Document{}
			b.Image = tc.i
			r := b.HasImage()
			if r != tc.e {
				t.Errorf("Got [%t], wanted [%t]", r, tc.e)
			}
		})
	}
}

func TestDocumentHasImageWhichWasFetched(t *testing.T) {
	var u, _ = url.FromRawURL("http://foo.example")
	var testcase = []struct {
		i *Image
		e bool
	}{
		{nil, false},
		{&Image{}, false},
		{&Image{Name: "foo"}, false},
		{&Image{Name: "foo", URL: u}, true},
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Test document image [%d]", idx)
		t.Run(name, func(t *testing.T) {
			b := Document{}
			b.Image = tc.i
			r := b.WasImageFetched()
			if r != tc.e {
				t.Errorf("Got [%t], wanted [%t]", r, tc.e)
			}
		})
	}
}
