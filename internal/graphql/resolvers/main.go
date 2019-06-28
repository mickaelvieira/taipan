package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	App           *AppResolver
	Users         *UsersResolver
	Documents     *DocumentsResolver
	Bookmarks     *BookmarksResolver
	Syndication   *SyndicationResolver
	Feeds         *FeedsResolver
	subscriptions *Subscription
}

// Favorites subscribes to favorites feed bookmarksEvents
func (r *RootResolver) Favorites(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &BookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(Favorites, s, ctx.Done())
	return c
}

// ReadingList subscribes to reading list feed bookmarksEvents
func (r *RootResolver) ReadingList(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &BookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(ReadingList, s, ctx.Done())
	return c
}

// News subscribes to news feed bookmarksEvents
func (r *RootResolver) News(ctx context.Context) <-chan *DocumentEventResolver {
	c := make(chan *DocumentEventResolver)
	s := &DocumentSubscriber{events: c}
	r.subscriptions.Subscribe(News, s, ctx.Done())
	return c
}

// GetRootResolver returns the root resolver.
// Queries, Mutations and Subscriptions are methods of this resolver
// The root resolver owns a subscription bus to broadcast feed events
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	var subscriptions = &Subscription{
		subscribers: make(map[FeedTopic]Subscribers),
	}
	return &RootResolver{
		App:           &AppResolver{},
		Users:         &UsersResolver{repositories: repositories},
		Documents:     &DocumentsResolver{repositories: repositories},
		Bookmarks:     &BookmarksResolver{repositories: repositories, subscriptions: subscriptions},
		Syndication:   &SyndicationResolver{repositories: repositories},
		Feeds:         &FeedsResolver{repositories: repositories},
		subscriptions: subscriptions,
	}
}
