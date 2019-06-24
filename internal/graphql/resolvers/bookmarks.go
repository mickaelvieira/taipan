package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/usecase"

	gql "github.com/graph-gophers/graphql-go"
)

// BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	Results []*BookmarkResolver
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	*bookmark.Bookmark
}

// ID resolves the ID field
func (r *BookmarkResolver) ID() gql.ID {
	return gql.ID(r.Bookmark.ID)
}

// URL resolves the URL
func (r *BookmarkResolver) URL() scalars.URL {
	return scalars.URL{URL: r.Bookmark.URL}
}

// Image resolves the Image field
func (r *BookmarkResolver) Image() *BookmarkImageResolver {
	if r.Bookmark.Image == nil || r.Bookmark.Image.Name == "" {
		return nil
	}

	return &BookmarkImageResolver{
		Image: r.Bookmark.Image,
	}
}

// Lang resolves the Lang field
func (r *BookmarkResolver) Lang() string {
	return r.Bookmark.Lang
}

// Charset resolves the Charset field
func (r *BookmarkResolver) Charset() string {
	return r.Bookmark.Charset
}

// Title resolves the Title field
func (r *BookmarkResolver) Title() string {
	return r.Bookmark.Title
}

// Description resolves the Description field
func (r *BookmarkResolver) Description() string {
	return r.Bookmark.Description
}

// AddedAt resolves the AddedAt field
func (r *BookmarkResolver) AddedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Bookmark.AddedAt}
}

// UpdatedAt resolves the UpdatedAt field
func (r *BookmarkResolver) UpdatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Bookmark.UpdatedAt}
}

// IsRead resolves the IsRead field
func (r *BookmarkResolver) IsRead() bool {
	return bool(r.Bookmark.IsRead)
}

// GetBookmark resolves the query
func (r *RootResolver) GetBookmark(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	b, err := r.repositories.Bookmarks.GetByURL(ctx, user, u)
	if err != nil {
		return nil, err
	}

	res := BookmarkResolver{Bookmark: b}

	return &res, nil
}

// GetFavorites resolves the query
func (r *RootResolver) GetFavorites(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetFavorites(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := GetBookmarksBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Bookmarks.GetTotalFavorites(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks = make([]*BookmarkResolver, 0)
	for _, result := range results {
		bookmarks = append(bookmarks, &BookmarkResolver{Bookmark: result})
	}

	reso := BookmarkCollectionResolver{
		Results: bookmarks,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &reso, nil
}

// FavoritesFeed subscribes to favorites feed bookmarksEvents
func (r *RootResolver) FavoritesFeed(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &BookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(Favorites, s, ctx.Done())
	return c
}

// ReadingListFeed subscribes to reading list feed bookmarksEvents
func (r *RootResolver) ReadingListFeed(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &BookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(ReadingList, s, ctx.Done())
	return c
}

// GetReadingList resolves the query
func (r *RootResolver) GetReadingList(ctx context.Context, args struct {
	Pagination CursorPaginationInput
}) (*BookmarkCollectionResolver, error) {
	fromArgs := GetCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Bookmarks.GetReadingList(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := GetBookmarksBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Bookmarks.GetTotalReadingList(ctx, user)
	if err != nil {
		return nil, err
	}

	var bookmarks = make([]*BookmarkResolver, 0)
	for _, result := range results {
		bookmarks = append(bookmarks, &BookmarkResolver{Bookmark: result})
	}

	reso := BookmarkCollectionResolver{
		Results: bookmarks,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &reso, nil
}

// CreateBookmark creates a new document and add it to user's bookmarks
func (r *RootResolver) CreateBookmark(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	d, err := usecase.Document(ctx, u, r.repositories)
	if err != nil {
		return nil, err
	}

	var isRead = false
	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, isRead, r.repositories)
	if err != nil {
		return nil, err
	}

	r.subscriptions.Publish(NewFeedEvent(ReadingList, Add, b))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Bookmark bookmarks a URL
func (r *RootResolver) Bookmark(ctx context.Context, args struct {
	URL    scalars.URL
	IsRead bool
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	d, err := r.repositories.Documents.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, args.IsRead, r.repositories)
	if err != nil {
		return nil, err
	}

	var removeFrom = News
	var addTo = ReadingList
	if b.IsRead {
		addTo = Favorites
	}

	r.subscriptions.Publish(NewFeedEvent(addTo, Add, b))
	r.subscriptions.Publish(NewFeedEvent(removeFrom, Remove, d))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// ChangeBookmarkReadStatus marks the bookmark as read or unread
func (r *RootResolver) ChangeBookmarkReadStatus(ctx context.Context, args struct {
	URL    scalars.URL
	IsRead bool
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	b, err := usecase.ReadStatus(ctx, user, u, args.IsRead, r.repositories)
	if err != nil {
		return nil, err
	}

	var removeFrom = ReadingList
	var addTo = Favorites
	if !b.IsRead {
		removeFrom = Favorites
		addTo = ReadingList
	}

	r.subscriptions.Publish(NewFeedEvent(addTo, Add, b))
	r.subscriptions.Publish(NewFeedEvent(removeFrom, Remove, b))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Unbookmark removes bookmark from user's list
func (r *RootResolver) Unbookmark(ctx context.Context, args struct {
	URL scalars.URL
}) (*DocumentResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	d, err := usecase.Unbookmark(ctx, user, u, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &DocumentResolver{Document: d}

	return res, nil
}
