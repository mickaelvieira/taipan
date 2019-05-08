package usecase

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/parser"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
	"io"
	"log"
	"net/url"
)

// Bookmark bookmark use case
func Bookmark(ctx context.Context, rawURL string, repositories *repository.Repositories) (*bookmark.Bookmark, error) {
	cl := client.Client{}
	URL, err := url.ParseRequestURI(rawURL)
	if err != nil || !URL.IsAbs() {
		return nil, errors.New("Invalid URL")
	}

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
