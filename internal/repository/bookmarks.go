package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"time"
)

// BookmarkRepository the User Bookmark repository
type BookmarkRepository struct {
	db *sql.DB
}

// GetReadingList find latest entries
func (r *BookmarkRepository) GetReadingList(ctx context.Context, user *user.User, cursor int32, limit int32) ([]*bookmark.Bookmark, error) {
	var results []*bookmark.Bookmark

	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.updated_at, b.linked, b.marked_as_read
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.marked_as_read = 0 AND b.user_id = ?
		ORDER BY b.updated_at DESC
		LIMIT ?, ?
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), user.ID, cursor, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, b)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetFavorites find latest entries
func (r *BookmarkRepository) GetFavorites(ctx context.Context, user *user.User, cursor int32, limit int32) ([]*bookmark.Bookmark, error) {
	var results []*bookmark.Bookmark

	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.updated_at, b.linked, b.marked_as_read
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.marked_as_read = 1 AND b.user_id = ?
		ORDER BY b.updated_at DESC
		LIMIT ?, ?
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), user.ID, cursor, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, b)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetTotalFavorites count latest entries
func (r *BookmarkRepository) GetTotalFavorites(ctx context.Context, user *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.marked_as_read = 1 AND b.user_id = ?
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query), user.ID).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetTotalReadingList count latest entries
func (r *BookmarkRepository) GetTotalReadingList(ctx context.Context, user *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.marked_as_read = 0 AND b.user_id = ?
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query), user.ID).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry
func (r *BookmarkRepository) GetByURL(ctx context.Context, user *user.User, u *url.URL) (*bookmark.Bookmark, error) {
	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.updated_at, b.linked, b.marked_as_read
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE b.user_id = ? AND d.url = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), user.ID, u.UnescapeString())
	b, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// BookmarkDocument the bookmark to the user
func (r *BookmarkRepository) BookmarkDocument(ctx context.Context, user *user.User, d *document.Document) error {
	query := `
		INSERT INTO bookmarks
		(user_id, document_id, added_at, updated_at, marked_as_read, linked)
		VALUES
		(?, ?, ?, ?, 0, 1)
		ON DUPLICATE KEY UPDATE updated_at = ?, linked = 1
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		user.ID,
		d.ID,
		time.Now(),
		time.Now(),
		time.Now(),
	)

	return err
}

// ChangeReadStatus change bookmarks read status .i.e READ/UNREAD
func (r *BookmarkRepository) ChangeReadStatus(ctx context.Context, user *user.User, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET updated_at = ?, marked_as_read = %s
		WHERE user_id = ? AND document_id = ?
	`

	// Workaround to overcome MySQL driver limitation
	isRead := "0"
	if b.IsRead == bookmark.READ {
		isRead = "1"
	}
	query = fmt.Sprintf(query, isRead)

	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		b.UpdatedAt,
		user.ID,
		b.ID,
	)

	return err
}

// Remove bookmarks from user list
func (r *BookmarkRepository) Remove(ctx context.Context, user *user.User, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET updated_at = ?, linked = ?
		WHERE user_id = ? AND document_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		b.UpdatedAt,
		b.IsLinked,
		user.ID,
		b.ID,
	)

	return err
}

func (r *BookmarkRepository) scan(rows Scanable) (*bookmark.Bookmark, error) {
	var b bookmark.Bookmark
	var imageURL, imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&b.ID,
		&b.URL,
		&b.Charset,
		&b.Lang,
		&b.Title,
		&b.Description,
		&imageURL,
		&imageName,
		&imageWidth,
		&imageHeight,
		&imageFormat,
		&b.AddedAt,
		&b.UpdatedAt,
		&b.IsLinked,
		&b.IsRead,
	)

	if err != nil {
		return nil, err
	}

	if imageURL != "" {
		image, err := getImageEntity(imageURL, imageName, imageWidth, imageHeight, imageFormat)
		if err != nil {
			return nil, err
		}
		b.Image = image
	}

	return &b, nil
}
