package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	neturl "net/url"
	"os"
)

// ImageResolver interface
type ImageResolver interface {
	// URL resolves the URL
	URL() scalars.URL

	// Name resolves the Name field
	Name() string

	// Width resolves the Width field
	Width() int32

	// Height resolves the Height field
	Height() int32
}

// BookmarkImageResolver resolves the bookmark's image entity
type BookmarkImageResolver struct {
	*document.Image
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

// UserImageResolver resolves the user's image entity
type UserImageResolver struct {
	*user.Image
}

// URL resolves the URL
func (r *UserImageResolver) URL() scalars.URL {
	var URL = &neturl.URL{}
	URL.Scheme = "https"
	URL.Host = os.Getenv("AWS_BUCKET")
	URL.Path = r.Image.Name

	return scalars.URL{URL: &url.URL{URL: URL}}
}

// Name resolves the Name field
func (r *UserImageResolver) Name() string {
	return r.Image.Name
}

// Width resolves the Width field
func (r *UserImageResolver) Width() int32 {
	return r.Image.Width
}

// Height resolves the Height field
func (r *UserImageResolver) Height() int32 {
	return r.Image.Height
}

// Format resolves the Format field
func (r *UserImageResolver) Format() string {
	return r.Image.Format
}
