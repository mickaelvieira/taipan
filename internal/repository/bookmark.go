package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"strings"
)

// BookmarkRepository the User Bookmark repository
type BookmarkRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *BookmarkRepository) GetByID(ctx context.Context, id string) *bookmark.Bookmark {
	var bookmark bookmark.Bookmark

	query := "SELECT id, url, charset, language, title, description, image_url, status, created_at, updated_at FROM bookmarks WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&bookmark.ID, &bookmark.URL, &bookmark.Charset, &bookmark.Lang, &bookmark.Title, &bookmark.Description, &bookmark.Image, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return &bookmark
		}
		log.Fatal(err)
	}

	return &bookmark
}

// GetByURL find a single entry
func (r *BookmarkRepository) GetByURL(ctx context.Context, URL string) string {
	var id string

	query := "SELECT id FROM bookmarks WHERE url = ?"
	err := r.db.QueryRowContext(ctx, query, URL).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return id
		}
		log.Fatal(err)
	}

	return id
}

// GetByIDs find all entries
func (r *BookmarkRepository) GetByIDs(ctx context.Context, ids []string) []*bookmark.Bookmark {
	var bookmarks []*bookmark.Bookmark

	params := make([]interface{}, len(ids))
	for i := range ids {
		params[i] = ids[i]
	}

	query := fmt.Sprintf("SELECT id, url, charset, language, title, description, image_url, status, created_at, updated_at FROM bookmarks WHERE id IN (?%s)", strings.Repeat(",?", len(ids)-1))
	rows, err := r.db.QueryContext(ctx, query, params...)

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Charset, &bookmark.Lang, &bookmark.Title, &bookmark.Description, &bookmark.Image, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return bookmarks
}

// Insert creates a new bookmark in the DB
func (r *BookmarkRepository) Insert(ctx context.Context, b *bookmark.Bookmark) *bookmark.Bookmark {
	query := "INSERT INTO bookmarks(id, url, charset, language, title, description, image_url, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, b.ID, b.URL, b.Charset, b.Lang, b.Title, b.Description, b.Image, b.Status, b.CreatedAt, b.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

// Update updates a bookmark in the DB
func (r *BookmarkRepository) Update(ctx context.Context, b *bookmark.Bookmark) *bookmark.Bookmark {
	query := "UPDATE bookmarks SET charset = ?, language = ?, title = ?, description = ?, image_url = ?, status = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, b.Charset, b.Lang, b.Title, b.Description, b.Image, b.Status, b.UpdatedAt, b.ID)

	if err != nil {
		log.Fatal(err)
	}

	return b
}
