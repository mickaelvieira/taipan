package repository

import (
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"net/url"
)

// Scanable sql.Rows or sql.Row
type Scanable interface {
	Scan(...interface{}) error
}

// Repositories holds a reference to the repositories
type Repositories struct {
	Users     *UserRepository
	Feeds     *FeedRepository
	Documents *DocumentRepository
	Bookmarks *BookmarkRepository
	Botlogs   *BotlogRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:     &UserRepository{db: db},
		Feeds:     &FeedRepository{db: db},
		Documents: &DocumentRepository{db: db},
		Bookmarks: &BookmarkRepository{db: db},
		Botlogs:   &BotlogRepository{db: db},
	}
}

func getImageEntity(rawURL string, name string, width int32, height int32, format string) (*image.Image, error) {
	URL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	var image = image.Image{
		URL:    &uri.URI{URL: URL},
		Name:   name,
		Width:  width,
		Height: height,
		Format: format,
	}

	return &image, nil
}
