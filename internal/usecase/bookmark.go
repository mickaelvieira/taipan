package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
	"io"
	"log"
	"net/url"
	"time"
)

/* https://godoc.org/golang.org/x/xerrors */

// Bookmarks use cases errors
var (
	ErrInvalidURI           = errors.New("Invalid URL")
	ErrContentHasNotChanged = errors.New("Content has not changed")
)

// Document in this use case, given a provided URL, we will from:
// - Fetch the corresponding document
// - Parse the document
// - Upload the document's image to AWS S3
// - Insert/Update the bookmark in the DB
// - Insert new feeds URL in the DB
// - And finally returns the bookmark entity
func Document(ctx context.Context, rawURL string, repositories *repository.Repositories) (*document.Document, error) {
	cl := client.Client{}
	URL, err := url.ParseRequestURI(rawURL)
	if err != nil || !URL.IsAbs() {
		return nil, ErrInvalidURI
	}

	// @TODO that might be nice to do a HEAD request
	// to get the last modified date before fetching the entire document
	var reader io.Reader
	var result *client.Result
	result, reader, err = cl.Fetch(URL)
	// We might have a non-HTTP error
	if err != nil {
		return nil, err
	}

	// Store the result of HTTP request
	if result != nil {
		err = repositories.Botlogs.Insert(ctx, result)
		if err != nil {
			return nil, err
		}
	}

	if result.RespStatusCode != 200 {
		return nil, fmt.Errorf("Unable to fetch the document: %s", result.RespReasonPhrase)
	}

	// The problem that we have here is the URL provided by the user or in the feed might be different from
	// the URL we actually store in the DB. (.i.e we get a "cleaned-up" URL from the document itself). So we can't
	// really rely on the URL to identify properly a document.
	// - Can we retrieve the document using its checksum?
	// - What are we going to do when the content changes?
	//
	var d *document.Document
	d, err = repositories.Documents.GetByChecksum(ctx, result.Checksum)
	if err == nil {
		fmt.Println("Document has not changed")
		return d, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// @TODO I need to check client's result before parsing
	d, err = parser.Parse(URL, reader)
	if err != nil {
		return nil, err
	}

	d.Checksum = result.Checksum

	fmt.Printf("Documents: %s\n", d.URL)

	// @TODO there is a bug here
	// If the document already exists with a URL starting with http:// the document gets duplicated
	err = repositories.Documents.Upsert(ctx, d)
	if err != nil {
		return nil, err
	}

	// An image was found in the document
	if d.Image != nil {
		err = s3.Upload(d.Image)
		if err != nil {
			log.Println(err) // @TODO we might eventually better handle this case
		} else {
			err = repositories.Documents.UpdateImage(ctx, d)
			if err != nil {
				return nil, err
			}
		}
	}

	err = repositories.Feeds.InsertAllIfNotExists(ctx, d.Feeds, d)
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

// ReadStatus bla
func ReadStatus(ctx context.Context, user *user.User, URL *uri.URI, isRead bookmark.ReadStatus, repositories *repository.Repositories) (*bookmark.Bookmark, error) {
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
func Unbookmark(ctx context.Context, user *user.User, URL *uri.URI, repositories *repository.Repositories) (*document.Document, error) {
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
