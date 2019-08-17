package image

import (
	"bytes"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	img "image"
	"io"
	"io/ioutil"
	"strings"

	// import gif, jpg, png image formats
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// import webp image format
	_ "golang.org/x/image/webp"
)

// Dimensions represents the image's dimensions
type Dimensions struct {
	Width  int
	Height int
}

const (
	gif  = "gif"
	jpg  = "jpg"
	jpeg = "jpeg"
	png  = "png"
	webp = "webp"
)

var mapping = map[string]string{
	gif:  "image/gif",
	jpeg: "image/jpeg",
	png:  "image/png",
	webp: "image/webp",
}

// GetExtension returns the image extension based on the its content type
func GetExtension(c string) (e string) {
	e = strings.TrimPrefix(c, "image/")
	if e == jpg {
		e = jpeg
	}
	if mapping[e] == "" {
		e = ""
	}
	return
}

// GetContentType returns the content type based on the provided extensions
func GetContentType(c string) string {
	if c == jpg {
		c = jpeg
	}
	return mapping[c]
}

// Image interface
type Image interface {
	SetDimensions(w int, h int)
}

// GetName builds image name from its checksum and content type
func GetName(cs checksum.Checksum, ct string) string {
	if f := GetExtension(ct); f != "" {
		return cs.String() + "." + f
	}
	return cs.String()
}

// GetDimensions retrieves image's sizes from reader
// @TODO I need to write a test for that to avoid having the issue with unimported format
func GetDimensions(in io.Reader) (*Dimensions, io.Reader) {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	c, _, err := img.DecodeConfig(bytes.NewReader(buf))
	d := &Dimensions{}

	if err == nil || err != img.ErrFormat {
		d.Width = c.Width
		d.Height = c.Height
	}

	return d, bytes.NewReader(buf)
}
