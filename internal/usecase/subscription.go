package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
)

// SubscribeToSource --
func SubscribeToSource(ctx context.Context, repos *repository.Repositories, u *user.User, src *syndication.Source) (*subscription.Subscription, error) {
	err := repos.Subscriptions.Subscribe(ctx, u, src)
	if err != nil {
		return nil, err
	}

	s, err := repos.Subscriptions.GetByURL(ctx, u, src.URL)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UnubscribeFromSource --
func UnubscribeFromSource(ctx context.Context, repos *repository.Repositories, u *user.User, src *syndication.Source) (*subscription.Subscription, error) {
	err := repos.Subscriptions.Unsubscribe(ctx, u, src)
	if err != nil {
		return nil, err
	}

	s, err := repos.Subscriptions.GetByURL(ctx, u, src.URL)
	if err != nil {
		return nil, err
	}

	return s, nil
}
