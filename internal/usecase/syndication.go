package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	nethttp "net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

// DeleteSyndicationSource soft deletes a source
func DeleteSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) (err error) {
	fmt.Printf("Soft deleting source '%s'\n", s.URL)
	s.Deleted = true
	s.UpdatedAt = time.Now()
	return repos.Syndication.Delete(ctx, s)
}

// DisableSyndicationSource soft deletes a source
func DisableSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) (err error) {
	fmt.Printf("Disabling source '%s'\n", s.URL)
	s.Deleted = true
	s.UpdatedAt = time.Now()
	return repos.Syndication.UpdateStatus(ctx, s)
}

// EnableSyndicationSource soft deletes a source
func EnableSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) (err error) {
	fmt.Printf("Enabling source '%s'\n", s.URL)
	s.Deleted = false
	s.UpdatedAt = time.Now()
	return repos.Syndication.UpdateStatus(ctx, s)
}

func handleFeedHTTPErrors(ctx context.Context, repos *repository.Repositories, r *http.Result, s *syndication.Source) (err error) {
	if r.RespStatusCode == nethttp.StatusNotFound {
		var logs []*http.Result
		logs, err = repos.Botlogs.FindByURLAndStatus(ctx, r.ReqURI, r.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			fmt.Printf("Too many '%d' errors\n", r.RespStatusCode)
			err = DeleteSyndicationSource(ctx, repos, s)
			if err != nil {
				return
			}
			fmt.Printf("Source '%s' was marked as deleted\n", s.URL)
		}
	}

	if r.RespStatusCode == nethttp.StatusNotAcceptable ||
		r.RespStatusCode == nethttp.StatusTooManyRequests ||
		r.RespStatusCode == nethttp.StatusInternalServerError {
		var logs []*http.Result
		logs, err = repos.Botlogs.FindByURLAndStatus(ctx, r.ReqURI, r.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			fmt.Printf("Too many '%d' errors\n", r.RespStatusCode)
			err = DisableSyndicationSource(ctx, repos, s)
			if err != nil {
				return
			}
			fmt.Printf("Source '%s' was marked as paused\n", s.URL)
		}
	}
	return
}

func handleDuplicateFeed(ctx context.Context, repos *repository.Repositories, FinalURI *url.URL, s *syndication.Source) (*syndication.Source, error) {
	var b bool
	var err error
	b, err = repos.Syndication.ExistWithURL(ctx, FinalURI)
	if err != nil {
		return s, err
	}

	if !b {
		fmt.Printf("Source's URL needs to be updated %s => %s\n", s.URL, FinalURI)
		s.URL = FinalURI
		s.UpdatedAt = time.Now()
		err = repos.Syndication.UpdateURL(ctx, s)
	} else {
		err = DisableSyndicationSource(ctx, repos, s)
		if err == nil {
			err = fmt.Errorf("Source '%s' was a duplicate. It's been deleted", s.URL)
		}
	}
	return s, err
}

