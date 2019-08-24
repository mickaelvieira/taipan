package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/graphql/loaders"
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
}

// ID resolves the ID field
func (r *Subscription) ID() gql.ID {
	return gql.ID(r.subscription.ID)
}

// User resolves the User field
func (r *Subscription) User(ctx context.Context) (*User, error) {
	// NOTE: When users look up the list of sources, we do return a list
	// of subscription but the logged in user might not have subscribed to it
	// so the UserID is empty in this case
	if r.subscription.UserID == "" {
		return nil, nil
	}

	l := loaders.FromContext(ctx)
	if l == nil {
		return nil, ErrLoadersNotFound
	}

	d, err := l.Users.Load(ctx, dataloader.StringKey(r.subscription.UserID))()
	if err != nil {
		return nil, err
	}

	u, ok := d.(*user.User)
	if !ok {
		return nil, ErrDataTypeIsNotValid
	}

	return resolve(r.repositories).user(u), nil
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
func (r *SubscriptionRootResolver) Subscription(ctx context.Context, a struct {
	URL scalars.URL
}) (*Subscription, error) {
	u := a.URL.ToDomain()
	user := auth.FromContext(ctx)

	_, err := usecase.CreateSyndicationSource(ctx, r.repositories, u, false)
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
func (r *SubscriptionRootResolver) Subscribe(ctx context.Context, a struct {
	URL scalars.URL
}) (*Subscription, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.SubscribeToSource(ctx, r.repositories, user, a.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).subscription(s), nil
}

// Unsubscribe --
func (r *SubscriptionRootResolver) Unsubscribe(ctx context.Context, a struct {
	URL scalars.URL
}) (*Subscription, error) {
	user := auth.FromContext(ctx)

	s, err := usecase.UnubscribeFromSource(ctx, r.repositories, user, a.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).subscription(s), nil
}

// Subscriptions --
func (r *SubscriptionRootResolver) Subscriptions(ctx context.Context, a struct {
	Pagination OffsetPaginationInput
	Search     *SubscriptionSearchInput
}) (*SubscriptionCollection, error) {
	user := auth.FromContext(ctx)
	p := offsetPagination(10)(a.Pagination)
	s := &repository.SubscriptionSearchParams{}

	if a.Search != nil {
		s.Terms = a.Search.Terms
	}

	if a.Search.Tags != nil {
		s.Tags = a.Search.Tags
	}

	results, err := r.repositories.Subscriptions.FindAll(ctx, user, s, p)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Subscriptions.GetTotal(ctx, user, s)
	if err != nil {
		return nil, err
	}

	res := SubscriptionCollection{
		Results: resolve(r.repositories).subscriptions(results),
		Total:   total,
		Offset:  p.Offset,
		Limit:   p.Limit,
	}

	return &res, nil
}

// Tags resolves the query
func (r *SubscriptionRootResolver) Tags(ctx context.Context) (*TagCollection, error) {
	ids, err := r.repositories.Syndication.GetActiveTagIDs(ctx)
	if err != nil {
		return nil, err
	}

	l := loaders.FromContext(ctx)

	var keys = make([]dataloader.Key, len(ids))
	for i, id := range ids {
		keys[i] = dataloader.StringKey(id)
	}

	future := l.SyndicationTag.LoadMany(ctx, keys)
	data, e := future()
	if len(e) > 0 {
		return nil, e[0]
	}

	var resolver = resolve(r.repositories)
	var tags = make([]*Tag, len(data))

	for i, datum := range data {
		result, ok := datum.(*syndication.Tag)
		if !ok {
			return nil, ErrDataTypeIsNotValid
		}
		tags[i] = resolver.tag(result)
	}

	res := TagCollection{
		Results: tags,
		Total:   int32(len(tags)),
	}

	return &res, nil
}
