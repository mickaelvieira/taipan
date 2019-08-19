package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"time"

	"github.com/mmcdole/gofeed"
)

// DisableSyndicationSource soft deletes a source
func DisableSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) error {
	logger.Warn(fmt.Sprintf("Disabling source '%s'", s.URL))
	s.IsDeleted = true
	s.UpdatedAt = time.Now()

	if err := repos.Syndication.UpdateVisibility(ctx, s); err != nil {
		return err
	}

	return nil
}

// EnableSyndicationSource soft deletes a source
func EnableSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) error {
	logger.Warn(fmt.Sprintf("Enabling source '%s'", s.URL))
	s.IsDeleted = false
	s.UpdatedAt = time.Now()

	if err := repos.Syndication.UpdateVisibility(ctx, s); err != nil {
		return err
	}

	return nil
}

// PauseSyndicationSource soft deletes a source
func PauseSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) error {
	logger.Warn(fmt.Sprintf("Pausing source '%s'", s.URL))
	s.IsPaused = true
	s.UpdatedAt = time.Now()

	if err := repos.Syndication.UpdateStatus(ctx, s); err != nil {
		return err
	}

	return nil
}

// ResumeSyndicationSource soft deletes a source
func ResumeSyndicationSource(ctx context.Context, repos *repository.Repositories, s *syndication.Source) error {
	logger.Warn(fmt.Sprintf("Resuming source '%s'", s.URL))
	s.IsPaused = false
	s.UpdatedAt = time.Now()

	if err := repos.Syndication.UpdateStatus(ctx, s); err != nil {
		return err
	}

	return nil
}

// UpdateSourceTitle soft deletes a source
func UpdateSourceTitle(ctx context.Context, repos *repository.Repositories, s *syndication.Source, t string) error {
	s.Title = t
	s.UpdatedAt = time.Now()

	if err := repos.Syndication.UpdateTitle(ctx, s); err != nil {
		return err
	}

	return nil
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

		logger.Warn(fmt.Sprintf("Failed request: '%s' was marked as paused", s.URL))
		return PauseSyndicationSource(ctx, repos, s)
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
			logger.Warn(fmt.Sprintf("Unexisting source: '%s' was marked as deleted", s.URL))
			return DisableSyndicationSource(ctx, repos, s)
		}
		if syndication.IsHTTPErrorTemporary(r.RespStatusCode) {
			logger.Warn(fmt.Sprintf("Server error: '%s' was marked as paused", s.URL))
			return PauseSyndicationSource(ctx, repos, s)
		}
	}

	return nil
}

func handleDuplicateFeed(ctx context.Context, repos *repository.Repositories, FinalURI *url.URL, s *syndication.Source) (*syndication.Source, error) {
	b, err := repos.Syndication.ExistWithURL(ctx, FinalURI)
	if err != nil {
		return s, err
	}

	if !b {
		logger.Warn(fmt.Sprintf("Source's URL needs to be updated %s => %s", s.URL, FinalURI))
		s.URL = FinalURI
		s.UpdatedAt = time.Now()
		if err := repos.Syndication.UpdateURL(ctx, s); err != nil {
			return s, err
		}
	} else {
		if err := PauseSyndicationSource(ctx, repos, s); err != nil {
			return s, err
		}
		logger.Warn(fmt.Sprintf("Source '%s' was a duplicate. It's been deleted", s.URL))
	}

	return s, nil
}

// CreateSyndicationSource in this use case given a url, we will:
// - fetch the related feed
// - parse the feed
// - And finally return a web syndication source
func CreateSyndicationSource(ctx context.Context, repos *repository.Repositories, u *url.URL, isPaused bool) (*syndication.Source, error) {
	if syndication.IsBlacklisted(u.String()) {
		return nil, fmt.Errorf("URL %s is blacklisted", u.String())
	}

	s, err := repos.Syndication.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			s = syndication.NewSource(u, "", "")
			s.IsPaused = isPaused
			err = repos.Syndication.Insert(ctx, s)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	result, err := FetchResource(ctx, repos, s.URL)
	if err != nil {
		return nil, err
	}

	// @TODO URLs must be pushed to otherwise if the feed does not change they will never be sent to the parser
	_, err = ParseSyndicationSource(ctx, repos, result, s)
	if err != nil {
		return nil, err
	}

	return s, nil
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
	if r.RequestHasFailed() {
		return urls, fmt.Errorf("%s", r.GetFailureReason())
	}

	if r.RequestWasRedirected() {
		var err error
		s, err = handleDuplicateFeed(ctx, repos, r.FinalURI, s)
		if err != nil {
			return urls, err
		}
	}

	pr, err := repos.Botlogs.FindPreviousByURL(ctx, s.URL, r)
	if err != nil && err != sql.ErrNoRows {
		return urls, err
	}

	if r.IsContentDifferent(pr) {
		c, err := gofeed.NewParser().Parse(r.Content)
		if err != nil {
			return urls, fmt.Errorf("Parsing error: %s - URL %s", err, s.URL)
		}

		if c.Title != "" {
			if s.Title == "" || s.Title == syndication.DefaultWPFeedTitle {
				s.Title = c.Title
			}
		}

		if c.Link != "" {
			l, err := url.FromRawURL(c.Link)
			if err == nil {
				s.Domain = l
			}
		}

		if s.Type == "" {
			feedType, err := syndication.FromGoFeedType(c.FeedType)
			if err == nil {
				s.Type = feedType
			} else {
				logger.Error(err)
			}
		}

		for _, item := range c.Items {
			u, err := url.FromRawURL(item.Link)
			if err != nil {
				logger.Error(err)
				continue // Just skip invalid URLs
			}

			// @TODO Add a list of Source proxy and resolve source's URLs before pushing to the queue
			b, err := repos.Documents.ExistWithURL(ctx, u)
			if err != nil {
				logger.Error(err)
				continue
			}
			if !b {
				logger.Warn(fmt.Sprintf("Adding URL [%s]", u))
				urls = append(urls, u)
			} else {
				logger.Warn(fmt.Sprintf("URL [%s] already exists", u))
			}
		}
	} else {
		logger.Info("Feed content has not changed")
	}

	// Reverse results
	for l, r := 0, len(urls)-1; l < r; l, r = l+1, r-1 {
		urls[l], urls[r] = urls[r], urls[l]
	}

	var results []*http.Result
	results, err = repos.Botlogs.FindByURL(ctx, s.URL)
	if err != nil {
		return urls, err
	}

	f := http.CalculateFrequency(results)
	logger.Info(fmt.Sprintf("Source frequency: [%s], previous: [%s]", f, s.Frequency))

	s.Frequency = f
	s.ParsedAt = time.Now()

	if err := repos.Syndication.Update(ctx, s); err != nil {
		return urls, err
	}

	return urls, nil
}