// ParseSyndicationSource in this usecase given an source entity:
// - Fetches the related RSS/ATOM document
// - Parses it the document
// - And returns a list of URLs found in the document
// The document is not parsed if the document has not changed since the last time it was fetched
func ParseSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) (urls []*url.URL, err error) {
	fmt.Printf("Parsing %s\n", s.URL)
	parser := gofeed.NewParser()

	var result, prevResult *http.Result
	prevResult, err = repos.Botlogs.FindLatestByURL(ctx, s.URL)
	result, err = FetchResource(ctx, repos, s.URL)
	if err != nil {
		return
	}

	if result.Failed {
		err = fmt.Errorf("Could not fetch source: '%s'", result.FailureReason)
		return
	}

	err = handleFeedHTTPErrors(ctx, repos, result, s)
	if err != nil {
		return
	}

	if result.RequestWasRedirected() {
		s, err = handleDuplicateFeed(ctx, repos, result.FinalURI, s)
		if err != nil {
			return
		}
	}

	if result.IsContentDifferent(prevResult) {
		fmt.Println("Source's content has changed")
		var content *gofeed.Feed
		content, err = parser.Parse(result.Content)
		if err != nil {
			err = fmt.Errorf("Parsing error: %s - URL %s", err, s.URL)
			return
		}

		s.Title = content.Title
		feedType, errType := syndication.FromGoFeedType(content.FeedType)
		if errType == nil {
			s.Type = feedType
		} else {
			log.Println(errType)
		}

		for _, item := range content.Items {
			u, e := url.FromRawURL(item.Link)
			if e != nil {
				continue // Just skip invalid URLs
			}

			// @TODO Add a list of Source proxy and resolve source's URLs before pushing to the queue

			var b bool
			b, e = repos.Documents.ExistWithURL(ctx, u)
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
		fmt.Println("Source's content has not changed")
	}

	// Reverse results
	for l, r := 0, len(urls)-1; l < r; l, r = l+1, r-1 {
		urls[l], urls[r] = urls[r], urls[l]
	}

	// @TODO Calculate the source update frequency
	// var results []*http.Result
	// results, err = repositories.Botlogs.FindByURL(ctx, s.URL)
	// if err != nil {
	// 	return
	// }

	// f := http.CalculateFrequency(results)
	// fmt.Printf("Source frequency: [%s], previous: [%s]", f, s.Frequency)

	// s.Frequency = f
	s.ParsedAt = time.Now()
	err = repos.Syndication.Update(ctx, s)

	return
}

// ParseSyndicationSourceNew in this usecase given an source entity:
// - Fetches the related RSS/ATOM document
// - Parses it the document
// - And returns a list of URLs found in the document
// The document is not parsed if the document has not changed since the last time it was fetched
func ParseSyndicationSourceNew(ctx context.Context, repos *repository.Repositories, r *http.Result, s *syndication.Source) (urls []*url.URL, err error) {
	fmt.Printf("Parsing %s\n", s.URL)
	parser := gofeed.NewParser()

	pr, err := repos.Botlogs.FindLatestByURL(ctx, s.URL)

	// Store the result of HTTP request
	err = repos.Botlogs.Insert(ctx, r)
	if err != nil {
		return
	}

	// We only want successful requests
	if r.RespStatusCode != nethttp.StatusOK {
		err = fmt.Errorf("Unable to fetch the document: %s", r.RespReasonPhrase)
		return
	}

	if err != nil {
		if r != nil {
			err = handleFeedHTTPErrors(ctx, repos, r, s)
			if err != nil {
				return
			}
		}
		return
	}

	if r.RequestWasRedirected() {
		s, err = handleDuplicateFeed(ctx, repos, r.FinalURI, s)
		if err != nil {
			return
		}
	}

	if r.IsContentDifferent(pr) {
		fmt.Println("Source's content has changed")
		var content *gofeed.Feed
		content, err = parser.Parse(r.Content)
		if err != nil {
			err = fmt.Errorf("Parsing error: %s - URL %s", err, s.URL)
			return
		}

		s.Title = content.Title
		feedType, errType := syndication.FromGoFeedType(content.FeedType)
		if errType == nil {
			s.Type = feedType
		} else {
			log.Println(errType)
		}

		for _, item := range content.Items {
			u, e := url.FromRawURL(item.Link)
			if e != nil {
				continue // Just skip invalid URLs
			}

			// @TODO Add a list of Source proxy and resolve source's URLs before pushing to the queue

			var b bool
			b, e = repos.Documents.ExistWithURL(ctx, u)
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
		fmt.Println("Source's content has not changed")
	}

	// Reverse results
	for l, r := 0, len(urls)-1; l < r; l, r = l+1, r-1 {
		urls[l], urls[r] = urls[r], urls[l]
	}

	// @TODO Calculate the source update frequency
	// var results []*http.Result
	// results, err = repositories.Botlogs.FindByURL(ctx, s.URL)
	// if err != nil {
	// 	return
	// }

	// f := http.CalculateFrequency(results)
	// fmt.Printf("Source frequency: [%s], previous: [%s]", f, s.Frequency)

	// s.Frequency = f
	s.ParsedAt = time.Now()
	err = repos.Syndication.Update(ctx, s)

	fmt.Println("=====================================")

	return
}
