package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

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

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.Bookmark
}

// ID resolves the ID field
func (r *BookmarkResolver) ID() graphql.ID {
	return graphql.ID(r.Bookmark.ID)
}

// URL resolves the URL
func (r *BookmarkResolver) URL() string {
	return r.Bookmark.URL
}

// Image resolves the Image field
func (r *BookmarkResolver) Image() *BookmarkImageResolver {
	if r.Bookmark.Image == nil {
		return nil
	}

	return &BookmarkImageResolver{
		Image: r.Bookmark.Image,
	}
}

// Lang resolves the Lang field
func (r *BookmarkResolver) Lang() string {
	return r.Bookmark.Lang
}

// Charset resolves the Charset field
func (r *BookmarkResolver) Charset() string {
	return r.Bookmark.Charset
}

// Title resolves the Title field
func (r *BookmarkResolver) Title() string {
	return r.Bookmark.Title
}

// Description resolves the Description field
func (r *BookmarkResolver) Description() string {
	return r.Bookmark.Description
}

// CreatedAt resolves the CreatedAt field
func (r *BookmarkResolver) CreatedAt() string {
	return r.Bookmark.CreatedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *BookmarkResolver) UpdatedAt() string {
	return r.Bookmark.UpdatedAt.Format(time.RFC3339)
}

// UserBookmarkResolver resolves the bookmark entity
type UserBookmarkResolver struct {
	*bookmark.UserBookmark
}

// ID resolves the ID field
func (r *UserBookmarkResolver) ID() graphql.ID {
	return graphql.ID(r.UserBookmark.ID)
}

// URL resolves the URL
func (r *UserBookmarkResolver) URL() string {
	return r.UserBookmark.URL
}

// Image resolves the Image field
func (r *UserBookmarkResolver) Image() *BookmarkImageResolver {
	if r.UserBookmark.Image == nil {
		return nil
	}

	return &BookmarkImageResolver{
		Image: r.UserBookmark.Image,
	}
}

// Lang resolves the Lang field
func (r *UserBookmarkResolver) Lang() string {
	return r.UserBookmark.Lang
}

// Charset resolves the Charset field
func (r *UserBookmarkResolver) Charset() string {
	return r.UserBookmark.Charset
}

// Title resolves the Title field
func (r *UserBookmarkResolver) Title() string {
	return r.UserBookmark.Title
}

// Description resolves the Description field
func (r *UserBookmarkResolver) Description() string {
	return r.UserBookmark.Description
}

// AddedAt resolves the AddedAt field
func (r *UserBookmarkResolver) AddedAt() string {
	return r.UserBookmark.AddedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *UserBookmarkResolver) UpdatedAt() string {
	return r.UserBookmark.UpdatedAt.Format(time.RFC3339)
}

// IsRead resolves the IsRead field
func (r *UserBookmarkResolver) IsRead() bool {
	return r.UserBookmark.IsRead
}
