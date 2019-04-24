package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"strings"
	"time"
)

var userID = "c1479a73-2f8a-11e8-ade8-fa163ea9b6ed"
var fields = []string{"id", "url", "title", "description", "status", "created_at", "updated_at"}

// BookmarkRepository the Bookmark repository
type BookmarkRepository struct {
	conn *sql.DB
}

// GetByID find a single entry
func (r *BookmarkRepository) GetByID(ctx context.Context, id string) *bookmark.Bookmark {
	var bookmark bookmark.Bookmark

	query := fmt.Sprintf("SELECT %s FROM bookmarks WHERE id = ?", strings.Join(fields, ", "))
	err := r.conn.QueryRowContext(ctx, query, id).Scan(&bookmark.ID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &bookmark
}

// GetByURL find a single entry
func (r *BookmarkRepository) GetByURL(ctx context.Context, URL string) string {
	var id string

	query := "SELECT id FROM bookmarks WHERE url = ?"
	err := r.conn.QueryRowContext(ctx, query, URL).Scan(&id)

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

	query := fmt.Sprintf("SELECT %s FROM bookmarks WHERE id IN (?%s)", strings.Join(fields, ", "), strings.Repeat(",?", len(ids)-1))
	rows, err := r.conn.QueryContext(ctx, query, params...)

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return bookmarks
}

// FindLatest find latest entries
func (r *BookmarkRepository) FindLatest(ctx context.Context, cursor int32, limit int32) []string {
	var ids []string

	query := "SELECT bookmarks.id FROM bookmarks INNER JOIN users_bookmarks ON users_bookmarks.bookmark_id = bookmarks.id WHERE linked = 1 AND user_id = ? ORDER BY added_at DESC LIMIT ?, ?"
	rows, err := r.conn.QueryContext(ctx, query, userID, cursor, limit)

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

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return ids
}

// GetTotal count latest entries
func (r *BookmarkRepository) GetTotal(ctx context.Context) int32 {
	var total int32

	sql := "SELECT COUNT(bookmarks.id) as total FROM bookmarks INNER JOIN users_bookmarks ON users_bookmarks.bookmark_id = bookmarks.id WHERE linked = 1 AND user_id = ?"
	r.conn.QueryRowContext(ctx, sql, userID).Scan(&total)

	return total
}

// Insert creates a new bookmark in the DB
func (r *BookmarkRepository) Insert(ctx context.Context, b *bookmark.Bookmark) *bookmark.Bookmark {
	stmt, err := r.conn.Prepare("INSERT INTO bookmarks(id, url, title, description, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(b.ID, b.URL, b.Title, b.Description, b.Status, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// Update updates a bookmark in the DB
func (r *BookmarkRepository) Update(ctx context.Context, b *bookmark.Bookmark) *bookmark.Bookmark {
	stmt, err := r.conn.Prepare("UPDATE bookmarks SET title = ?, description = ?, status = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(b.Title, b.Description, b.Status, b.UpdatedAt, b.ID)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// IsLinked checked whether the bookmark is linked to the user
func (r *BookmarkRepository) IsLinked(ctx context.Context, b *bookmark.Bookmark) (string, int32) {
	var id string
	var linked int32

	query := "SELECT id, linked FROM users_bookmarks WHERE bookmark_id = ? AND user_id = ?"
	err := r.conn.QueryRowContext(ctx, query, b.ID, userID).Scan(&id, &linked)

	if err != nil {
		if err == sql.ErrNoRows {
			return id, linked
		}
		log.Fatal(err)
	}

	return id, linked
}

// Link the bookmark to the user
func (r *BookmarkRepository) Link(ctx context.Context, b *bookmark.Bookmark) error {
	stmt, err := r.conn.Prepare("INSERT INTO users_bookmarks (user_id, bookmark_id, added_at, accessed_at, marked_as_read, linked) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, b.ID, time.Now(), time.Now(), 0, 1)

	return err
}

// ReLink the bookmark to the user
func (r *BookmarkRepository) ReLink(ctx context.Context, b *bookmark.Bookmark) error {
	stmt, err := r.conn.Prepare("UPDATE users_bookmarks SET linked = 1 WHERE bookmark_id = ? AND user_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.ID, userID)

	return err
}
