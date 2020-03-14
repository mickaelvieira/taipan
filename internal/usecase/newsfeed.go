package usecase

import (
	"context"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/domain/newsfeed"
	"github.com/mickaelvieira/taipan/internal/logger"
	"github.com/mickaelvieira/taipan/internal/repository"
)

// AddDocumentToNewsFeeds adds an entry to the users's newsfeeds
func AddDocumentToNewsFeeds(ctx context.Context, repos *repository.Repositories, sourceID string, documentID string) error {
	subscribers, err := repos.Subscriptions.FindSubscribersIDs(ctx, sourceID)
	if err != nil {
		return err
	}

	if len(subscribers) == 0 {
		logger.Warn(fmt.Sprintf("Source [%s] does not have any subscribers", sourceID))
		return nil
	}

	entries := make([]*newsfeed.Entry, len(subscribers))
	for i, s := range subscribers {
		entries[i] = newsfeed.NewEntry(s, documentID)
	}

	if err := repos.NewsFeed.AddEntries(ctx, entries); err != nil {
		return err
	}

	return nil
}
