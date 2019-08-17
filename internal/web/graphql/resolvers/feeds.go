package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/auth"
)

// FeedsResolver feeds' root resolver
type FeedsResolver struct {
	repositories *repository.Repositories
}

// Favorites resolves the query
func (r *FeedsResolver) Favorites(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*BookmarkCollectionResolver, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetFavorites(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := getBookmarksBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Bookmarks.CountFavorites(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks = make([]*BookmarkResolver, len(results))
	for i, result := range results {
		bookmarks[i] = &BookmarkResolver{
			Bookmark: result,
		}
	}

	res := BookmarkCollectionResolver{
		Results: bookmarks,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// ReadingList resolves the query
func (r *FeedsResolver) ReadingList(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*BookmarkCollectionResolver, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetReadingList(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := getBookmarksBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Bookmarks.CountReadingList(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks = make([]*BookmarkResolver, len(results))
	for i, result := range results {
		bookmarks[i] = &BookmarkResolver{
			Bookmark: result,
		}
	}

	res := BookmarkCollectionResolver{
		Results: bookmarks,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// News resolves the query
func (r *FeedsResolver) News(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollectionResolver, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, from, to, limit, true)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Documents.GetTotalNews(ctx, user)
	if err != nil {
		return nil, err
	}

	var documents = make([]*DocumentResolver, len(results))
	for i, result := range results {
		documents[i] = &DocumentResolver{
			Document:     result,
			repositories: r.repositories,
		}
	}

	res := DocumentCollectionResolver{
		Results: documents,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// LatestNews resolves the query
func (r *FeedsResolver) LatestNews(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollectionResolver, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.GetNews(ctx, user, from, to, limit, false)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Documents.GetTotalLatestNews(ctx, user, from, to, false)
	if err != nil {
		return nil, err
	}

	var documents = make([]*DocumentResolver, len(results))
	for i, result := range results {
		documents[i] = &DocumentResolver{
			Document:     result,
			repositories: r.repositories,
		}
	}

	res := DocumentCollectionResolver{
		Results: documents,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}
