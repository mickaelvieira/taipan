package resolvers

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

var bookmarksLoader = loaders.GetBookmarksLoader()

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.Bookmark
}

// ID resolves the ID field
func (rslv *BookmarkResolver) ID() graphql.ID {
	return graphql.ID(rslv.Bookmark.ID)
}

// URL resolves the URL
func (rslv *BookmarkResolver) URL() string {
	return rslv.Bookmark.URL
}

// Hash resolves the Hash field
func (rslv *BookmarkResolver) Hash() string {
	return rslv.Bookmark.Hash
}

// Title resolves the Title field
func (rslv *BookmarkResolver) Title() string {
	return rslv.Bookmark.Title
}

// Description resolves the Description field
func (rslv *BookmarkResolver) Description() string {
	return rslv.Bookmark.Description
}

// Status resolves the Status field
func (rslv *BookmarkResolver) Status() bookmark.Status {
	return rslv.Bookmark.Status
}

// CreatedAt resolves the CreatedAt field
func (rslv *BookmarkResolver) CreatedAt() string {
	return rslv.Bookmark.CreatedAt.Format(time.UnixDate)
}

// UpdatedAt resolves the UpdatedAt field
func (rslv *BookmarkResolver) UpdatedAt() string {
	return rslv.Bookmark.UpdatedAt.Format(time.UnixDate)
}

// AddedAt resolves the AddedAt field
func (rslv *BookmarkResolver) AddedAt() string {
	return rslv.Bookmark.AddedAt.Format(time.UnixDate)
}

// IsRead resolves the IsRead field
func (rslv *BookmarkResolver) IsRead() bool {
	return rslv.Bookmark.IsRead
}

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct{ ID string }) (*BookmarkResolver, error) {
	thunk := bookmarksLoader.Load(ctx, dataloader.StringKey(args.ID))
	result, err := thunk()

	if err != nil {
		return nil, err
	}

	bookmark, ok := result.(*bookmark.Bookmark)

	if !ok {
		return nil, errors.New("Wrong data")
	}

	res := BookmarkResolver{Bookmark: bookmark}

	return &res, nil
}
