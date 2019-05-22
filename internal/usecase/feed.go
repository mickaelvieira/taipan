package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/repository"
	"io"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// ParseFeed parses a feed
func ParseFeed(ctx context.Context, feed *feed.Feed, repositories *repository.Repositories) ([]string, error) {
	var err error
	var reader io.Reader
	var content *gofeed.Feed
	var curLogEntry *client.Result
	var preLogEntry *client.Result
	var entries []string

	log.Printf("Parsing %s", feed.URL)
	parser := gofeed.NewParser()

	preLogEntry, err = repositories.Botlogs.FindLatestByURI(ctx, feed.URL.String())

	http := client.Client{}
	curLogEntry, reader, err = http.Fetch(feed.URL.URL)
	if err != nil {
		return entries, err
	}

	err = repositories.Botlogs.Insert(ctx, curLogEntry)
	if err != nil {
		return entries, err
	}

	if curLogEntry.IsContentDifferent(preLogEntry) {
		content, err = parser.Parse(reader)
		if err != nil {
			return entries, err
		}

		for _, item := range content.Items {
			entries = append(entries, item.Link)
		}

		feed.ParsedAt = time.Now()

		repositories.Feeds.Update(ctx, feed)
	} else {
		log.Println("content has not changed")
	}

	return entries, nil
}
