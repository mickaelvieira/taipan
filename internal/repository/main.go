package repository

import "github/mickaelvieira/taipan/internal/db"

var newDB = db.GetDB()

// NewBookmarkRepository initialize repository
func NewBookmarkRepository() *BookmarkRepository {

	var repository = BookmarkRepository{db: newDB}

	return &repository
}
