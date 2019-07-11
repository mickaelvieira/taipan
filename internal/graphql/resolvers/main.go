package resolvers

import (
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
	subscriptions *subscription
}

// GetRootResolver returns the root resolver.
// Queries, Mutations and Subscriptions are methods of this resolver
// The root resolver owns a subscription bus to broadcast feed events
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	var subscriptions = &subscription{
		subscribers: make(map[feedTopic]subscribers),
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
