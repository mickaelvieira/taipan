package resolvers

import (
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	App           *AppRootResolver
	Users         *UserRootResolver
	Documents     *DocumentRootResolver
	Bookmarks     *BookmarkRootResolver
	Syndication   *SyndicationRootResolver
	Subscriptions *SubscriptionRootResolver
	Feeds         *FeedsRootResolver
	Bot           *LogRootResolver
	publisher     *publisher.Subscription
}

// GetRootResolver returns the root resolver.
// Queries, Mutations and Subscriptions are methods of this resolver
// The root resolver owns a publisher bus to broadcast feed events
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	var publisher = publisher.NewEventBus()
	return &RootResolver{
		App:           &AppRootResolver{},
		Users:         &UserRootResolver{repositories: repositories, publisher: publisher},
		Documents:     &DocumentRootResolver{repositories: repositories},
		Bookmarks:     &BookmarkRootResolver{repositories: repositories, publisher: publisher},
		Syndication:   &SyndicationRootResolver{repositories: repositories},
		Subscriptions: &SubscriptionRootResolver{repositories: repositories},
		Feeds:         &FeedsRootResolver{repositories: repositories},
		Bot:           &LogRootResolver{repositories: repositories},
		publisher:     publisher,
	}
}
