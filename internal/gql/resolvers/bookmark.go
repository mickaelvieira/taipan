package resolvers

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

var bookmarksLoader = loaders.GetBookmarksLoader()

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.Bookmark
}

//BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	results *[]*BookmarkResolver
}

func (rslv *BookmarkCollectionResolver) Cursor() int32 {
	return 0
}

func (rslv *BookmarkCollectionResolver) Total() int32 {
	return 0
}

func (rslv *BookmarkCollectionResolver) Results() *[]*BookmarkResolver {
	return rslv.results
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

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset int32
	Limit  int32
}) (*BookmarkCollectionResolver, error) {

	var repository = repository.NewBookmarkRepository()
	var ids = repository.FindLatest(ctx, args.Offset, args.Limit)

	log.Println("GET LATEST")
	thunk := bookmarksLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(ids))
	results, err := thunk()

	if err != nil {
		return nil, err[0]
	}

	log.Println(len(results))

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		bookmark, ok := result.(*bookmark.Bookmark)

		if !ok {
			return nil, errors.New("Wrong data")
		}

		res := BookmarkResolver{Bookmark: bookmark}

		bookmarks = append(bookmarks, &res)
	}

	reso := BookmarkCollectionResolver{results: &bookmarks}

	return &reso, nil
}
