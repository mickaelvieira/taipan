package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	repositories  *repository.Repositories
	subscriptions *Subscription
}

// Documents resolves documents' queries and mutation
func (r *RootResolver) Documents() *DocumentsResolver {
	return &DocumentsResolver{
		repositories: r.repositories,
	}
}

// Bookmarks resolves bookmarks' queries and mutation
func (r *RootResolver) Bookmarks() *BookmarksResolver {
	return &BookmarksResolver{
		subscriptions: r.subscriptions,
		repositories:  r.repositories,
	}
}

// Syndication resolves syndication's queries and mutations
func (r *RootResolver) Syndication() *SyndicationResolver {
	return &SyndicationResolver{
		repositories: r.repositories,
	}
}

// Feeds resolves feeds' queries
func (r *RootResolver) Feeds() *FeedsResolver {
	return &FeedsResolver{
		repositories: r.repositories,
	}
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

// NewsFeed subscribes to news feed bookmarksEvents
func (r *RootResolver) NewsFeed(ctx context.Context) <-chan *DocumentEventResolver {
	c := make(chan *DocumentEventResolver)
	s := &DocumentSubscriber{events: c}
	r.subscriptions.Subscribe(News, s, ctx.Done())
	return c
}

// GetRootResolver returns the root resolver.
// Queries, Mutations and Subscriptions are methods of this resolver
// The root resolver owns a subscription bus to broadcast feed events
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	return &RootResolver{
		repositories: repositories,
		subscriptions: &Subscription{
			subscribers: make(map[FeedTopic]Subscribers),
		},
	}
}
