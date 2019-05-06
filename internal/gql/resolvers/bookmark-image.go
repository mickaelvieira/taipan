package resolvers

import "github/mickaelvieira/taipan/internal/domain/bookmark"

// BookmarkImageResolver resolves the bookmark's image entity
type BookmarkImageResolver struct {
	*bookmark.Image
}

// URL resolves the URL
func (r *BookmarkImageResolver) URL() string {
	return r.Image.String()
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
