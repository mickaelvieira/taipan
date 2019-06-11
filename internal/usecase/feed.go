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

// HandleFeedHTTPErrors handles HTTP errors
func HandleFeedHTTPErrors(ctx context.Context, rs *client.Result, f *feed.Feed, repositories *repository.Repositories) (err error) {
	if rs.RespStatusCode == 404 || rs.RespStatusCode == 429 || rs.RespStatusCode == 500 {
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
func ParseFeed(ctx context.Context, f *feed.Feed, repositories *repository.Repositories) (urls []*url.URL, err error) {
	fmt.Printf("Parsing %s\n", f.URL)
	parser := gofeed.NewParser()

	var result, prevResult *client.Result
	prevResult, err = repositories.Botlogs.FindLatestByURL(ctx, f.URL)
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

	if result.WasRedirected {
		// Is there already a feed with the new URL?
		var b bool
		b, err = repositories.Feeds.ExistWithURL(ctx, result.FinalURI)
		if err != nil {
			return
		}

		if !b {
			// There is no duplicate
			// We can update the feed with the new URL safely
			fmt.Printf("Feed's URL needs to be updated %s => %s\n", f.URL, result.FinalURI)
			f.URL = result.FinalURI
			err = repositories.Feeds.UpdateURL(ctx, f)
			if err != nil {
				return
			}
		} else {
			fmt.Printf("Delete duplicate feed %s\n", f.URL)
			// There is a duplicate
			// delete this feed
			err = repositories.Feeds.Delete(ctx, f)
			// and stop processing
			// we get data from the duplicate
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

			// @TODO we can probable do it in a concurrent way
			// r, e := ResolveURL(u)
			// if e != nil {
			// 	log.Println(e)
			// 	continue // Just skip URLs we could not resolve
			// }

			// @TODO I need to see how I can handle servers that don't allow HEAD method
			// if r.RespStatusCode != 200 {
			// 	log.Printf("Incorrect status code: %d %s", r.RespStatusCode, r.RespReasonPhrase)
			// 	continue // Just skip unsuccessful request
			// }

			var b bool
			b, e = repositories.Documents.ExistWithURL(ctx, u)
			if e != nil {
				log.Println(e)
				continue
			}
			if !b {
				fmt.Printf("===> ADD URL %s\n", u)
				urls = append(urls, u)
			} else {
				fmt.Printf("Document already exists %s\n", u)
			}
		}
	} else {
		fmt.Println("Feed's content has not changed")
	}

	f.ParsedAt = time.Now()
	err = repositories.Feeds.Update(ctx, f)

	return
}
