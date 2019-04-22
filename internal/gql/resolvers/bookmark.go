package resolvers

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/parser"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/repository"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

const defBkmkLimit = 10

var bookmarksLoader = loaders.GetBookmarksLoader()

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.Bookmark // @TODO replace this with UserBookmark eventually
}

//BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	Results *[]*BookmarkResolver
	Total   int32
	Offset  int32
	Limit   int32
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

// Lang resolves the Lang field
func (rslv *BookmarkResolver) Lang() string {
	return rslv.Bookmark.Lang
}

// Charset resolves the Charset field
func (rslv *BookmarkResolver) Charset() string {
	return rslv.Bookmark.Charset
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

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	ID string
}) (*BookmarkResolver, error) {
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

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	repository := repository.NewBookmarkRepository()
	ids := repository.FindLatest(ctx, offset, limit)
	total := repository.CountLatest(ctx)

	thunk := bookmarksLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(ids))
	results, err := thunk()

	// @TODO better error handling
	if err != nil {
		return nil, err[0]
	}

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		bookmark, ok := result.(*bookmark.Bookmark)

		if !ok {
			// @TODO better error handling
			return nil, errors.New("Wrong data")
		}

		res := BookmarkResolver{Bookmark: bookmark}

		bookmarks = append(bookmarks, &res)
	}

	reso := BookmarkCollectionResolver{Results: &bookmarks, Total: total, Offset: offset, Limit: limit}

	return &reso, nil
}

func (r *Resolvers) CreateBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {

	bookmark, err := parser.FetchAndParse(args.URL)

	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{Bookmark: bookmark}

	return &res, nil
}
