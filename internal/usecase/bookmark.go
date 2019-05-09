package usecase

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/helpers"
	"github/mickaelvieira/taipan/internal/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
	"io"
	"log"
	"net/url"
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
		return nil, errors.New("Invalid URL")
	}

	URL = helpers.RemoveFragment(URL)

	var reader io.Reader
	var result *client.Result
	result, reader, err = cl.Fetch(URL)
	if err != nil {
		return nil, err
	}

	// @TODO Don't parse the document if it hasn't changed
	// @TODO I need to check client's result before parsing
	var document *parser.Document
	document, err = parser.Parse(URL, reader)
	if err != nil {
		return nil, err
	}

	bookmark := document.ToBookmark()

	// log.Println(reqLog)
	log.Println(document)

	if bookmark.Image != nil {
		image, err := s3.Upload(bookmark.Image.URL.String())
		if err != nil {
			log.Println(err) // @TODO we might eventually better handle this case
		} else {
			bookmark.Image = image
		}
	}

	err = repositories.Bookmarks.Upsert(ctx, bookmark)
	if err != nil {
		return nil, err
	}

	if bookmark.Image != nil {
		err = repositories.Bookmarks.UpdateImage(ctx, bookmark)
		if err != nil {
			return nil, err
		}
	}

	err = repositories.Botlogs.Insert(ctx, result)
	if err != nil {
		return nil, err
	}

	err = repositories.Feeds.InsertAllIfNotExists(ctx, document.Feeds)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
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

	userBookmark, err := repositories.UserBookmarks.GetByURL(ctx, user, bookmark.URL)
	if err != nil {
		return nil, err
	}

	return userBookmark, nil
}
