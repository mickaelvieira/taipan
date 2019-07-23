package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"time"
)

/* https://godoc.org/golang.org/x/xerrors */

// Bookmarks use cases errors
var (
	ErrInvalidURI           = errors.New("Invalid URL")
	ErrContentHasNotChanged = errors.New("Content has not changed")
)

// DeleteDocument soft deletes a document
func DeleteDocument(ctx context.Context, repos *repository.Repositories, d *document.Document) (err error) {
	fmt.Printf("Soft deleting document '%s'\n", d.URL)
	d.Deleted = true
	d.UpdatedAt = time.Now()
	return repos.Documents.Delete(ctx, d)
}

func handleDuplicateDocument(ctx context.Context, repos *repository.Repositories, originalURI *url.URL, finalURI *url.URL) (err error) {
	var b bool
	var e *document.Document
	fmt.Printf("Request was redirected %s => %s", originalURI, finalURI)
	fmt.Println("Let's check for duplicate")
	// Let's check whether there a document with the requested URI
	e, err = repos.Documents.GetByURL(ctx, originalURI)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		fmt.Println("There is no duplicate")
	} else {
		fmt.Printf("A duplicate was found %v\n", e)

		// There is a document with the old URL
		// Let's check whether there a document with the final URI
		b, err = repos.Documents.ExistWithURL(ctx, finalURI)
		if err != nil {
			return err
		}

		if b {
			// Delete the old one
			// @TODO recreate the users' bookmark with the new URL
			return DeleteDocument(ctx, repos, e)
		}

		fmt.Printf("Document's URL needs to be updated %s => %s\n", e.URL, finalURI)
		e.URL = finalURI
		e.UpdatedAt = time.Now()
		repos.Documents.UpdateURL(ctx, e)
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
func Document(ctx context.Context, repos *repository.Repositories, u *url.URL, findFeeds bool) (*document.Document, error) {
	fmt.Printf("Fetching %s\n", u.String())

	result, err := FetchResource(ctx, repos, u)
	if err != nil {
		return nil, err
	}

	// The problem that we have here is the URL provided by the user or in the feed might be different from
	// the URL we actually store in the DB. (.i.e we get a "cleaned-up" URL from the document itself). So we can't
	// really rely on the URL to identify properly a document.
	// - Can we retrieve the document using its checksum?
	// - What are we going to do when the content changes?
	//
	var d *document.Document
	d, err = repos.Documents.GetByChecksum(ctx, result.Checksum)
	if !findFeeds && d != nil {
		fmt.Println("Document's content has not changed")
		return d, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	d, err = parser.Parse(result.FinalURI, result.Content, findFeeds)
	if err != nil {
		return nil, err
	}
	// Assign document checksum to document
	d.Checksum = result.Checksum

	fmt.Printf("Document was parsed: %s\n", d.URL)

	if result.RequestWasRedirected() {
		err = handleDuplicateDocument(ctx, repos, result.ReqURI, result.FinalURI)
	}

	err = repos.Documents.Upsert(ctx, d)
	if err != nil {
		return nil, err
	}

	err = HandleImage(ctx, repos, d)
	if err != nil {
		return nil, err
	}

	err = repos.Syndication.InsertAllIfNotExists(ctx, d.Feeds)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Bookmark in this use case given a user and a document, we will:
// - Link the document to the user
// - Save it in the DB
// - And finally return the user's bookmark
func Bookmark(ctx context.Context, repos *repository.Repositories, usr *user.User, d *document.Document, isFavorite bool) (*bookmark.Bookmark, error) {
	err := repos.Bookmarks.BookmarkDocument(ctx, usr, d, isFavorite)
	if err != nil {
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

	err = repos.Bookmarks.ChangeFavoriteStatus(ctx, usr, b)
	if err != nil {
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

	err = repos.Bookmarks.ChangeFavoriteStatus(ctx, usr, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Unbookmark removes bookmark from user list
func Unbookmark(ctx context.Context, repos *repository.Repositories, usr *user.User, u *url.URL) (*document.Document, error) {
	var (
		err error
		b   *bookmark.Bookmark
		d   *document.Document
	)

	b, err = repos.Bookmarks.GetByURL(ctx, usr, u)
	if err != nil {
		return nil, err
	}

	b.IsLinked = false
	b.UpdatedAt = time.Now()

	err = repos.Bookmarks.Remove(ctx, usr, b)
	if err != nil {
		return nil, err
	}

	d, err = repos.Documents.GetByID(ctx, b.ID)

	return d, nil
}
