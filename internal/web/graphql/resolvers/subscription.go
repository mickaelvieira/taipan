package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// SubscriptionRootResolver syndication's root resolver
type SubscriptionRootResolver struct {
	repositories *repository.Repositories
}

// SubscriptionCollection resolver
type SubscriptionCollection struct {
	Results []*Subscription
	Total   int32
	Offset  int32
	Limit   int32
}

// Subscription resolves the bookmark entity
type Subscription struct {
	subscription *subscription.Subscription
	repositories *repository.Repositories
	logLoader    *dataloader.Loader
}

// ID resolves the ID field
func (r *Subscription) ID() gql.ID {
	return gql.ID(r.subscription.ID)
}

// URL resolves the URL field
func (r *Subscription) URL() scalars.URL {
	return scalars.NewURL(r.subscription.URL)
}

// Domain resolves the Domain field
func (r *Subscription) Domain() *scalars.URL {
	d := scalars.NewURL(r.subscription.Domain)
	return &d
}

// Title resolves the Title field
func (r *Subscription) Title() string {
	return r.subscription.Title
}

// Type resolves the Type field
func (r *Subscription) Type() string {
	return string(r.subscription.Type)
}

// IsSubscribed resolves the IsPaused field
func (r *Subscription) IsSubscribed() bool {
	return r.subscription.Subscribed
}

// Frequency resolves the Frequency field
func (r *Subscription) Frequency() string {
	return string(r.subscription.Frequency)
}

// CreatedAt resolves the CreatedAt field
func (r *Subscription) CreatedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.subscription.CreatedAt)
	return &t
}

// UpdatedAt resolves the UpdatedAt field
func (r *Subscription) UpdatedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.subscription.UpdatedAt)
	return &t
}

// Subscription adds a syndication source and subscribes to it
func (r *SubscriptionRootResolver) Subscription(ctx context.Context, args struct {
	URL scalars.URL
}) (*Subscription, error) {
	u := args.URL.ToDomain()
	user := auth.FromContext(ctx)

	_, err := usecase.CreateSyndicationSource(ctx, r.repositories, u)
	if err != nil {
		return nil, err
	}

	s, err := usecase.SubscribeToSource(ctx, r.repositories, user, u)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).subscription(s), nil
}

// Subscribe --
func (r *SubscriptionRootResolver) Subscribe(ctx context.Context, args struct {
	URL scalars.URL
}) (*Subscription, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.SubscribeToSource(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).subscription(s), nil
}

// Unsubscribe --
func (r *SubscriptionRootResolver) Unsubscribe(ctx context.Context, args struct {
	URL scalars.URL
}) (*Subscription, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.UnubscribeFromSource(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).subscription(s), nil
}

// Subscriptions --
func (r *SubscriptionRootResolver) Subscriptions(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     *subscriptionSearchInput
}) (*SubscriptionCollection, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)

	var terms []string
	showDeleted := false
	pausedOnly := false

	if args.Search != nil {
		terms = args.Search.Terms
		if user.ID == "1" {
			showDeleted = args.Search.ShowDeleted
			pausedOnly = args.Search.PausedOnly
		}
	}

	results, err := r.repositories.Subscriptions.FindAll(ctx, user, terms, showDeleted, pausedOnly, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Subscriptions.GetTotal(ctx, user, terms, showDeleted, pausedOnly)
	if err != nil {
		return nil, err
	}

	res := SubscriptionCollection{
		Results: resolve(r.repositories).subscriptions(results),
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
