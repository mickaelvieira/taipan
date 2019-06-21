package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	neturl "net/url"
	"os"
)

// BookmarkImageResolver resolves the bookmark's image entity
type BookmarkImageResolver struct {
	*image.Image
}

// URL resolves the URL
func (r *BookmarkImageResolver) URL() scalars.URL {
	var URL = &neturl.URL{}
	URL.Scheme = "https"
	URL.Host = os.Getenv("AWS_BUCKET")
	URL.Path = r.Image.Name

	return scalars.URL{URL: &url.URL{URL: URL}}
}

// Name resolves the Name field
func (r *BookmarkImageResolver) Name() string {
	return r.Image.Name
}

// Width resolves the Width field
func (r *BookmarkImageResolver) Width() int32 {
	return r.Image.Width
}

// Height resolves the Height field
func (r *BookmarkImageResolver) Height() int32 {
	return r.Image.Height
}

// Format resolves the Format field
func (r *BookmarkImageResolver) Format() string {
	return r.Image.Format
}
