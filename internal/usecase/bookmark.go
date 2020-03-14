package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/domain/bookmark"
	"github.com/mickaelvieira/taipan/internal/domain/document"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"github.com/mickaelvieira/taipan/internal/logger"
	"github.com/mickaelvieira/taipan/internal/parser"
	"github.com/mickaelvieira/taipan/internal/repository"
	"time"
)

/* https://godoc.org/golang.org/x/xerrors */

// Bookmarks use cases errors
// @TODO move this into the domain
var (
	ErrInvalidURI           = errors.New("Invalid URL")
	ErrContentHasNotChanged = errors.New("Content has not changed")
)

// DeleteDocument soft deletes a document
func DeleteDocument(ctx context.Context, repos *repository.Repositories, d *document.Document) error {
	logger.Info(fmt.Sprintf("Soft deleting document '%s'", d.URL))
	d.Deleted = true
	d.UpdatedAt = time.Now()

	if err := repos.Documents.Delete(ctx, d); err != nil {
		return err
	}

	return nil
}

func handleDuplicateDocument(ctx context.Context, repos *repository.Repositories, originalURI *url.URL, finalURI *url.URL) error {
	logger.Warn(fmt.Sprintf("Request was redirected %s => %s", originalURI, finalURI))
	logger.Warn("Looking up duplicates...")
	// Let's check whether there a document with the requested URI
	e, err := repos.Documents.GetByURL(ctx, originalURI)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		logger.Warn("No duplicates!")
	} else {
		logger.Warn(fmt.Sprintf("A duplicate was found %v", e))

		// There is a document with the old URL
		// Let's check whether there a document with the final URI
		b, err := repos.Documents.ExistWithURL(ctx, finalURI)
		if err != nil {
			return err
		}

		if b {
			// Delete the old one
			// @TODO recreate the users' bookmark with the new URL
			if err := DeleteDocument(ctx, repos, e); err != nil {
				return err
			}
		}

		logger.Warn(fmt.Sprintf("Document's URL needs to be updated %s => %s", e.URL, finalURI))
		e.URL = finalURI
		e.UpdatedAt = time.Now()

		if err := repos.Documents.UpdateURL(ctx, e); err != nil {
			return err
		}
	}

	return nil
}

// Document in this use case, given a provided URL, we will from:
// - Fetch the corresponding document
// - Parse the document
// - Upload the document's image to AWS S3
// - Insert/Update the bookmark in the DB
// - Insert new feeds URL in the DB
// - And finally returns the bookmark entity
func Document(ctx context.Context, repos *repository.Repositories, u *url.URL, sourceID string) (*document.Document, error) {
	logger.Info(fmt.Sprintf("Fetching %s", u.String()))

	result, err := FetchResource(ctx, repos, u)
	if err != nil {
		return nil, err
	}

	if !result.RequestWasSuccessful() {
		return nil, fmt.Errorf(result.GetFailureReason())
	}

	// The problem that we have here is the URL provided by the user or in the feed might be different from
	// the URL we actually store in the DB. (.i.e we get a "cleaned-up" URL from the document itself). So we can't
	// really rely on the URL to identify properly a document.
	// - Can we retrieve the document using its checksum?
	// - What are we going to do when the content changes?
	//
	d, err := repos.Documents.GetByChecksum(ctx, result.Checksum)
	if d != nil {
		logger.Info(fmt.Sprintln("Document's content has not changed"))
		return d, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	d, err = parser.Parse(result.FinalURI, result.Content)
	if err != nil {
		return nil, err
	}
	// Assign document checksum to document
	d.Checksum = result.Checksum

	logger.Info(fmt.Sprintf("Document was parsed: %s", d.URL))

	if result.RequestWasRedirected() {
		if err := handleDuplicateDocument(ctx, repos, result.ReqURI, result.FinalURI); err != nil {
			return nil, err
		}
	}

	if err := repos.Documents.Upsert(ctx, d); err != nil {
		return nil, err
	}

	if sourceID != "" {
		d.SourceID = sourceID
		if err := repos.Documents.UpdateSource(ctx, d); err != nil {
			return nil, err
		}
	}

	if err := HandleImage(ctx, repos, d); err != nil {
		// We just log those errors, no need to send them back to the user
		logger.Error(err)
	}

	if err := repos.Syndication.InsertAllIfNotExists(ctx, d.Feeds); err != nil {
		return nil, err
	}

	return d, nil
}

// Bookmark in this use case given a user and a document, we will:
// - Link the document to the user
// - Save it in the DB
// - And finally return the user's bookmark
func Bookmark(ctx context.Context, repos *repository.Repositories, usr *user.User, d *document.Document, isFavorite bool) (*bookmark.Bookmark, error) {
	if err := repos.Bookmarks.BookmarkDocument(ctx, usr, d, isFavorite); err != nil {
		return nil, err
	}

	b, err := repos.Bookmarks.GetByURL(ctx, usr, d.URL)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Favorite mark or unmark a bookmark as favorite
func Favorite(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*bookmark.Bookmark, error) {
	b, err := repos.Bookmarks.GetByURL(ctx, usr, u)
	if err != nil {
		return nil, err
	}

	b.IsFavorite = true
	if b.FavoritedAt.IsZero() {
		b.FavoritedAt = time.Now()
	}
	b.UpdatedAt = time.Now()

	if err := repos.Bookmarks.ChangeFavoriteStatus(ctx, usr, b); err != nil {
		return nil, err
	}

	return b, nil
}

// Unfavorite unmark a bookmark as favorite
func Unfavorite(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*bookmark.Bookmark, error) {
	b, err := repos.Bookmarks.GetByURL(ctx, usr, u)
	if err != nil {
		return nil, err
	}

	b.IsFavorite = false
	b.UpdatedAt = time.Now()

	if err := repos.Bookmarks.ChangeFavoriteStatus(ctx, usr, b); err != nil {
		return nil, err
	}

	return b, nil
}

// Unbookmark removes bookmark from user list
func Unbookmark(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*document.Document, error) {
	b, err := repos.Bookmarks.GetByURL(ctx, usr, u)
	if err != nil {
		return nil, err
	}

	b.IsLinked = false
	b.UpdatedAt = time.Now()

	if err = repos.Bookmarks.Remove(ctx, usr, b); err != nil {
		return nil, err
	}

	d, err := repos.Documents.GetByID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	return d, nil
}
