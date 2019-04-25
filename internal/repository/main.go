package repository

import (
	"github/mickaelvieira/taipan/internal/db"
)

// Repositories holds a reference to the repositories
type Repositories struct {
	Users         *UserRepository
	Bookmarks     *BookmarkRepository
	UserBookmarks *UserBookmarkRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:         &UserRepository{db: db},
		Bookmarks:     &BookmarkRepository{db: db},
		UserBookmarks: &UserBookmarkRepository{db: db},
	}
}
