package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"github/mickaelvieira/taipan/internal/usecase"
	"net/url"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	Results *[]*BookmarkResolver
	Total   int32
	Offset  int32
	Limit   int32
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
	return r.Bookmark.URL.String()
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

// AddedAt resolves the AddedAt field
func (r *BookmarkResolver) AddedAt() string {
	return r.Bookmark.AddedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *BookmarkResolver) UpdatedAt() string {
	return r.Bookmark.UpdatedAt.Format(time.RFC3339)
}

// IsRead resolves the IsRead field
func (r *BookmarkResolver) IsRead() bool {
	return r.Bookmark.IsRead
}

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)

	u, err := url.ParseRequestURI(args.URL)
	b, err := r.Repositories.Bookmarks.GetByURL(ctx, user, &uri.URI{URL: u})
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{Bookmark: b}

	return &res, nil
}

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(10)
	offset, limit := fromArgs(args.Offset, args.Limit)
	user := auth.FromContext(ctx)

	results, err := r.Repositories.Bookmarks.FindLatest(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.Repositories.Bookmarks.GetTotal(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		res := BookmarkResolver{Bookmark: result}
		bookmarks = append(bookmarks, &res)
	}

	reso := BookmarkCollectionResolver{
		Results: &bookmarks,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}

// Bookmark bookmarks a URL
func (r *Resolvers) Bookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)

	d, err := usecase.Document(ctx, args.URL, r.Repositories)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, r.Repositories)

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}
