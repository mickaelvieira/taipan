package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/clientid"
	"github/mickaelvieira/taipan/internal/web/graphql/loaders"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"
	"log"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// BookmarkRootResolver bookmarks' root resolver
type BookmarkRootResolver struct {
	repositories *repository.Repositories
	publisher    *publisher.Subscription
}

// BookmarkCollection resolver
type BookmarkCollection struct {
	Results []*Bookmark
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// BookmarkSearchResults resolver
type BookmarkSearchResults struct {
	Results []*Bookmark
	Total   int32
	Offset  int32
	Limit   int32
}

// Bookmark resolves the bookmark entity
type Bookmark struct {
	bookmark     *bookmark.Bookmark
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *Bookmark) ID() gql.ID {
	return gql.ID(r.bookmark.ID)
}

// User resolves User field
func (r *Bookmark) User(ctx context.Context) (*User, error) {
	l := loaders.FromContext(ctx)
	if l == nil {
		return nil, ErrLoadersNotFound
	}

	d, err := l.Users.Load(ctx, dataloader.StringKey(r.bookmark.UserID))()
	if err != nil {
		return nil, err
	}

	u, ok := d.(*user.User)
	if !ok {
		return nil, ErrDataTypeIsNotValid
	}

	return resolve(r.repositories).user(u), nil
}

// URL resolves the URL
func (r *Bookmark) URL() scalars.URL {
	return scalars.NewURL(r.bookmark.URL)
}

// Image resolves the Image field
func (r *Bookmark) Image() *BookmarkImage {
	if !r.bookmark.HasImage() {
		return nil
	}
	return &BookmarkImage{Image: r.bookmark.Image}
}

// Lang resolves the Lang field
func (r *Bookmark) Lang() string {
	return r.bookmark.Lang
}

// Charset resolves the Charset field
func (r *Bookmark) Charset() string {
	return r.bookmark.Charset
}

// Title resolves the Title field
func (r *Bookmark) Title() string {
	return r.bookmark.Title
}

// Description resolves the Description field
func (r *Bookmark) Description() string {
	return r.bookmark.Description
}

// AddedAt resolves the AddedAt field
func (r *Bookmark) AddedAt() scalars.Datetime {
	return scalars.NewDatetime(r.bookmark.AddedAt)
}

// FavoritedAt resolves the FavoritedAt field
func (r *Bookmark) FavoritedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.bookmark.FavoritedAt)
	return &t
}

// UpdatedAt resolves the UpdatedAt field
func (r *Bookmark) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.bookmark.UpdatedAt)
}

// IsFavorite resolves the IsFavorite field
func (r *Bookmark) IsFavorite() bool {
	return bool(r.bookmark.IsFavorite)
}

// Source resolves the Source field
func (r *Bookmark) Source(ctx context.Context) (*Source, error) {
	if r.bookmark.SourceID == "" {
		return nil, nil
	}

	l := loaders.FromContext(ctx)
	if l == nil {
		return nil, ErrLoadersNotFound
	}

	d, err := l.Sources.Load(ctx, dataloader.StringKey(r.bookmark.SourceID))()
	if err != nil {
		return nil, err
	}

	result, ok := d.(*syndication.Source)
	if !ok {
		return nil, ErrDataTypeIsNotValid
	}

	return resolve(r.repositories).source(result), nil
}

// BookmarkEvent resolves an bookmark event
type BookmarkEvent struct {
	event        *publisher.Event
	repositories *repository.Repositories
}

// Item returns the event's message
func (r *BookmarkEvent) Item() *Bookmark {
	b, ok := r.event.Payload.(*bookmark.Bookmark)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a bookmark")
	}

	return resolve(r.repositories).bookmark(b)
}

// Emitter returns the event's emitter ID
func (r *BookmarkEvent) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *BookmarkEvent) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *BookmarkEvent) Action() string {
	return string(r.event.Action)
}

type bookmarkSubscriber struct {
	repositories *repository.Repositories
	events       chan<- *BookmarkEvent
}

