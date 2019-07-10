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

const gif = "gif"
const jpg = "jpg"
const jpeg = "jpeg"
const png = "png"
const webp = "webp"

var mapping = map[string]string{
	gif:  "image/gif",
	jpeg: "image/jpeg",
	png:  "image/png",
	webp: "image/webp",
}

// GetExtensionFromContentType returns the extensions based on the provided content type
func GetExtensionFromContentType(i string) (o string) {
	o = strings.TrimPrefix(i, "image/")
	if o == jpg {
		o = jpeg
	}
	if mapping[o] == "" {
		o = ""
	}
	return
}

// GetContentTypeFromExtension returns the content type based on the provided extensions
func GetContentTypeFromExtension(i string) string {
	if i == jpg {
		i = jpeg
	}
	return mapping[i]
}

// Image interface
type Image interface {
	SetDimensions(w int, h int)
}

// GetName builds image name from its checksum and content type
func GetName(cs checksum.Checksum, ct string) string {
	f := GetExtensionFromContentType(ct)
	if f != "" {
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
