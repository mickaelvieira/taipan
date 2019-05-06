package repository

import (
	"github/mickaelvieira/taipan/internal/db"
)

// Scanable sql.Rows or sql.Row
type Scanable interface {
	Scan(...interface{}) error
}

// Repositories holds a reference to the repositories
type Repositories struct {
	Users         *UserRepository
	Feeds         *FeedRepository
	Bookmarks     *BookmarkRepository
	UserBookmarks *UserBookmarkRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:         &UserRepository{db: db},
		Feeds:         &FeedRepository{db: db},
		Bookmarks:     &BookmarkRepository{db: db},
		UserBookmarks: &UserBookmarkRepository{db: db},
	}
}
