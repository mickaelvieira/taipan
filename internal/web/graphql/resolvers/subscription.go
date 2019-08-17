package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

	gql "github.com/graph-gophers/graphql-go"
)

// SubscriptionsResolver syndication's root resolver
type SubscriptionsResolver struct {
	repositories *repository.Repositories
}

// SubscriptionCollectionResolver resolver
type SubscriptionCollectionResolver struct {
	Results []*SubscriptionResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// SubscriptionResolver resolves the bookmark entity
type SubscriptionResolver struct {
	*subscription.Subscription
}

// ID resolves the ID field
func (r *SubscriptionResolver) ID() gql.ID {
	return gql.ID(r.Subscription.ID)
}

// URL resolves the URL field
func (r *SubscriptionResolver) URL() scalars.URL {
	return scalars.NewURL(r.Subscription.URL)
}

// Domain resolves the Domain field
func (r *SubscriptionResolver) Domain() *scalars.URL {
	d := scalars.NewURL(r.Subscription.Domain)
	return &d
}

// Title resolves the Title field
func (r *SubscriptionResolver) Title() string {
	return r.Subscription.Title
}

// Type resolves the Type field
func (r *SubscriptionResolver) Type() string {
	return string(r.Subscription.Type)
}

// IsSubscribed resolves the IsPaused field
func (r *SubscriptionResolver) IsSubscribed() bool {
	return r.Subscription.Subscribed
}

// Frequency resolves the Frequency field
func (r *SubscriptionResolver) Frequency() string {
	return string(r.Subscription.Frequency)
}

// CreatedAt resolves the CreatedAt field
func (r *SubscriptionResolver) CreatedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.Subscription.CreatedAt)
	return &t
}

// UpdatedAt resolves the UpdatedAt field
func (r *SubscriptionResolver) UpdatedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.Subscription.UpdatedAt)
	return &t
}

// Subscription adds a syndication source and subscribes to it
func (r *SubscriptionsResolver) Subscription(ctx context.Context, args struct {
	URL scalars.URL
}) (*SubscriptionResolver, error) {
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

	res := &SubscriptionResolver{Subscription: s}

	return res, nil
}

// Subscribe --
func (r *SubscriptionsResolver) Subscribe(ctx context.Context, args struct {
	URL scalars.URL
}) (*SubscriptionResolver, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.SubscribeToSource(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	res := &SubscriptionResolver{Subscription: s}

	return res, nil
}

// Unsubscribe --
func (r *SubscriptionsResolver) Unsubscribe(ctx context.Context, args struct {
	URL scalars.URL
}) (*SubscriptionResolver, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.UnubscribeFromSource(ctx, r.repositories, user, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	res := &SubscriptionResolver{Subscription: s}

	return res, nil
}

// Subscriptions --
func (r *SubscriptionsResolver) Subscriptions(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     *subscriptionSearchInput
}) (*SubscriptionCollectionResolver, error) {
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

	var total int32
	total, err = r.repositories.Subscriptions.GetTotal(ctx, user, terms, showDeleted, pausedOnly)
	if err != nil {
		return nil, err
	}

	var sources []*SubscriptionResolver
	for _, result := range results {
		sources = append(sources, &SubscriptionResolver{
			Subscription: result,
		})
	}

	res := SubscriptionCollectionResolver{
		Results: sources,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
