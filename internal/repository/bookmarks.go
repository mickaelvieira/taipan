package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"strings"
)

// NewBookmarkRepository initialize repository
func NewBookmarkRepository() *BookmarkRepository {
	db := db.GetDB()

	var repository = BookmarkRepository{conn: db}

	return &repository
}

// BookmarkRepository the GPX repository
type BookmarkRepository struct {
	conn *sql.DB
}

// FindOne find a single entry
func (r *BookmarkRepository) FindOne(ctx context.Context, id string) *bookmark.Bookmark {
	var bookmark bookmark.Bookmark

	err := r.conn.QueryRowContext(ctx, "SELECT id, title, description, hash FROM bookmarks WHERE id = ?", id).Scan(&bookmark.ID, &bookmark.Title, &bookmark.Description, &bookmark.Hash)

	if err != nil {
		log.Fatal(err)
	}

	return &bookmark
}

// FindAll find all entries
func (r *BookmarkRepository) FindAll(ctx context.Context, ids []string) []*bookmark.Bookmark {
	var bookmarks []*bookmark.Bookmark

	// @TODO need to double whether this is safe
	rows, err := r.conn.QueryContext(ctx, "SELECT id, title, description, hash, status, created_at, updated_at FROM bookmarks WHERE id IN (?);", strings.Join(ids, ","))

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.Title, &bookmark.Description, &bookmark.Hash, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	return bookmarks
}
