package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/parser"
	"log"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

const defBkmkLimit = 10

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.UserBookmark // @TODO replace this with UserBookmark eventually
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
	return graphql.ID(rslv.UserBookmark.ID)
}

// URL resolves the URL
func (rslv *BookmarkResolver) URL() string {
	return rslv.UserBookmark.URL
}

// Image resolves the Image field
func (rslv *BookmarkResolver) Image() string {
	return rslv.UserBookmark.Image
}

// Lang resolves the Lang field
func (rslv *BookmarkResolver) Lang() string {
	return rslv.UserBookmark.Lang
}

// Charset resolves the Charset field
func (rslv *BookmarkResolver) Charset() string {
	return rslv.UserBookmark.Charset
}

// Title resolves the Title field
func (rslv *BookmarkResolver) Title() string {
	return rslv.UserBookmark.Title
}

// Description resolves the Description field
func (rslv *BookmarkResolver) Description() string {
	return rslv.UserBookmark.Description
}

// AddedAt resolves the AddedAt field
func (rslv *BookmarkResolver) AddedAt() string {
	return rslv.UserBookmark.AddedAt.Format(time.UnixDate)
}

// UpdatedAt resolves the UpdatedAt field
func (rslv *BookmarkResolver) UpdatedAt() string {
	return rslv.UserBookmark.UpdatedAt.Format(time.UnixDate)
}

// IsRead resolves the IsRead field
func (rslv *BookmarkResolver) IsRead() bool {
	return rslv.UserBookmark.IsRead
}

// // GetBookmark resolves the query
// func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
// 	ID string
// }) (*BookmarkResolver, error) {
// 	var bookmarksLoader = r.Dataloaders.GetBookmarksLoader()
// 	thunk := bookmarksLoader.Load(ctx, dataloader.StringKey(args.ID))
// 	result, err := thunk()

// 	if err != nil {
// 		return nil, err
// 	}

// 	bookmark, ok := result.(*bookmark.Bookmark)

// 	if !ok {
// 		return nil, errors.New("Wrong data")
// 	}

// 	res := BookmarkResolver{Bookmark: bookmark}

// 	return &res, nil
// }

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	ubRepo := r.Repositories.UserBookmarks
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	results, err := ubRepo.FindLatest(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := ubRepo.GetTotal(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks []*BookmarkResolver
	for _, result := range results {
		res := BookmarkResolver{UserBookmark: result}
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

// CreateBookmark creates a new bookmark or updates and existing one
func (r *Resolvers) CreateBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	fRepo := r.Repositories.Feeds
	bRepo := r.Repositories.Bookmarks
	ubRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	bookmark, feeds, err := parser.FetchAndParse(args.URL)
	if err != nil {
		return nil, err
	}

	err = bRepo.Upsert(ctx, bookmark)
	if err != nil {
		return nil, err
	}

	err = fRepo.InsertAllIfNotExists(ctx, feeds)
	if err != nil {
		return nil, err
	}

	err = ubRepo.AddToUserCollection(ctx, user, bookmark)
	if err != nil {
		return nil, err
	}

	userBookmark, err := ubRepo.GetByURL(ctx, user, bookmark.URL)
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}

// UpdateBookmark updates a bookmark
func (r *Resolvers) UpdateBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	fRepo := r.Repositories.Feeds
	bRepo := r.Repositories.Bookmarks
	ubRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	bookmark, feeds, err := parser.FetchAndParse(args.URL)
	if err != nil {
		return nil, err
	}

	b, err := bRepo.GetByURL(ctx, bookmark.URL)
	if err != nil {
		return nil, err
	}

	bookmark.ID = b.ID
	err = bRepo.Update(ctx, bookmark)
	if err != nil {
		return nil, err
	}

	err = fRepo.InsertAllIfNotExists(ctx, feeds)
	if err != nil {
		return nil, err
	}

	err = ubRepo.AddToUserCollection(ctx, user, bookmark)
	if err != nil {
		return nil, err
	}

	userBookmark, err := ubRepo.GetByURL(ctx, user, bookmark.URL)
	if err != nil {
		log.Fatal(err)
	}

	res := BookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}
