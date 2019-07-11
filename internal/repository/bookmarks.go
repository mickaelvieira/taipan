package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"strings"
	"time"
)

// BookmarkRepository the User Bookmark repository
type BookmarkRepository struct {
	db *sql.DB
}

// GetReadingList find latest entries
func (r *BookmarkRepository) GetReadingList(ctx context.Context, user *user.User, fromID string, toID string, limit int32) ([]*bookmark.Bookmark, error) {
	var results []*bookmark.Bookmark

	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.updated_at, b.linked, b.marked_as_read
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE %s
		ORDER BY b.updated_at DESC
		LIMIT ?
	`
	where, args, err := r.getPagination(ctx, fromID, toID)
	if err != nil {
		return nil, err
	}

	where = append(where, "d.deleted = 0")
	where = append(where, "b.linked = 1")
	where = append(where, "b.marked_as_read = 0")
	where = append(where, "b.user_id = ?")
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

	args = append(args, user.ID)
	args = append(args, limit)
	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
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
func (r *BookmarkRepository) GetFavorites(ctx context.Context, user *user.User, fromID string, toID string, limit int32) ([]*bookmark.Bookmark, error) {
	var results []*bookmark.Bookmark

	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.updated_at, b.linked, b.marked_as_read
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE %s
		ORDER BY b.updated_at DESC
		LIMIT ?
	`
	where, args, err := r.getPagination(ctx, fromID, toID)
	if err != nil {
		return nil, err
	}

	where = append(where, "d.deleted = 0")
	where = append(where, "b.linked = 1")
	where = append(where, "b.marked_as_read = 1")
	where = append(where, "b.user_id = ?")
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

	args = append(args, user.ID)
	args = append(args, limit)
	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
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

func (r *BookmarkRepository) getCursorDate(ctx context.Context, id string) (t time.Time, err error) {
	query := `
		SELECT b.updated_at
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.id = ?
	`
	t = time.Now()
	if id == "" {
		return t, err
	}

	err = r.db.QueryRowContext(ctx, formatQuery(query), id).Scan(&t)
	if err != nil && err != sql.ErrNoRows {
		return time.Now(), err
	}

	return t, nil
}

func (r *BookmarkRepository) getPagination(ctx context.Context, fromID string, toID string) (where []string, args []interface{}, err error) {
	var fromDate, toDate time.Time
	fromDate, err = r.getCursorDate(ctx, fromID)
	if err != nil {
		return
	}

	toDate, err = r.getCursorDate(ctx, toID)
	if err != nil {
		return
	}

	var query string
	if fromID != "" && toID != "" {
		query = "b.updated_at <= ? AND b.updated_at >= ? AND d.id NOT IN (?, ?)"
		args = append(args, fromDate)
		args = append(args, toDate)
		args = append(args, fromID)
		args = append(args, toID)
	} else if fromID != "" && toID == "" {
		query = "b.updated_at <= ? AND d.id != ?"
		args = append(args, fromDate)
		args = append(args, fromID)
	} else if fromID == "" && toID != "" {
		query = "b.updated_at >= ? AND d.id != ?"
		args = append(args, toDate)
		args = append(args, toID)
	} else {
		return
	}

	where = append(where, query)

	return
}

// BookmarkDocument the bookmark to the user
func (r *BookmarkRepository) BookmarkDocument(ctx context.Context, u *user.User, d *document.Document, isFavorite bool) error {
	query := `
		INSERT INTO bookmarks
		(user_id, document_id, added_at, updated_at, marked_as_read, linked)
		VALUES
		(?, ?, ?, ?, ?, 1)
		ON DUPLICATE KEY UPDATE updated_at = ?, linked = 1
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.ID,
		d.ID,
		time.Now(),
		time.Now(),
		isFavorite,
		time.Now(),
	)

	return err
}

// ChangeFavoriteStatus change bookmarks read status .i.e READ/UNREAD
func (r *BookmarkRepository) ChangeFavoriteStatus(ctx context.Context, u *user.User, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET marked_as_read = ?, updated_at = ?
		WHERE user_id = ? AND document_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		b.IsFavorite,
		b.UpdatedAt,
		u.ID,
		b.ID,
	)

	return err
}

// Remove bookmarks from user list
func (r *BookmarkRepository) Remove(ctx context.Context, u *user.User, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET marked_as_read = ?, linked = ?, updated_at = ?
		WHERE user_id = ? AND document_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		b.IsFavorite,
		b.IsLinked,
		b.UpdatedAt,
		u.ID,
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
		&b.IsFavorite,
	)

	if err != nil {
		return nil, err
	}

	if imageURL != "" {
		i, err := document.NewImage(imageURL, imageName, imageWidth, imageHeight, imageFormat)
		if err != nil {
			return nil, err
		}
		b.Image = i
	}

	return &b, nil
}
