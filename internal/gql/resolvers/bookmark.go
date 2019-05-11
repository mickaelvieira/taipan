package resolvers

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"github/mickaelvieira/taipan/internal/usecase"
	"net/url"
)

const defBkmkLimit = 10

// GetBookmark resolves the query
func (r *Resolvers) GetBookmark(ctx context.Context, args struct {
	URL string
}) (*BookmarkResolver, error) {
	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, usecase.ErrUserDoesNotExist
		}
		return nil, err
	}

	var u *url.URL
	u, err = url.ParseRequestURI(args.URL)

	b, err := r.Repositories.Bookmarks.GetByURL(ctx, user, &uri.URI{URL: u})
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{Bookmark: b}

	return &res, nil
}

// GetNewBookmarks resolves the query
func (r *Resolvers) GetNewBookmarks(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*DocumentCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, usecase.ErrUserDoesNotExist
		}
		return nil, err
	}

	results, err := r.Repositories.Documents.FindNew(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.Repositories.Documents.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var documents []*DocumentResolver
	for _, result := range results {
		res := DocumentResolver{Document: result}
		documents = append(documents, &res)
	}

	reso := DocumentCollectionResolver{
		Results: &documents,
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
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(defBkmkLimit)
	offset, limit := fromArgs(args.Offset, args.Limit)

	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, usecase.ErrUserDoesNotExist
		}
		return nil, err
	}

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
	user, err := r.getUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, usecase.ErrUserDoesNotExist
		}
		return nil, err
	}

	var d *document.Document
	d, err = usecase.Document(ctx, args.URL, r.Repositories)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, r.Repositories)

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}
