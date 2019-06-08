package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// HandleFeedHTTPErrors handles HTTP errors
func HandleFeedHTTPErrors(ctx context.Context, rs *client.Result, f *feed.Feed, repositories *repository.Repositories) (err error) {
	if rs.RespStatusCode == 404 || rs.RespStatusCode == 500 {
		var logs []*client.Result
		logs, err = repositories.Botlogs.FindByURLAndStatus(ctx, rs.ReqURI, rs.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			err = repositories.Feeds.Delete(ctx, f)
			if err != nil {
				return
			}
			fmt.Printf("Too many %d errors\n", rs.RespStatusCode)
			fmt.Printf("Feed %s was marked as deleted\n", f.URL)
		}
	}
	return
}

// ParseFeed in this usecase given an feed entity:
// - Fetches the related RSS/ATOM document
// - Parses it the document
// - And returns a list of URLs found in the document
// The document is not parsed if the document has not changed since the last time it was fetched
func ParseFeed(ctx context.Context, f *feed.Feed, repositories *repository.Repositories) (urls []string, err error) {
	fmt.Printf("Parsing %s\n", f.URL)
	parser := gofeed.NewParser()

	var result, prevResult *client.Result
	prevResult, err = repositories.Botlogs.FindLatestByURL(ctx, f.URL.String())
	result, err = FetchResource(ctx, f.URL, repositories)
	if err != nil {
		if result != nil {
			err = HandleFeedHTTPErrors(ctx, result, f, repositories)
			if err != nil {
				return
			}
		}
		return
	}

	if result.IsContentDifferent(prevResult) {
		var content *gofeed.Feed
		content, err = parser.Parse(result.Content)
		if err != nil {
			err = fmt.Errorf("Parsing error: %s - URL %s", err, f.URL)
			return
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
			urls = append(urls, item.Link)
		}
	} else {
		fmt.Println("Content has not changed")
	}

	f.ParsedAt = time.Now()
	err = repositories.Feeds.Update(ctx, f)

	return
}
