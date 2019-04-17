package repository

import "github/mickaelvieira/taipan/internal/db"

var conn = db.GetDB()

// NewBookmarkRepository initialize repository
func NewBookmarkRepository() *BookmarkRepository {

	var repository = BookmarkRepository{conn: conn}

	return &repository
}
