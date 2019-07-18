package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/newsfeed"
	"github/mickaelvieira/taipan/internal/repository"
)

// AddDocumentToNewsFeeds adds an entry to the users's newsfeeds
func AddDocumentToNewsFeeds(ctx context.Context, repos *repository.Repositories, sourceID string, documentID string) error {
	subscribers, err := repos.Subscriptions.FindSubscribersIDs(ctx, sourceID)
	if err != nil {
		return err
	}

	if len(subscribers) == 0 {
		return fmt.Errorf("No subscribers to source [%s]", sourceID)
	}

	fmt.Println(subscribers)

	entries := make([]*newsfeed.Entry, len(subscribers))
	for i, s := range subscribers {
		entries[i] = newsfeed.NewEntry(s, documentID)
	}
	fmt.Println(entries)

	err = repos.NewsFeed.AddEntries(ctx, entries)
	if err != nil {
		return err
	}

	return nil
}
