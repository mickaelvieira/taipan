package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/repository"
	"io"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// ParseFeed in this usecase given an feed entity:
// - Fetches the related RSS/ATOM document
// - Parses it the document
// - And returns a list of URLs found in the document
// The document is not parsed if the document has not changed since the last time it was fetched
func ParseFeed(ctx context.Context, f *feed.Feed, repositories *repository.Repositories) ([]string, error) {
	var err error
	var reader io.Reader
	var content *gofeed.Feed
	var curLogEntry *client.Result
	var preLogEntry *client.Result
	var entries []string

	fmt.Printf("Parsing %s\n", f.URL)
	parser := gofeed.NewParser()

	preLogEntry, err = repositories.Botlogs.FindLatestByURL(ctx, f.URL.String())

	curLogEntry, reader, err = FetchResource(ctx, f.URL, repositories)
	if err != nil {
		if curLogEntry != nil && curLogEntry.RespStatusCode == 404 {
			var logs []*client.Result
			logs, err = repositories.Botlogs.FindByURLAndStatus(ctx, curLogEntry.ReqURI, curLogEntry.RespStatusCode)
			if err != nil {
				return nil, err
			}
			if len(logs) >= 5 {
				err = repositories.Feeds.Delete(ctx, f)
				if err != nil {
					return nil, err
				}
				fmt.Printf("Feed %s was marked as deleted\n", f.URL)
				return entries, nil
			}
		}
		return nil, err
	}

	if curLogEntry.IsContentDifferent(preLogEntry) {
		content, err = parser.Parse(reader)

		// @TODO We are getting a lot of "Failed to detect feed type" errors,
		// We need to handle this issue
		if err != nil {
			return nil, fmt.Errorf("Parsing error: %s - URL %s", err, f.URL)
		}

		f.Title = content.Title
		feedType, errType := feed.FromGoFeedType(content.FeedType)
		if errType == nil {
			f.Type = feedType
		} else {
			log.Println(errType)
		}

		for _, item := range content.Items {
			fmt.Printf("URL %s\n", item.Link)
			entries = append(entries, item.Link)
		}
	} else {
		fmt.Println("Content has not changed")
	}

	f.ParsedAt = time.Now()
	err = repositories.Feeds.Update(ctx, f)

	if err != nil {
		return nil, err
	}

	return entries, nil
}
