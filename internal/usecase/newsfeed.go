package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/newsfeed"
	"github/mickaelvieira/taipan/internal/repository"
)

// AddDocumentToNewsFeeds adds an entry to the users's newsfeeds
func AddDocumentToNewsFeeds(ctx context.Context, repos *repository.Repositories, sourceID string, documentID string) error {
	subscribers, err := repos.Subscriptions.FindSubscribersIDs(ctx, sourceID)
	if err != nil {
		return err
	}

	entries := make([]*newsfeed.Entry, len(subscribers))
	for i, s := range subscribers {
		entries[i] = newsfeed.NewEntry(s, documentID)
	}

	err = repos.NewsFeed.AddEntries(ctx, entries)
	if err != nil {
		return err
	}

	return nil
}
