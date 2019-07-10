package image

import (
	"github/mickaelvieira/taipan/internal/domain/checksum"
	img "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	_ "golang.org/x/image/webp"
)

// Extension image formats available
type Extension struct {
	GIF  string
	JPEG string
	PNG  string
	WEBP string
}

// ContentType content types available
type ContentType struct {
	GIF  string
	JPG  string
	JPEG string
	PNG  string
	WEBP string
}

var contentType = &ContentType{
	GIF:  "image/gif",
	JPG:  "image/jpg",
	JPEG: "image/jpeg",
	PNG:  "image/png",
	WEBP: "image/webp",
}

var extension = &Extension{
	GIF:  "gif",
	JPEG: "jpeg",
	PNG:  "png",
	WEBP: "webp",
}

// GetExtensionFromContentType returns the extension based on the provided content type
func GetExtensionFromContentType(i string) (o string) {
	switch i {
	case contentType.JPG:
		o = extension.JPEG
	case contentType.JPEG:
		o = extension.JPEG
	case contentType.GIF:
		o = extension.GIF
	case contentType.PNG:
		o = extension.PNG
	case contentType.WEBP:
		o = extension.WEBP
	}
	return
}

// GetContentTypeFromExtension returns the content type based on the provided extension
func GetContentTypeFromExtension(i string) (o string) {
	switch i {
	case extension.JPEG:
		o = contentType.JPEG
	case extension.GIF:
		o = contentType.GIF
	case extension.PNG:
		o = contentType.PNG
	case extension.WEBP:
		o = contentType.WEBP
	}
	return
}

// Image interface
type Image interface {
	SetSizes(w int, h int)
}

// GetName builds image name from its checksum and content type
func GetName(cs checksum.Checksum, ct string) string {
	var f string
	if ct != "" {
		f = GetExtensionFromContentType(ct)
	}
	if f != "" {
		return cs.String() + "." + f
	}
	return cs.String()
}

// GetSizes retrieves image's sizes from reader
// @TODO I need to write a test for that to avoid having the issue with unimported format
func GetSizes(i Image, r io.Reader) error {
	c, _, err := img.DecodeConfig(r)
	if err != nil && err != img.ErrFormat {
		return err
	}

	i.SetSizes(c.Width, c.Height)

	return nil
}
