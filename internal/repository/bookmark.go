package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/types"
	"github/mickaelvieira/taipan/internal/domain/user"
	"strconv"
	"strings"
)

// BookmarkRepository the User Bookmark repository
type BookmarkRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *BookmarkRepository) GetByID(ctx context.Context, id string) (*bookmark.Bookmark, error) {
	query := `
		SELECT b.id, b.url, HEX(b.checksum), b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, b.created_at, b.updated_at
		FROM bookmarks AS b
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)
	bookmark, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

// GetByURL find a single entry
func (r *BookmarkRepository) GetByURL(ctx context.Context, u *types.URI) (*bookmark.Bookmark, error) {
	query := `
		SELECT b.id, b.url, HEX(b.checksum), b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, b.created_at, b.updated_at
		FROM bookmarks AS b
		WHERE url = ?
	`
	row := r.db.QueryRowContext(ctx, query, u.String())
	bookmark, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

// GetByChecksum find a single entry
func (r *BookmarkRepository) GetByChecksum(ctx context.Context, c types.Checksum) (*bookmark.Bookmark, error) {
	query := `
		SELECT b.id, b.url, HEX(b.checksum), b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, b.created_at, b.updated_at
		FROM bookmarks AS b
		WHERE b.checksum = UNHEX(?)
	`
	row := r.db.QueryRowContext(ctx, query, c.String())
	bookmark, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

// GetByIDs find all entries
func (r *BookmarkRepository) GetByIDs(ctx context.Context, ids []string) ([]*bookmark.Bookmark, error) {
	var bookmarks []*bookmark.Bookmark

	params := make([]interface{}, len(ids))
	for i := range ids {
		params[i] = ids[i]
	}

	query := `
		SELECT b.id, b.url, HEX(b.checksum), b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, b.created_at, b.updated_at
		FROM bookmarks AS b
		WHERE id IN (?%s)
	`
	query = fmt.Sprintf(query, strings.Repeat(",?", len(ids)-1))
	rows, err := r.db.QueryContext(ctx, query, params...)
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

// FindNew find newest entries
func (r *BookmarkRepository) FindNew(ctx context.Context, user *user.User, cursor int32, limit int32) ([]*bookmark.Bookmark, error) {
	var bookmarks []*bookmark.Bookmark

	query := `
		SELECT b.id, b.url, HEX(b.checksum), b.charset, b.language, b.title, b.description, b.image_url, b.image_name, b.image_width, b.image_height, b.image_format, b.created_at, b.updated_at
		FROM bookmarks AS b
		LEFT JOIN users_bookmarks AS ub ON ub.bookmark_id = b.id
		WHERE ub.user_id IS NULL OR ub.user_id != ?
		ORDER BY b.updated_at DESC
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

// GetTotal count latest entries
func (r *BookmarkRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(b.id) as total FROM bookmarks AS b
	`
	err := r.db.QueryRowContext(ctx, query).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// Insert creates a new bookmark in the DB
func (r *BookmarkRepository) Insert(ctx context.Context, b *bookmark.Bookmark) error {
	query := `
		INSERT INTO bookmarks
		(url, checksum, charset, language, title, description, status, created_at, updated_at)
		VALUES
		(?, UNHEX(?), ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(
		ctx,
		query,
		b.URL,
		b.Checksum,
		b.Charset,
		b.Lang,
		b.Title,
		b.Description,
		b.Status,
		b.CreatedAt,
		b.UpdatedAt,
	)

	if err == nil {
		var ID int64
		ID, err = result.LastInsertId()
		if err == nil {
			b.ID = strconv.FormatInt(ID, 10)
		}
	}

	return err
}

// Update updates a bookmark in the DB
func (r *BookmarkRepository) Update(ctx context.Context, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET checksum = UNHEX(?), charset = ?, language = ?, title = ?, description = ?, status = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		b.Checksum,
		b.Charset,
		b.Lang,
		b.Title,
		b.Description,
		b.Status,
		b.UpdatedAt,
		b.ID,
	)

	return err
}

// UpdateImage updates a bookmark's image in the DB
func (r *BookmarkRepository) UpdateImage(ctx context.Context, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET image_url = ?, image_name = ?, image_width = ?, image_height = ?, image_format = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		b.Image.URL.String(),
		b.Image.Name,
		b.Image.Width,
		b.Image.Height,
		b.Image.Format,
		b.ID,
	)

	return err
}

// Upsert insert the bookmark or update if there is already one with the same URL
func (r *BookmarkRepository) Upsert(ctx context.Context, b *bookmark.Bookmark) error {
	bookmark, err := r.GetByURL(ctx, b.URL)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return r.Insert(ctx, b)
	}

	b.ID = bookmark.ID

	return r.Update(ctx, b)
}

func (r *BookmarkRepository) scan(rows Scanable) (*bookmark.Bookmark, error) {
	var bookmark bookmark.Bookmark
	var imageURL, imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&bookmark.ID,
		&bookmark.URL,
		&bookmark.Checksum,
		&bookmark.Charset,
		&bookmark.Lang,
		&bookmark.Title,
		&bookmark.Description,
		&imageURL,
		&imageName,
		&imageWidth,
		&imageHeight,
		&imageFormat,
		&bookmark.CreatedAt,
		&bookmark.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if imageURL != "" {
		image, err := getBookmarkImage(imageURL, imageName, imageWidth, imageHeight, imageFormat)
		if err != nil {
			return nil, err
		}
		bookmark.Image = image
	}

	return &bookmark, nil
}
