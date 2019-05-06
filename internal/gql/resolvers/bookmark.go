package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/parser"
	"github/mickaelvieira/taipan/internal/s3"
	"log"
)

const defBkmkLimit = 10

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	userBookmarksRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	userBookmark, err := userBookmarksRepo.GetByURL(ctx, user, args.URL)
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}

// GetLatestBookmarks resolves the query
func (r *Resolvers) GetLatestBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*BookmarkCollectionResolver, error) {
	userBookmarksRepo := r.Repositories.UserBookmarks
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
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
	feedsRepo := r.Repositories.Feeds
	bookmarksRepo := r.Repositories.Bookmarks
	userBookmarksRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	document, err := parser.FetchAndParse(args.URL)
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

	res := BookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}

// UpdateBookmark updates a bookmark
func (r *Resolvers) UpdateBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	feedsRepo := r.Repositories.Feeds
	bookmarksRepo := r.Repositories.Bookmarks
	userBookmarksRepo := r.Repositories.UserBookmarks

	user, err := r.getUser(ctx)
	if err != nil {
		return nil, err
	}

	document, err := parser.FetchAndParse(args.URL)
	if err != nil {
		return nil, err
	}

	bookmark := document.ToBookmark()

	if bookmark.Image != nil {
		log.Printf("Uploading %s", bookmark.Image.URL.String())
		image, err := s3.Upload(bookmark.Image.URL.String())
		if err != nil {
			log.Println(err) // @TODO we might eventually better handle this case
		} else {
			log.Printf("Success %s", image.URL.String())
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

	res := BookmarkResolver{UserBookmark: userBookmark}

	return &res, nil
}
