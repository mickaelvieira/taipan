package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/auth"
)

// FeedsRootResolver feeds' root resolver
type FeedsRootResolver struct {
	repositories *repository.Repositories
}

// Favorites resolves the query
func (r *FeedsRootResolver) Favorites(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*BookmarkCollection, error) {
	page := cursorPagination(10)(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetFavorites(ctx, user, page)
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
		Limit:   page.Limit,
	}

	return &res, nil
}

// ReadingList resolves the query
func (r *FeedsRootResolver) ReadingList(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*BookmarkCollection, error) {
	page := cursorPagination(10)(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetReadingList(ctx, user, page)
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
		Limit:   page.Limit,
	}

	return &res, nil
}

// News resolves the query
func (r *FeedsRootResolver) News(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*DocumentCollection, error) {
	page := cursorPagination(10)(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, page, true)
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
		Limit:   page.Limit,
	}

	return &res, nil
}

// LatestNews resolves the query
func (r *FeedsRootResolver) LatestNews(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*DocumentCollection, error) {
	page := cursorPagination(10)(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, page, false)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Documents.GetTotalLatestNews(ctx, user, page, false)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	res := DocumentCollection{
		Results: resolve(r.repositories).documents(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   page.Limit,
	}

	return &res, nil
}
