package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"

	gql "github.com/graph-gophers/graphql-go"
)

// BookmarksResolver bookmarks' root resolver
type BookmarksResolver struct {
	repositories  *repository.Repositories
	subscriptions *Subscription
}

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

// IsFavorite resolves the IsFavorite field
func (r *BookmarkResolver) IsFavorite() bool {
	return bool(r.Bookmark.IsFavorite)
}

// Bookmark resolves the query
func (r *BookmarksResolver) Bookmark(ctx context.Context, args struct {
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

// Create creates a new document and add it to user's bookmarks
func (r *BookmarksResolver) Create(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	d, err := usecase.Document(ctx, u, r.repositories)
	if err != nil {
		return nil, err
	}

	var isFavorite = false
	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, isFavorite, r.repositories)
	if err != nil {
		return nil, err
	}

	r.subscriptions.Publish(NewFeedEvent(ReadingList, Add, b))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Add bookmarks a URL
func (r *BookmarksResolver) Add(ctx context.Context, args struct {
	URL        scalars.URL
	IsFavorite bool
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	d, err := r.repositories.Documents.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	var b *bookmark.Bookmark
	b, err = usecase.Bookmark(ctx, user, d, args.IsFavorite, r.repositories)
	if err != nil {
		return nil, err
	}

	var addTo = ReadingList
	if b.IsFavorite {
		addTo = Favorites
	}

	r.subscriptions.Publish(NewFeedEvent(addTo, Add, b))
	r.subscriptions.Publish(NewFeedEvent(News, Remove, d))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Favorite adds the bookmark to favorites
func (r *BookmarksResolver) Favorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	b, err := usecase.FavoriteStatus(ctx, user, u, true, r.repositories)
	if err != nil {
		return nil, err
	}

	r.subscriptions.Publish(NewFeedEvent(Favorites, Add, b))
	r.subscriptions.Publish(NewFeedEvent(ReadingList, Remove, b))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Unfavorite removes the bookmark from favorites
func (r *BookmarksResolver) Unfavorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.URL

	b, err := usecase.FavoriteStatus(ctx, user, u, false, r.repositories)
	if err != nil {
		return nil, err
	}

	r.subscriptions.Publish(NewFeedEvent(ReadingList, Add, b))
	r.subscriptions.Publish(NewFeedEvent(Favorites, Remove, b))

	res := &BookmarkResolver{Bookmark: b}

	return res, nil
}

// Remove removes bookmark from user's list
func (r *BookmarksResolver) Remove(ctx context.Context, args struct {
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