func (s *bookmarkSubscriber) Publish(e *publisher.Event) {
	s.events <- &BookmarkEvent{
		event:        e,
		repositories: s.repositories,
	}
}

// BookmarkChanged --
func (r *RootResolver) BookmarkChanged(ctx context.Context) <-chan *BookmarkEvent {
	// @TODO better handle authentication
	c := make(chan *BookmarkEvent)
	s := &bookmarkSubscriber{
		events:       c,
		repositories: r.repositories,
	}
	r.publisher.Subscribe(publisher.TopicBookmark, s, ctx.Done())
	return c
}

// Bookmark resolves the query
func (r *BookmarkRootResolver) Bookmark(ctx context.Context, args struct {
	URL scalars.URL
}) (*Bookmark, error) {
	user := auth.FromContext(ctx)
	u := args.URL.ToDomain()

	b, err := r.repositories.Bookmarks.GetByURL(ctx, user, u)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).bookmark(b), nil
}

// Create creates a new document and add it to user's bookmarks
func (r *BookmarkRootResolver) Create(ctx context.Context, args struct {
	URL        scalars.URL
	IsFavorite bool
}) (*Bookmark, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	d, err := usecase.Document(ctx, r.repositories, args.URL.ToDomain(), "")
	if err != nil {
		return nil, err
	}

	b, err := usecase.Bookmark(ctx, r.repositories, user, d, args.IsFavorite)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicBookmark, publisher.Bookmark, b),
	)

	return resolve(r.repositories).bookmark(b), nil
}

// Add bookmarks a URL
func (r *BookmarkRootResolver) Add(ctx context.Context, args struct {
	URL           scalars.URL
	IsFavorite    bool
	Subscriptions *[]scalars.URL
}) (*Bookmark, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	d, err := r.repositories.Documents.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	b, err := usecase.Bookmark(ctx, r.repositories, user, d, args.IsFavorite)
	if err != nil {
		return nil, err
	}

	// subscribes to sources sent along
	if args.Subscriptions != nil {
		subscriptions := *args.Subscriptions
		for _, u := range subscriptions {
			_, err := usecase.SubscribeToSource(ctx, r.repositories, user, u.ToDomain())
			if err != nil {
				return nil, err
			}
		}
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicBookmark, publisher.Bookmark, b),
	)

	return resolve(r.repositories).bookmark(b), nil
}

// Favorite adds the bookmark to favorites
func (r *BookmarkRootResolver) Favorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*Bookmark, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	b, err := usecase.Favorite(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicBookmark, publisher.Favorite, b),
	)

	return resolve(r.repositories).bookmark(b), nil
}

// Unfavorite removes the bookmark from favorites
func (r *BookmarkRootResolver) Unfavorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*Bookmark, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	b, err := usecase.Unfavorite(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicBookmark, publisher.Unfavorite, b),
	)

	return resolve(r.repositories).bookmark(b), nil
}

// Remove removes bookmark from user's list
func (r *BookmarkRootResolver) Remove(ctx context.Context, args struct {
	URL scalars.URL
}) (*Document, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	d, err := usecase.Unbookmark(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicDocument, publisher.Unbookmark, d),
	)

	return resolve(r.repositories).document(d), nil
}

// Search --
func (r *BookmarkRootResolver) Search(ctx context.Context, args struct {
	Pagination OffsetPaginationInput
	Search     BookmarkSearchInput
}) (*BookmarkSearchResults, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)
	terms := args.Search.Terms

	var bookmarks []*Bookmark
	var total int32

	if len(terms) > 0 {
		results, err := r.repositories.Bookmarks.FindAll(ctx, user, terms, offset, limit)
		if err != nil {
			return nil, err
		}

		total, err = r.repositories.Bookmarks.CountAll(ctx, user, terms)
		if err != nil {
			return nil, err
		}

		bookmarks = resolve(r.repositories).bookmarks(results)
	}

	res := BookmarkSearchResults{
		Results: bookmarks,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
