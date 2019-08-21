package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"
	"os"
)

// makeImageURL returns an image's URL based on its name
func makeImageURL(name string) *url.URL {
	u, err := url.FromRawURL("https://" + os.Getenv("AWS_BUCKET") + "/" + name)
	if err != nil {
		u = &url.URL{}
	}
	return u
}

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

// BookmarkImage resolves the bookmark's image entity
type BookmarkImage struct {
	*document.Image
}

// URL resolves the URL
func (r *BookmarkImage) URL() scalars.URL {
	if r.Image.Name == "" {
		return scalars.NewURL(r.Image.URL)
	}
	return scalars.NewURL(makeImageURL(r.Image.Name))
}

// Name resolves the Name field
func (r *BookmarkImage) Name() string {
	return r.Image.Name
}

// Width resolves the Width field
func (r *BookmarkImage) Width() int32 {
	return r.Image.Width
}

// Height resolves the Height field
func (r *BookmarkImage) Height() int32 {
	return r.Image.Height
}

// Format resolves the Format field
func (r *BookmarkImage) Format() string {
	return r.Image.Format
}

// UserImage resolves the user's image entity
type UserImage struct {
	*user.Image
}

// URL resolves the URL
func (r *UserImage) URL() scalars.URL {
	return scalars.NewURL(makeImageURL(r.Image.Name))
}

// Name resolves the Name field
func (r *UserImage) Name() string {
	return r.Image.Name
}

// Width resolves the Width field
func (r *UserImage) Width() int32 {
	return r.Image.Width
}

// Height resolves the Height field
func (r *UserImage) Height() int32 {
	return r.Image.Height
}

// Format resolves the Format field
func (r *UserImage) Format() string {
	return r.Image.Format
}
