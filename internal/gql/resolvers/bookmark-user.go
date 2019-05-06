package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.UserBookmark
}

// ID resolves the ID field
func (r *BookmarkResolver) ID() graphql.ID {
	return graphql.ID(r.UserBookmark.ID)
}

// URL resolves the URL
func (r *BookmarkResolver) URL() string {
	return r.UserBookmark.URL
}

// Image resolves the Image field
func (r *BookmarkResolver) Image() *BookmarkImageResolver {
	if r.UserBookmark.Image == nil {
		return nil
	}

	return &BookmarkImageResolver{
		Image: r.UserBookmark.Image,
	}
}

// Lang resolves the Lang field
func (r *BookmarkResolver) Lang() string {
	return r.UserBookmark.Lang
}

// Charset resolves the Charset field
func (r *BookmarkResolver) Charset() string {
	return r.UserBookmark.Charset
}

// Title resolves the Title field
func (r *BookmarkResolver) Title() string {
	return r.UserBookmark.Title
}

// Description resolves the Description field
func (r *BookmarkResolver) Description() string {
	return r.UserBookmark.Description
}

// AddedAt resolves the AddedAt field
func (r *BookmarkResolver) AddedAt() string {
	return r.UserBookmark.AddedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *BookmarkResolver) UpdatedAt() string {
	return r.UserBookmark.UpdatedAt.Format(time.RFC3339)
}

// IsRead resolves the IsRead field
func (r *BookmarkResolver) IsRead() bool {
	return r.UserBookmark.IsRead
}
