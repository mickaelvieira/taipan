package aggregator

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
)

// CalculateUserStatistics --
func CalculateUserStatistics(ctx context.Context, repos *repository.Repositories, u *user.User) (*user.Stats, error) {
	var terms []string

	totalBookmarks, err := repos.Bookmarks.CountAll(ctx, u, terms)
	if err != nil {
		return nil, err
	}

	totalFavorites, err := repos.Bookmarks.CountFavorites(ctx, u)
	if err != nil {
		return nil, err
	}

	totalReadingList, err := repos.Bookmarks.CountReadingList(ctx, u)
	if err != nil {
		return nil, err
	}

	totalSubscriptions, err := repos.Subscriptions.CountUserSubscription(ctx, u)
	if err != nil {
		return nil, err
	}

	s := user.Stats{
		Bookmarks:     totalBookmarks,
		Favorites:     totalFavorites,
		ReadingList:   totalReadingList,
		Subscriptions: totalSubscriptions,
	}

	return &s, nil
}
