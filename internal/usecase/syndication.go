package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// DeleteSyndicationSource soft deletes a source
func DeleteSyndicationSource(ctx context.Context, s *syndication.Source, r *repository.SyndicationRepository) (err error) {
	fmt.Printf("Soft deleting source '%s'\n", s.URL)
	s.Deleted = true
	s.UpdatedAt = time.Now()
	return r.Delete(ctx, s)
}

// DisableSyndicationSource soft deletes a source
func DisableSyndicationSource(ctx context.Context, s *syndication.Source, r *repository.SyndicationRepository) (err error) {
	fmt.Printf("Disabling source '%s'\n", s.URL)
	s.Deleted = true
	s.UpdatedAt = time.Now()
	return r.UpdateStatus(ctx, s)
}

// EnableSyndicationSource soft deletes a source
func EnableSyndicationSource(ctx context.Context, s *syndication.Source, r *repository.SyndicationRepository) (err error) {
	fmt.Printf("Enabling source '%s'\n", s.URL)
	s.Deleted = false
	s.UpdatedAt = time.Now()
	return r.UpdateStatus(ctx, s)
}

func handleFeedHTTPErrors(ctx context.Context, rs *http.Result, s *syndication.Source, repositories *repository.Repositories) (err error) {
	if rs.RespStatusCode == 404 {
		var logs []*http.Result
		logs, err = repositories.Botlogs.FindByURLAndStatus(ctx, rs.ReqURI, rs.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			fmt.Printf("Too many '%d' errors\n", rs.RespStatusCode)
			err = DeleteSyndicationSource(ctx, s, repositories.Syndication)
			if err != nil {
				return
			}
			fmt.Printf("Source '%s' was marked as deleted\n", s.URL)
		}
	}

	if rs.RespStatusCode == 406 || rs.RespStatusCode == 429 || rs.RespStatusCode == 500 {
		var logs []*http.Result
		logs, err = repositories.Botlogs.FindByURLAndStatus(ctx, rs.ReqURI, rs.RespStatusCode)
		if err != nil {
			return
		}
		// @TODO Should we check whether they are actually 5 successive errors?
		if len(logs) >= 5 {
			fmt.Printf("Too many '%d' errors\n", rs.RespStatusCode)
			err = DisableSyndicationSource(ctx, s, repositories.Syndication)
			if err != nil {
				return
			}
			fmt.Printf("Source '%s' was marked as paused\n", s.URL)
		}
	}
	return
}

func handleDuplicateFeed(ctx context.Context, FinalURI *url.URL, s *syndication.Source, repositories *repository.Repositories) (*syndication.Source, error) {
	var b bool
	var err error
	b, err = repositories.Syndication.ExistWithURL(ctx, FinalURI)
	if err != nil {
		return s, err
	}

	if !b {
		fmt.Printf("Source's URL needs to be updated %s => %s\n", s.URL, FinalURI)
		s.URL = FinalURI
		s.UpdatedAt = time.Now()
		err = repositories.Syndication.UpdateURL(ctx, s)
	} else {
		err = DisableSyndicationSource(ctx, s, repositories.Syndication)
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
func ParseSyndicationSource(ctx context.Context, s *syndication.Source, repositories *repository.Repositories) (urls []*url.URL, err error) {
	fmt.Printf("Parsing %s\n", s.URL)
	parser := gofeed.NewParser()

	var result, prevResult *http.Result
	prevResult, err = repositories.Botlogs.FindLatestByURL(ctx, s.URL)
	result, err = FetchResource(ctx, s.URL, repositories)
	if err != nil {
		if result != nil {
			err = handleFeedHTTPErrors(ctx, result, s, repositories)
			if err != nil {
				return
			}
		}
		return
	}

	if result.WasRedirected {
		s, err = handleDuplicateFeed(ctx, result.FinalURI, s, repositories)
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
		fmt.Println("Source's content has not changed")
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
	err = repositories.Syndication.Update(ctx, s)

	return
}
