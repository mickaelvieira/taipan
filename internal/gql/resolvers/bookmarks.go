package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/usecase"
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
	if r.Bookmark.Image == nil || r.Bookmark.Image.Name == "" {
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
	return bool(r.Bookmark.IsRead)
}

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)

	u, err := url.FromRawURL(args.URL)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = r.repositories.Bookmarks.GetByURL(ctx, user, u)
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{Bookmark: b}

	return &res, nil
}

// GetFavorites resolves the query
func (r *Resolvers) GetFavorites(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(10)
	offset, limit := fromArgs(args.Offset, args.Limit)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetFavorites(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Bookmarks.GetTotalFavorites(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		bookmarks = append(bookmarks, &BookmarkResolver{Bookmark: result})
	}

	reso := BookmarkCollectionResolver{
		Results: &bookmarks,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}

// GetReadingList resolves the query
func (r *Resolvers) GetReadingList(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(10)
	offset, limit := fromArgs(args.Offset, args.Limit)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetReadingList(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Bookmarks.GetTotalReadingList(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		bookmarks = append(bookmarks, &BookmarkResolver{Bookmark: result})
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
	url, err := url.FromRawURL(args.URL)
	if err != nil {
		return nil, err
	}

	d, err := usecase.Document(ctx, url, r.repositories)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// ChangeBookmarkReadStatus marks the bookmark as read or unread
func (r *Resolvers) ChangeBookmarkReadStatus(ctx context.Context, args struct {
	URL    string
	IsRead bool
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	isRead := args.IsRead
	url, err := url.FromRawURL(args.URL)
	if err != nil {
		return nil, err
	}

	b, err := usecase.ReadStatus(ctx, user, url, isRead, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Unbookmark removes bookmark from user's list
func (r *Resolvers) Unbookmark(ctx context.Context, args struct {
	URL string
}) (*DocumentResolver, error) {
	user := auth.FromContext(ctx)
	url, err := url.FromRawURL(args.URL)
	if err != nil {
		return nil, err
	}

	d, err := usecase.Unbookmark(ctx, user, url, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &DocumentResolver{Document: d}

	return res, nil
}
