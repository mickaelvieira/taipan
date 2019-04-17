package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"strings"
)

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

	params := make([]interface{}, len(ids))
	for i := range ids {
		params[i] = ids[i]
	}

	sql := "SELECT id, url, title, description, hash, status, created_at, updated_at FROM bookmarks WHERE id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"
	rows, err := r.conn.QueryContext(ctx, sql, params...)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.Hash, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	return bookmarks
}

// FindLatest find latest entries
func (r *BookmarkRepository) FindLatest(ctx context.Context, cursor int32, limit int32) []string {
	var ids []string

	rows, err := r.conn.QueryContext(ctx, "SELECT bookmarks.id FROM bookmarks INNER JOIN users_bookmarks ON users_bookmarks.bookmark_id = bookmarks.id WHERE linked = 1 ORDER BY added_at DESC LIMIT ?, ?;", cursor, limit)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	return ids
}
