package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/user"
	"log"
	"time"
)

// UserBookmarkRepository the User Bookmark repository
type UserBookmarkRepository struct {
	db *sql.DB
}

// FindLatest find latest entries
func (r *UserBookmarkRepository) FindLatest(ctx context.Context, user *user.User, cursor int32, limit int32) []*bookmark.Bookmark {
	var bookmarks []*bookmark.Bookmark

	query := "SELECT bookmarks.id, url, title, description, image_url, status, created_at, updated_at FROM bookmarks INNER JOIN users_bookmarks ON users_bookmarks.bookmark_id = bookmarks.id WHERE linked = 1 AND user_id = ? ORDER BY added_at DESC LIMIT ?, ?"
	rows, err := r.db.QueryContext(ctx, query, user.ID, cursor, limit)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.Image, &bookmark.Status, &bookmark.CreatedAt, &bookmark.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return bookmarks
}

// GetTotal count latest entries
func (r *UserBookmarkRepository) GetTotal(ctx context.Context, user *user.User) int32 {
	var total int32

	sql := "SELECT COUNT(bookmarks.id) as total FROM bookmarks INNER JOIN users_bookmarks ON users_bookmarks.bookmark_id = bookmarks.id WHERE linked = 1 AND user_id = ?"
	r.db.QueryRowContext(ctx, sql, user.ID).Scan(&total)

	return total
}

// IsLinked checked whether the bookmark is linked to the user
func (r *UserBookmarkRepository) IsLinked(ctx context.Context, user *user.User, b *bookmark.Bookmark) (string, int32) {
	var id string
	var linked int32

	query := "SELECT id, linked FROM users_bookmarks WHERE bookmark_id = ? AND user_id = ?"
	err := r.db.QueryRowContext(ctx, query, b.ID, user.ID).Scan(&id, &linked)

	if err != nil {
		if err == sql.ErrNoRows {
			return id, linked
		}
		log.Fatal(err)
	}

	return id, linked
}

// Link the bookmark to the user
func (r *UserBookmarkRepository) Link(ctx context.Context, user *user.User, b *bookmark.Bookmark) error {
	query := "INSERT INTO users_bookmarks (user_id, bookmark_id, added_at, accessed_at, marked_as_read, linked) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.ID, b.ID, time.Now(), time.Now(), 0, 1)

	return err
}

// ReLink the bookmark to the user
func (r *UserBookmarkRepository) ReLink(ctx context.Context, user *user.User, b *bookmark.Bookmark) error {
	query := "UPDATE users_bookmarks SET linked = 1 WHERE bookmark_id = ? AND user_id = ?"
	_, err := r.db.ExecContext(ctx, query, user.ID)

	return err
}
