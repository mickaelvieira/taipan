package resolvers

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/parser"
	"github/mickaelvieira/taipan/internal/s3"
	"log"
)

const defBkmkLimit = 10

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	URL string
}) (*UserBookmarkResolver, error) {
	userBookmarksRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Unknown user")
		}
		return nil, err
	}

	userBookmark, err := userBookmarksRepo.GetByURL(ctx, user, args.URL)
	if err != nil {
		return nil, err
	}

	res := UserBookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}

// GetNewBookmarks resolves the query
func (r *Resolvers) GetNewBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	bookmarksRepo := r.Repositories.Bookmarks
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Unknown user")
		}
		return nil, err
	}

	results, err := bookmarksRepo.FindNew(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := bookmarksRepo.GetTotal(ctx)
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

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*UserBookmarkCollectionResolver, error) {
	userBookmarksRepo := r.Repositories.UserBookmarks
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Unknown user")
		}
		return nil, err
	}

	results, err := userBookmarksRepo.FindLatest(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := userBookmarksRepo.GetTotal(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks []*UserBookmarkResolver
	for _, result := range results {
		res := UserBookmarkResolver{UserBookmark: result}
		bookmarks = append(bookmarks, &res)
	}

	reso := UserBookmarkCollectionResolver{
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
}) (*UserBookmarkResolver, error) {
	feedsRepo := r.Repositories.Feeds
	bookmarksRepo := r.Repositories.Bookmarks
	userBookmarksRepo := r.Repositories.UserBookmarks
	logsRepo := r.Repositories.Botlogs

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Unknown user")
		}
		return nil, err
	}

	document, reqLog, err := parser.FetchAndParse(args.URL)
	if err != nil {
		return nil, err
	}

	bookmark := document.ToBookmark()

	if bookmark.Image != nil {
		image, err := s3.Upload(bookmark.Image.URL.String())
		if err != nil {
			log.Println(err) // @TODO we might eventually better handle this case
		} else {
			bookmark.Image = image
		}
	}

	err = bookmarksRepo.Upsert(ctx, bookmark)
	if err != nil {
		return nil, err
	}

	if bookmark.Image != nil {
		err = bookmarksRepo.UpdateImage(ctx, bookmark)
		if err != nil {
			return nil, err
		}
	}

	err = logsRepo.Insert(ctx, reqLog)
	if err != nil {
		return nil, err
	}

	err = feedsRepo.InsertAllIfNotExists(ctx, document.Feeds)
	if err != nil {
		return nil, err
	}

	err = userBookmarksRepo.AddToUserCollection(ctx, user, bookmark)
	if err != nil {
		return nil, err
	}

	userBookmark, err := userBookmarksRepo.GetByURL(ctx, user, bookmark.URL)
	if err != nil {
		return nil, err
	}

	res := UserBookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}
