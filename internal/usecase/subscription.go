package usecase

import (
	"context"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/domain/subscription"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"github.com/mickaelvieira/taipan/internal/logger"
	"github.com/mickaelvieira/taipan/internal/repository"
)

// SubscribeToSource --
func SubscribeToSource(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*subscription.Subscription, error) {
	src, err := repos.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	// make sure the source is not paused
	if src.IsPaused {
		if err := ResumeSyndicationSource(ctx, repos, src); err != nil {
			return nil, err
		}
	}

	err = repos.Subscriptions.Subscribe(ctx, usr, src)
	if err != nil {
		return nil, err
	}

	s, err := repos.Subscriptions.GetByURL(ctx, usr, src.URL)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UnubscribeFromSource --
func UnubscribeFromSource(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*subscription.Subscription, error) {
	src, err := repos.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	err = repos.Subscriptions.Unsubscribe(ctx, usr, src)
	if err != nil {
		return nil, err
	}

	subscribers, err := repos.Subscriptions.FindSubscribersIDs(ctx, src.ID)
	if err != nil {
		return nil, err
	}

	if len(subscribers) == 0 && !src.IsPaused {
		if err := PauseSyndicationSource(ctx, repos, src); err != nil {
			return nil, err
		}
		logger.Warn(fmt.Sprintf("Source [%s] does not have any subscribers, it was marked as paused", src.URL))
	}

	s, err := repos.Subscriptions.GetByURL(ctx, usr, src.URL)
	if err != nil {
		return nil, err
	}

	return s, nil
}
