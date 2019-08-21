package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/auth"
)

// FeedsQuery feeds' root resolver
type FeedsQuery struct {
	repositories *repository.Repositories
}

// Favorites resolves the query
func (r *FeedsQuery) Favorites(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*BookmarkCollection, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetFavorites(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Bookmarks.CountFavorites(ctx, user)
	if err != nil {
		return nil, err
	}

	first, last := getBookmarksBoundaryIDs(results)

	res := BookmarkCollection{
		Results: resolve(r.repositories).bookmarks(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// ReadingList resolves the query
func (r *FeedsQuery) ReadingList(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*BookmarkCollection, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetReadingList(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Bookmarks.CountReadingList(ctx, user)
	if err != nil {
		return nil, err
	}

	first, last := getBookmarksBoundaryIDs(results)

	res := BookmarkCollection{
		Results: resolve(r.repositories).bookmarks(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// News resolves the query
func (r *FeedsQuery) News(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollection, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, from, to, limit, true)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Documents.GetTotalNews(ctx, user)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	res := DocumentCollection{
		Results: resolve(r.repositories).documents(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// LatestNews resolves the query
func (r *FeedsQuery) LatestNews(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollection, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, from, to, limit, false)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Documents.GetTotalLatestNews(ctx, user, from, to, false)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	res := DocumentCollection{
		Results: resolve(r.repositories).documents(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}
