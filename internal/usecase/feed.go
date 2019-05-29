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

	preLogEntry, err = repositories.Botlogs.FindLatestByURI(ctx, f.URL.String())

	http := client.Client{}
	curLogEntry, reader, err = http.Fetch(f.URL.URL)
	if err != nil {
		return entries, err
	}

	err = repositories.Botlogs.Insert(ctx, curLogEntry)
	if err != nil {
		return entries, err
	}

	if curLogEntry.IsContentDifferent(preLogEntry) {
		content, err = parser.Parse(reader)

		f.Title = content.Title
		feedType, errType := feed.FromGoFeedType(content.FeedType)
		if errType == nil {
			f.Type = feedType
		} else {
			log.Println(errType)
		}

		// @TODO We are getting a lot of "Failed to detect feed type" errors,
		// We need to handle this issue
		if err != nil {
			return entries, err
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
		return entries, err
	}

	return entries, nil
}
