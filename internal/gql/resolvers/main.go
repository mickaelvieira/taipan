package resolvers

import (
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	repositories  *repository.Repositories
	subscriptions *Subscription
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
