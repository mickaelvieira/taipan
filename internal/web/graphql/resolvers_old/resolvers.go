package resolvers

import (
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/pkg/errors"
)

// GraphQL general errors
var (
	ErrLoadersNotFound    = errors.New("Dataloaders cannot be retrieved from the context")
	ErrDataTypeIsNotValid = errors.New("The dataloader returns an incorrect data type")
)

// RootResolver resolvers
// type RootResolver struct {
// 	App           *AppRootResolver
// 	Users         *UserRootResolver
// 	Documents     *DocumentRootResolver
// 	Bookmarks     *BookmarkRootResolver
// 	Syndication   *SyndicationRootResolver
// 	Subscriptions *SubscriptionRootResolver
// 	Feeds         *FeedsRootResolver
// 	Bot           *LogRootResolver
// 	publisher     *publisher.Subscription
// 	repositories  *repository.Repositories
// }

type Query struct {
	App           *AppQuery
	Users         *UsersQuery
	Documents     *DocumentsQuery
	Bookmarks     *BookmarksQuery
	Syndication   *SyndicationQuery
	Subscriptions *SubscriptionsQuery
	Feeds         *FeedsQuery
	Bot           *BotQuery
}

type Mutation struct {
	Users         *UsersMutation
	Documents     *DocumentsMutation
	Bookmarks     *BookmarksMutation
	Syndication   *SyndicationMutation
	Subscriptions *SubscriptionsMutation
}

// // QueryResolver --
// type QueryResolver struct{}

// // App resolves the App query
// func (r *QueryResolver) App(ctx context.Context) (*gqlgen.AppQuery, error) {
// 	q := &gqlgen.AppQuery{
// 		Info: resolveAppInfo(),
// 	}
// 	return q, nil
// }

// // Users resolves the Users query
// func (r *QueryResolver) Users(ctx context.Context) (*gqlgen.UsersQuery, error) {

// }

// func (r *QueryResolver) Documents(ctx context.Context) (*gqlgen.DocumentsQuery, error)
// func (r *QueryResolver) Bookmarks(ctx context.Context) (*gqlgen.BookmarksQuery, error)
// func (r *QueryResolver) Feeds(ctx context.Context) (*gqlgen.FeedsQuery, error)
// func (r *QueryResolver) Syndication(ctx context.Context) (*gqlgen.SyndicationQuery, error)
// func (r *QueryResolver) Subscriptions(ctx context.Context) (*gqlgen.SubscriptionsQuery, error)
// func (r *QueryResolver) Bot(ctx context.Context) (*gqlgen.BotQuery, error)

// type Resolver struct{}

// func (r *Resolver) Mutation() gqlgen.MutationResolver {
// 	return &mutationResolver{r}
// }
// func (r *Resolver) Query() gqlgen.QueryResolver {
// 	return &queryResolver{r}
// }

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
		repositories:  repositories,
	}
}
