package image

import (
	"testing"
)

func TestGetBase64ContentType(t *testing.T) {
	i := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPo"
	o := GetContentType(i)
	e := "image/png"

	if o != e {
		t.Errorf("Incorrect content type: Wanted [%s] got [%s]", e, o)
	}
}

func TestGetBase64Data(t *testing.T) {
	i := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPo"
	o := GetBase64Data(i)
	e := "iVBORw0KGgoAAAANSUhEUgAAAPo"

	if o != e {
		t.Errorf("Incorrect data: Wanted [%s] got [%s]", e, o)
	}
}
