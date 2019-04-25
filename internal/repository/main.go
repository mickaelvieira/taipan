package repository

import (
	"github/mickaelvieira/taipan/internal/db"
	"log"
)

// NewBookmarkRepository initialize repository
func NewBookmarkRepository() *BookmarkRepository {

	log.Println("New Repo")
	var repository = BookmarkRepository{db: db.GetDB()}

	return &repository
}
