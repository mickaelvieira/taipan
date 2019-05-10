package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
	"io"
	"log"
	"net/url"
)

// Bookmarks use cases errors
var (
	ErrInvalidURI           = errors.New("Invalid URL")
	ErrContentHasNotChanged = errors.New("Content has not changed")
)

// Bookmark in this use case, given a provided URL, we will from:
// - Fetch the corresponding document
// - Parse the document
// - Upload the document's image to AWS S3
// - Insert/Update the bookmark in the DB
// - Insert new feeds URL in the DB
// - And finally returns the bookmark entity
func Bookmark(ctx context.Context, rawURL string, repositories *repository.Repositories) (*bookmark.Bookmark, error) {
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
	if err != nil {
		return nil, err
	}

	// The problem that we have here is the URL provided by the user or in the feed might be different from
	// the URL we actually store in the DB. (.i.e we get a "cleaned-up" URL from the document itself). So we can't
	// really rely on the URL to identify properly a document.
	// - Can we retrieve the document using its checksum?
	// - What are we going to do when the content changes?
	//
	var b *bookmark.Bookmark
	b, err = repositories.Bookmarks.GetByChecksum(ctx, result.Checksum)
	if err == nil {
		return b, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// @TODO Don't parse the document if it hasn't changed
	// @TODO I need to check client's result before parsing
	var d *parser.Document
	d, err = parser.Parse(URL, reader)
	if err != nil {
		return nil, err
	}

	b = d.ToBookmark()
	b.Checksum = result.Checksum

	// log.Println(reqLog)
	log.Println(d)

	if b.Image != nil {
		image, err := s3.Upload(b.Image.URL.String())
		if err != nil {
			log.Println(err) // @TODO we might eventually better handle this case
		} else {
			b.Image = image
		}
	}

	err = repositories.Bookmarks.Upsert(ctx, b)
	if err != nil {
		return nil, err
	}

	if b.Image != nil {
		err = repositories.Bookmarks.UpdateImage(ctx, b)
		if err != nil {
			return nil, err
		}
	}

	err = repositories.Botlogs.Insert(ctx, result)
	if err != nil {
		return nil, err
	}

	err = repositories.Feeds.InsertAllIfNotExists(ctx, d.Feeds)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// CreateUserBookmark in this use case given a user and a bookmarkwe will from:
// - Add the bookmark to the user's bookmark collection
// - Save it in the DB
// - And finally return the user's bookmark
func CreateUserBookmark(ctx context.Context, user *user.User, bookmark *bookmark.Bookmark, repositories *repository.Repositories) (*bookmark.UserBookmark, error) {
	err := repositories.UserBookmarks.AddToUserCollection(ctx, user, bookmark)
	if err != nil {
		return nil, err
	}

	ub, err := repositories.UserBookmarks.GetByURL(ctx, user, bookmark.URL)
	if err != nil {
		return nil, err
	}

	return ub, nil
}
