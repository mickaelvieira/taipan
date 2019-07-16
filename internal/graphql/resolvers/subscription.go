package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/repository"

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

// URL resolves the URL
func (r *SubscriptionResolver) URL() scalars.URL {
	return scalars.NewURL(r.Subscription.URL)
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

// CreatedAt resolves the CreatedAt field
func (r *SubscriptionResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.Subscription.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *SubscriptionResolver) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.Subscription.UpdatedAt)
}

// Subscriptions --
func (r *SubscriptionsResolver) Subscriptions(ctx context.Context, args struct {
	Pagination offsetPaginationInput
}) (*SubscriptionCollectionResolver, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)

	results, err := r.repositories.Subscriptions.FindAll(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Subscriptions.GetTotal(ctx, user)
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
