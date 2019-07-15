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

func handleFeedHTTPErrors(ctx context.Context, repos *repository.Repositories, r *http.Result, s *syndication.Source) error {
	// no error or failure, there is nothing to do
	if !syndication.IsHTTPError(r.RespStatusCode) && !r.Failed {
		return nil
	}

	if r.Failed {
		// @TODO Should we check whether they are actually 5 successive errors?
		l, err := repos.Botlogs.FindFailureByURL(ctx, r.ReqURI)
		if err != nil {
			return err
		}

		if len(l) < 5 {
			return nil
		}

		fmt.Printf("Failed request: '%s' was marked as paused\n", s.URL)
		return DisableSyndicationSource(ctx, repos, s)
	}

	if syndication.IsHTTPError(r.RespStatusCode) {
		// @TODO Should we check whether they are actually 5 successive errors?
		l, err := repos.Botlogs.FindByURLAndStatus(ctx, r.ReqURI, r.RespStatusCode)
		if err != nil {
			return err
		}

		if len(l) < 5 {
			return nil
		}

		if syndication.IsHTTPErrorPermanent(r.RespStatusCode) {
			fmt.Printf("Unexisting source: '%s' was marked as deleted\n", s.URL)
			return DeleteSyndicationSource(ctx, repos, s)
		}
		if syndication.IsHTTPErrorTemporary(r.RespStatusCode) {
			fmt.Printf("Server error: '%s' was marked as paused\n", s.URL)
			return DisableSyndicationSource(ctx, repos, s)
		}
	}

	return nil
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
func ParseSyndicationSource(ctx context.Context, repos *repository.Repositories, r *http.Result, s *syndication.Source) ([]*url.URL, error) {
	var urls []*url.URL

	if err := handleFeedHTTPErrors(ctx, repos, r, s); err != nil {
		return urls, err
	}

	// We only want successful requests at this point
	if r.RespStatusCode != nethttp.StatusOK {
		if r.Failed {
			return urls, fmt.Errorf("%s", r.FailureReason)
		}
		return urls, fmt.Errorf("%s", r.RespReasonPhrase)
	}

	if r.RequestWasRedirected() {
		var err error
		s, err = handleDuplicateFeed(ctx, repos, r.FinalURI, s)
		if err != nil {
			return urls, err
		}
	}

	pr, err := repos.Botlogs.FindPreviousByURL(ctx, s.URL, r)
	if err != nil {
		return urls, err
	}

	if r.IsContentDifferent(pr) {
		c, err := gofeed.NewParser().Parse(r.Content)
		if err != nil {
			return urls, fmt.Errorf("Parsing error: %s - URL %s", err, s.URL)
		}

		if s.Title == "" {
			s.Title = c.Title
		}

		if s.Type == "" {
			feedType, e := syndication.FromGoFeedType(c.FeedType)
			if e == nil {
				s.Type = feedType
			} else {
				log.Println(e)
			}
		}

		for _, item := range c.Items {
			u, e := url.FromRawURL(item.Link)
			if e != nil {
				continue // Just skip invalid URLs
			}

			// @TODO Add a list of Source proxy and resolve source's URLs before pushing to the queue
			b, e := repos.Documents.ExistWithURL(ctx, u)
			if e != nil {
				log.Println(e)
				continue
			}
			if !b {
				urls = append(urls, u)
			}
		}
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

	return urls, err
}
