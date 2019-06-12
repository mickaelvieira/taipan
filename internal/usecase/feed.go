package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// DeleteFeed soft deletes a feed
func DeleteFeed(ctx context.Context, f *feed.Feed, r *repository.FeedRepository) (err error) {
	fmt.Printf("Soft deleting feed '%s'\n", f.URL)
	f.Deleted = true
	f.UpdatedAt = time.Now()
	return r.Delete(ctx, f)
}

func handleFeedHTTPErrors(ctx context.Context, rs *client.Result, f *feed.Feed, repositories *repository.Repositories) (err error) {
	if rs.RespStatusCode == 404 || rs.RespStatusCode == 429 || rs.RespStatusCode == 500 {
		var logs []*client.Result
		logs, err = repositories.Botlogs.FindByURLAndStatus(ctx, rs.ReqURI, rs.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			fmt.Printf("Too many '%d' errors\n", rs.RespStatusCode)
			err = DeleteFeed(ctx, f, repositories.Feeds)
			if err != nil {
				return
			}
			fmt.Printf("Feed '%s' was marked as deleted\n", f.URL)
		}
	}
	return
}

func handleDuplicateFeed(ctx context.Context, FinalURI *url.URL, f *feed.Feed, repositories *repository.Repositories) (*feed.Feed, error) {
	var b bool
	var err error
	b, err = repositories.Feeds.ExistWithURL(ctx, FinalURI)
	if err != nil {
		return f, err
	}

	if !b {
		fmt.Printf("Feed's URL needs to be updated %s => %s\n", f.URL, FinalURI)
		f.URL = FinalURI
		f.UpdatedAt = time.Now()
		err = repositories.Feeds.UpdateURL(ctx, f)
	} else {
		err = DeleteFeed(ctx, f, repositories.Feeds)
		if err == nil {
			err = fmt.Errorf("Feed '%s' was a duplicate. It's been deleted", f.URL)
		}
	}
	return f, err
}

// ParseFeed in this usecase given an feed entity:
// - Fetches the related RSS/ATOM document
// - Parses it the document
// - And returns a list of URLs found in the document
// The document is not parsed if the document has not changed since the last time it was fetched
func ParseFeed(ctx context.Context, f *feed.Feed, repositories *repository.Repositories) (urls []*url.URL, err error) {
	fmt.Printf("Parsing %s\n", f.URL)
	parser := gofeed.NewParser()

	var result, prevResult *client.Result
	prevResult, err = repositories.Botlogs.FindLatestByURL(ctx, f.URL)
	result, err = FetchResource(ctx, f.URL, repositories)
	if err != nil {
		if result != nil {
			err = handleFeedHTTPErrors(ctx, result, f, repositories)
			if err != nil {
				return
			}
		}
		return
	}

	if result.WasRedirected {
		f, err = handleDuplicateFeed(ctx, result.FinalURI, f, repositories)
		if err != nil {
			return
		}
	}

	if result.IsContentDifferent(prevResult) {
		fmt.Println("Feed's content has changed")
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
			u, e := url.FromRawURL(item.Link)
			if e != nil {
				continue // Just skip invalid URLs
			}

			// @TODO Add a list of Feed proxy and resolve feed's URLs before pushing to the queue

			var b bool
			b, e = repositories.Documents.ExistWithURL(ctx, u)
			if e != nil {
				log.Println(e)
				continue
			}
			if !b {
				fmt.Printf("New document '%s'\n", u)
				urls = append(urls, u)
			} else {
				fmt.Printf("Document already exists '%s'\n", u)
			}
		}
	} else {
		fmt.Println("Feed's content has not changed")
	}

	f.ParsedAt = time.Now()
	err = repositories.Feeds.Update(ctx, f)

	return
}
