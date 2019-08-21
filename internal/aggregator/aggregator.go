package aggregator

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
)

// AggType type of aggregation
type AggType string

// List of aggregation available
const (
	Bookmarks     AggType = "bookmarks"
	Favorites     AggType = "favorites"
	ReadingList   AggType = "readingList"
	Subscriptions AggType = "subscriptions"
)

// LoaderKey contains the user and the type of aggregator we want to perform
type LoaderKey struct {
	User *user.User
	Type AggType
}

// String implementation of the dataloader.Key interface
func (k *LoaderKey) String() string {
	return k.User.ID + "-" + string(k.Type)
}

// Raw implementation of the dataloader.Key interface
func (k *LoaderKey) Raw() interface{} {
	return k
}

// NewLoaderKey creates a new LoaderKey
func NewLoaderKey(u *user.User, t AggType) *LoaderKey {
	return &LoaderKey{
		User: u,
		Type: t,
	}
}

// Aggregate --
func Aggregate(ctx context.Context, repos *repository.Repositories, u *user.User, t AggType) (int32, error) {
	var total int32
	var err error
	switch t {
	case Bookmarks:
		var terms []string
		total, err = repos.Bookmarks.CountAll(ctx, u, terms)
		if err != nil {
			return total, err
		}
	case Favorites:
		total, err = repos.Bookmarks.CountFavorites(ctx, u)
		if err != nil {
			return total, err
		}
	case ReadingList:
		total, err = repos.Bookmarks.CountReadingList(ctx, u)
		if err != nil {
			return total, err
		}
	case Subscriptions:
		total, err = repos.Subscriptions.CountUserSubscription(ctx, u)
		if err != nil {
			return total, err
		}
	default:
		return total, errors.New("Incorrect aggregation type")
	}

	return total, nil
}
