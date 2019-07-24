package resolvers

import (
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	App           *AppResolver
	Users         *UsersResolver
	Documents     *DocumentsResolver
	Bookmarks     *BookmarksResolver
	Syndication   *SyndicationResolver
	Subscriptions *SubscriptionsResolver
	Feeds         *FeedsResolver
	publisher     *publisher.Subscription
}

// GetRootResolver returns the root resolver.
// Queries, Mutations and Subscriptions are methods of this resolver
// The root resolver owns a publisher bus to broadcast feed events
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	var publisher = publisher.NewEventBus()
	return &RootResolver{
		App:           &AppResolver{},
		Users:         &UsersResolver{repositories: repositories, publisher: publisher},
		Documents:     &DocumentsResolver{repositories: repositories},
		Bookmarks:     &BookmarksResolver{repositories: repositories, publisher: publisher},
		Syndication:   &SyndicationResolver{repositories: repositories},
		Subscriptions: &SubscriptionsResolver{repositories: repositories},
		Feeds:         &FeedsResolver{repositories: repositories},
		publisher:     publisher,
	}
}
