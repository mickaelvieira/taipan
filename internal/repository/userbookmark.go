package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/user"
	"log"
	"net/url"
	"time"
)

type row struct {
	id          string
	url         string
	lang        string
	charset     string
	title       string
	description string
	imageURL    string
	imageName   string
	imageWidth  int32
	imageHeight int32
	imageFormat string
	addedAt     time.Time
	updatedAt   time.Time
	isRead      bool
	isLinked    bool
}

// UserBookmarkRepository the User Bookmark repository
type UserBookmarkRepository struct {
	db *sql.DB
}

// FindNew find newest entries
func (r *UserBookmarkRepository) FindNew(ctx context.Context, user *user.User, cursor int32, limit int32) ([]*bookmark.Bookmark, error) {
	var bookmarks []*bookmark.Bookmark

	query := `
		SELECT b.id, user_id, url, charset, language, title, description, image_url, status, created_at, updated_at
		FROM bookmarks AS b
		LEFT JOIN users_bookmarks ON ub.bookmark_id = b.id
		WHERE ub.user_id IS NULL OR ub.user_id != ?
		ORDER BY b.created_at_at DESC
		LIMIT ?, ?
	`
	rows, err := r.db.QueryContext(ctx, query, user.ID, cursor, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.Scan(
			&bookmark.ID,
			&bookmark.URL,
			&bookmark.Charset,
			&bookmark.Lang,
			&bookmark.Title,
			&bookmark.Description,
			&bookmark.Image,
			&bookmark.Status,
			&bookmark.CreatedAt,
			&bookmark.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

// FindLatest find latest entries
func (r *UserBookmarkRepository) FindLatest(ctx context.Context, user *user.User, cursor int32, limit int32) ([]*bookmark.UserBookmark, error) {
	var bookmarks []*bookmark.UserBookmark

	query := `
		SELECT b.id, b.url, b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, ub.added_at, ub.updated_at, ub.linked, ub.marked_as_read
		FROM bookmarks AS b
		INNER JOIN users_bookmarks AS ub ON ub.bookmark_id = b.id
		WHERE ub.linked = 1 AND ub.user_id = ?
		ORDER BY ub.updated_at DESC
		LIMIT ?, ?
	`
	rows, err := r.db.QueryContext(ctx, query, user.ID, cursor, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		bookmark, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

// GetTotal count latest entries
func (r *UserBookmarkRepository) GetTotal(ctx context.Context, user *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(b.id) as total
		FROM bookmarks AS b
		INNER JOIN users_bookmarks AS ub ON ub.bookmark_id = b.id
		WHERE ub.linked = 1 AND ub.user_id = ?
	`
	err := r.db.QueryRowContext(ctx, query, user.ID).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry
func (r *UserBookmarkRepository) GetByURL(ctx context.Context, user *user.User, URL string) (*bookmark.UserBookmark, error) {
	query := `
		SELECT b.id, b.url, b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, ub.added_at, ub.updated_at, ub.linked, ub.marked_as_read
		FROM bookmarks AS b
		INNER JOIN users_bookmarks AS ub ON ub.bookmark_id = b.id
		WHERE ub.user_id = ? AND b.url = ?
	`
	row := r.db.QueryRowContext(ctx, query, user.ID, URL)
	bookmark, err := r.scan(row)

	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

// AddToUserCollection the bookmark to the user
func (r *UserBookmarkRepository) AddToUserCollection(ctx context.Context, user *user.User, b *bookmark.Bookmark) error {
	query := `
		INSERT INTO users_bookmarks
		(user_id, bookmark_id, added_at, updated_at, marked_as_read, linked)
		VALUES
		(?, ?, ?, ?, 0, 1)
		ON DUPLICATE KEY UPDATE updated_at = ?, linked = 1
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		b.ID,
		time.Now(),
		time.Now(),
		time.Now(),
	)

	return err
}

func (r *UserBookmarkRepository) scan(rows Scanable) (*bookmark.UserBookmark, error) {
	var rw row

	err := rows.Scan(
		&rw.id,
		&rw.url,
		&rw.charset,
		&rw.lang,
		&rw.title,
		&rw.description,
		&rw.imageURL,
		&rw.imageName,
		&rw.imageWidth,
		&rw.imageHeight,
		&rw.imageFormat,
		&rw.addedAt,
		&rw.updatedAt,
		&rw.isLinked,
		&rw.isRead,
	)

	if err != nil {
		return nil, err
	}

	b := bookmark.UserBookmark{
		ID:          rw.id,
		URL:         rw.url,
		Charset:     rw.charset,
		Lang:        rw.lang,
		Title:       rw.title,
		Description: rw.description,
		AddedAt:     rw.addedAt,
		UpdatedAt:   rw.updatedAt,
		IsLinked:    rw.isLinked,
		IsRead:      rw.isRead,
	}

	if rw.imageURL != "" {
		u, err := url.ParseRequestURI(rw.imageURL)

		if err != nil {
			return nil, err
		}

		i := bookmark.Image{
			URL:    u,
			Name:   rw.imageName,
			Width:  rw.imageWidth,
			Height: rw.imageHeight,
			Format: rw.imageFormat,
		}

		b.Image = &i
	}

	return &b, nil
}
