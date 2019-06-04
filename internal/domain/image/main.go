package image

import "github/mickaelvieira/taipan/internal/domain/url"

// Image represents a bookmark's image
type Image struct {
	Name   string
	URL    *url.URL
	Width  int32
	Height int32
	Format string
}

func (i *Image) String() string {
	if i.URL == nil {
		return ""
	}
	return i.URL.String()
}
