package repository

import (
	"github/mickaelvieira/taipan/internal/db"
)

// NewBookmarkRepository initialize repository
func NewBookmarkRepository() *BookmarkRepository {
	db := db.GetDB()

	var repository = BookmarkRepository{conn: db}

	return &repository
}
