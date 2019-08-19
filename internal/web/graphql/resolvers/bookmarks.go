package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/clientid"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"
	"log"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// BookmarksResolver bookmarks' root resolver
type BookmarksResolver struct {
	repositories *repository.Repositories
	publisher    *publisher.Subscription
}

// BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	Results []*BookmarkResolver
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// BookmarkSearchResultsResolver resolver
type BookmarkSearchResultsResolver struct {
	Results []*BookmarkResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// BookmarkResolver resolves the bookmark entity
type BookmarkResolver struct {
	b  *bookmark.Bookmark
	r  *repository.Repositories
	sl *dataloader.Loader
	ll *dataloader.Loader
}

// ID resolves the ID field
func (r *BookmarkResolver) ID() gql.ID {
	return gql.ID(r.b.ID)
}

// URL resolves the URL
func (r *BookmarkResolver) URL() scalars.URL {
	return scalars.NewURL(r.b.URL)
}

// Image resolves the Image field
func (r *BookmarkResolver) Image() *BookmarkImageResolver {
	if !r.b.HasImage() {
		return nil
	}
	return &BookmarkImageResolver{Image: r.b.Image}
}

// Lang resolves the Lang field
func (r *BookmarkResolver) Lang() string {
	return r.b.Lang
}

// Charset resolves the Charset field
func (r *BookmarkResolver) Charset() string {
	return r.b.Charset
}

// Title resolves the Title field
func (r *BookmarkResolver) Title() string {
	return r.b.Title
}

// Description resolves the Description field
func (r *BookmarkResolver) Description() string {
	return r.b.Description
}

// AddedAt resolves the AddedAt field
func (r *BookmarkResolver) AddedAt() scalars.Datetime {
	return scalars.NewDatetime(r.b.AddedAt)
}

// FavoritedAt resolves the FavoritedAt field
func (r *BookmarkResolver) FavoritedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.b.FavoritedAt)
	return &t
}

// UpdatedAt resolves the UpdatedAt field
func (r *BookmarkResolver) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.b.UpdatedAt)
}

// IsFavorite resolves the IsFavorite field
func (r *BookmarkResolver) IsFavorite() bool {
	return bool(r.b.IsFavorite)
}

// Source resolves the Source field
func (r *BookmarkResolver) Source(ctx context.Context) (*SourceResolver, error) {
	if r.b.SourceID == "" {
		return nil, nil
	}

	data, err := r.sl.Load(ctx, dataloader.StringKey(r.b.SourceID))()
	if err != nil {
		return nil, err
	}

	result, ok := data.(*syndication.Source)
	if !ok {
		return nil, errors.New("Loader returns incorrect type")
	}

	return resolve(r.r).source(result), nil
}

// BookmarkEventResolver resolves an bookmark event
type BookmarkEventResolver struct {
	event        *publisher.Event
	repositories *repository.Repositories
}

// Item returns the event's message
func (r *BookmarkEventResolver) Item() *BookmarkResolver {
	b, ok := r.event.Payload.(*bookmark.Bookmark)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a bookmark")
	}

	return resolve(r.repositories).bookmark(b)
}

// Emitter returns the event's emitter ID
func (r *BookmarkEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *BookmarkEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *BookmarkEventResolver) Action() string {
	return string(r.event.Action)
}

type userSubscriber struct {
	events chan<- *UserEventResolver
}

func (s *userSubscriber) Publish(e *publisher.Event) {
	s.events <- &UserEventResolver{event: e}
}

type bookmarkSubscriber struct {
	repositories *repository.Repositories
	events       chan<- *BookmarkEventResolver
}

func (s *bookmarkSubscriber) Publish(e *publisher.Event) {
	s.events <- &BookmarkEventResolver{
		event:        e,
		repositories: s.repositories,
	}
}

type documentSubscriber struct {
	repositories *repository.Repositories
	events       chan<- *DocumentEventResolver
}

func (s *documentSubscriber) Publish(e *publisher.Event) {
	s.events <- &DocumentEventResolver{
		event:        e,
		repositories: s.repositories,
	}
}

// BookmarkChanged --
func (r *RootResolver) BookmarkChanged(ctx context.Context) <-chan *BookmarkEventResolver {
	// @TODO better handle authentication
	c := make(chan *BookmarkEventResolver)
	s := &bookmarkSubscriber{events: c}
	r.publisher.Subscribe(publisher.TopicBookmark, s, ctx.Done())
	return c
}

// Bookmark resolves the query
func (r *BookmarksResolver) Bookmark(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	u := args.URL.ToDomain()

	b, err := r.repositories.Bookmarks.GetByURL(ctx, user, u)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).bookmark(b), nil
}

// Create creates a new document and add it to user's bookmarks
func (r *BookmarksResolver) Create(ctx context.Context, args struct {
	URL        scalars.URL
	IsFavorite bool
	WithFeeds  bool
}) (*BookmarkResolver, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	d, err := usecase.Document(ctx, r.repositories, args.URL.ToDomain(), args.WithFeeds)
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
func (r *BookmarksResolver) Add(ctx context.Context, args struct {
	URL           scalars.URL
	IsFavorite    bool
	Subscriptions *[]scalars.URL
}) (*BookmarkResolver, error) {
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
func (r *BookmarksResolver) Favorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
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
func (r *BookmarksResolver) Unfavorite(ctx context.Context, args struct {
	URL scalars.URL
}) (*BookmarkResolver, error) {
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
func (r *BookmarksResolver) Remove(ctx context.Context, args struct {
	URL scalars.URL
}) (*DocumentResolver, error) {
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
func (r *BookmarksResolver) Search(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     bookmarkSearchInput
}) (*BookmarkSearchResultsResolver, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)
	terms := args.Search.Terms

	var bookmarks []*BookmarkResolver
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

	res := BookmarkSearchResultsResolver{
		Results: bookmarks,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
