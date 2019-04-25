package resolvers

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

const defBkmkLimit = 10

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

// Image resolves the Image field
func (rslv *BookmarkResolver) Image() string {
	return rslv.Bookmark.Image
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
	var bookmarksLoader = r.Dataloaders.GetBookmarksLoader()
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
	results := repository.FindLatest(ctx, offset, limit)
	total := repository.GetTotal(ctx)

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		res := BookmarkResolver{Bookmark: result}
		bookmarks = append(bookmarks, &res)
	}

	reso := BookmarkCollectionResolver{Results: &bookmarks, Total: total, Offset: offset, Limit: limit}

	return &reso, nil
}

// CreateBookmark creates a new bookmark or updates and existing one
func (r *Resolvers) CreateBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {

	bookmark, err := parser.FetchAndParse(args.URL)

	if err != nil {
		return nil, err
	}

	log.Println(bookmark)

	repository := repository.NewBookmarkRepository()
	ID := repository.GetByURL(ctx, bookmark.URL)

	if ID != "" {
		bookmark.ID = ID
		repository.Update(ctx, bookmark)
	} else {
		repository.Insert(ctx, bookmark)
	}

	linkID, isLinked := repository.IsLinked(ctx, bookmark)

	if linkID == "" {
		repository.Link(ctx, bookmark)
	} else if isLinked == 0 {
		repository.ReLink(ctx, bookmark)
	}

	res := BookmarkResolver{Bookmark: bookmark}

	return &res, nil
}
