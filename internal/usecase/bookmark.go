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
func DeleteDocument(ctx context.Context, d *document.Document, r *repository.DocumentRepository) (err error) {
	fmt.Printf("Soft deleting document '%s'\n", d.URL)
	d.Deleted = true
	d.UpdatedAt = time.Now()
	return r.Delete(ctx, d)
}

func handleDuplicateDocument(ctx context.Context, originalURI *url.URL, finalURI *url.URL, repositories *repository.Repositories) (err error) {
	var b bool
	var e *document.Document
	fmt.Printf("Request was redirected %s => %s", originalURI, finalURI)
	fmt.Println("Let's check for duplicate")
	// Let's check whether there a document with the requested URI
	e, err = repositories.Documents.GetByURL(ctx, originalURI)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		fmt.Println("There is no duplicate")
	} else {
		fmt.Printf("A duplicate was found %v\n", e)

		// There is a document with the old URL
		// Let's check whether there a document with the final URI
		b, err = repositories.Documents.ExistWithURL(ctx, finalURI)
		if err != nil {
			return err
		}

		if b {
			// Delete the old one
			// @TODO recreate the users' bookmark with the new URL
			return DeleteDocument(ctx, e, repositories.Documents)
		}

		fmt.Printf("Document's URL needs to be updated %s => %s\n", e.URL, finalURI)
		e.URL = finalURI
		e.UpdatedAt = time.Now()
		repositories.Documents.UpdateURL(ctx, e)
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
func Document(ctx context.Context, URL *url.URL, repositories *repository.Repositories) (*document.Document, error) {
	fmt.Printf("Fetching %s\n", URL.String())

	result, err := FetchResource(ctx, URL, repositories)
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
	d, err = repositories.Documents.GetByChecksum(ctx, result.Checksum)
	if d != nil {
		fmt.Println("Document's content has not changed")
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

	fmt.Printf("Document was parsed: %s\n", d.URL)

	if result.WasRedirected {
		err = handleDuplicateDocument(ctx, result.ReqURI, result.FinalURI, repositories)
	}

	err = repositories.Documents.Upsert(ctx, d)

	if err != nil {
		return nil, err
	}

	err = HandleImage(ctx, d, repositories)
	if err != nil {
		return nil, err
	}

	err = repositories.Feeds.InsertAllIfNotExists(ctx, d.Feeds)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Bookmark in this use case given a user and a document, we will:
// - Link the document to the user
// - Save it in the DB
// - And finally return the user's bookmark
func Bookmark(ctx context.Context, user *user.User, d *document.Document, repositories *repository.Repositories) (*bookmark.Bookmark, error) {
	err := repositories.Bookmarks.BookmarkDocument(ctx, user, d)
	if err != nil {
		return nil, err
	}

	b, err := repositories.Bookmarks.GetByURL(ctx, user, d.URL)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ReadStatus changes the bookmark read status
func ReadStatus(ctx context.Context, user *user.User, URL *url.URL, isRead bool, repositories *repository.Repositories) (*bookmark.Bookmark, error) {
	var err error
	var b *bookmark.Bookmark

	b, err = repositories.Bookmarks.GetByURL(ctx, user, URL)
	if err != nil {
		return nil, err
	}

	b.IsRead = isRead
	b.UpdatedAt = time.Now()

	err = repositories.Bookmarks.ChangeReadStatus(ctx, user, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Unbookmark removes bookmark from user list
func Unbookmark(ctx context.Context, user *user.User, URL *url.URL, repositories *repository.Repositories) (*document.Document, error) {
	var err error
	var b *bookmark.Bookmark
	var d *document.Document

	b, err = repositories.Bookmarks.GetByURL(ctx, user, URL)
	if err != nil {
		return nil, err
	}

	b.IsLinked = false
	// b.IsRead = false // @TODO
	b.UpdatedAt = time.Now()

	err = repositories.Bookmarks.Remove(ctx, user, b)
	if err != nil {
		return nil, err
	}

	d, err = repositories.Documents.GetByID(ctx, b.ID)

	return d, nil
}
