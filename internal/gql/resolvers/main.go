package resolvers

import (
	"github/mickaelvieira/taipan/internal/repository"
)

// RootResolver resolvers
type RootResolver struct {
	repositories          *repository.Repositories
	bookmarksSubscription *Subscription
	documentsSubscription *Subscription
}

// GetRootResolver returns the root resolver. Queries and mutations are methods of this resolver
func GetRootResolver(repositories *repository.Repositories) *RootResolver {
	return &RootResolver{
		repositories:          repositories,
		bookmarksSubscription: &Subscription{subscribers: make(map[Topic]map[string]Subscriber)},
		documentsSubscription: &Subscription{subscribers: make(map[Topic]map[string]Subscriber)},
	}
}
