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

	"github.com/go-sql-driver/mysql"
)

// BookmarkRepository the User Bookmark repository
type BookmarkRepository struct {
	db *sql.DB
}

// GetReadingList find latest entries
func (r *BookmarkRepository) GetReadingList(ctx context.Context, user *user.User, fromID string, toID string, limit int32) ([]*bookmark.Bookmark, error) {
	var results []*bookmark.Bookmark

	query := `
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.favorited_at, b.updated_at, b.linked, b.favorite
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE %s
		ORDER BY b.added_at DESC
		LIMIT ?
	`

	var where []string
	var args []interface{}

	fromDate, err := r.getCursorDate(ctx, fromID, "added_at")
	if err != nil {
		return nil, err
	}

	toDate, err := r.getCursorDate(ctx, toID, "added_at")
	if err != nil {
		return nil, err
	}

	if fromID != "" && toID != "" {
		where = append(where, "b.added_at <= ?")
		where = append(where, "b.added_at >= ?")
		where = append(where, "d.id NOT IN (?, ?)")
		args = append(args, fromDate)
		args = append(args, toDate)
		args = append(args, fromID)
		args = append(args, toID)
	} else if fromID != "" && toID == "" {
		where = append(where, "b.added_at <= ?")
		where = append(where, "d.id != ?")
		args = append(args, fromDate)
		args = append(args, fromID)
	} else if fromID == "" && toID != "" {
		where = append(where, "b.added_at >= ?")
		where = append(where, "d.id != ?")
		args = append(args, toDate)
		args = append(args, toID)
	}

	where = append(where, "d.deleted = 0")
	where = append(where, "b.linked = 1")
	where = append(where, "b.favorite = 0")
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
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.favorited_at, b.updated_at, b.linked, b.favorite
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE %s
		ORDER BY b.favorited_at DESC
		LIMIT ?
	`

	var where []string
	var args []interface{}

	fromDate, err := r.getCursorDate(ctx, fromID, "favorited_at")
	if err != nil {
		return nil, err
	}

	toDate, err := r.getCursorDate(ctx, toID, "favorited_at")
	if err != nil {
		return nil, err
	}

	if fromID != "" && toID != "" {
		where = append(where, "b.favorited_at <= ?")
		where = append(where, "b.favorited_at >= ?")
		where = append(where, "d.id NOT IN (?, ?)")
		args = append(args, fromDate)
		args = append(args, toDate)
		args = append(args, fromID)
		args = append(args, toID)
	} else if fromID != "" && toID == "" {
		where = append(where, "b.favorited_at <= ?")
		where = append(where, "d.id != ?")
		args = append(args, fromDate)
		args = append(args, fromID)
	} else if fromID == "" && toID != "" {
		where = append(where, "b.favorited_at >= ?")
		where = append(where, "d.id != ?")
		args = append(args, toDate)
		args = append(args, toID)
	}

	where = append(where, "d.deleted = 0")
	where = append(where, "b.linked = 1")
	where = append(where, "b.favorite = 1")
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

// CountFavorites count latest entries
func (r *BookmarkRepository) CountFavorites(ctx context.Context, user *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.favorite = 1 AND b.user_id = ?
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query), user.ID).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// CountReadingList count latest entries
func (r *BookmarkRepository) CountReadingList(ctx context.Context, user *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.deleted = 0 AND b.linked = 1 AND b.favorite = 0 AND b.user_id = ?
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
		SELECT d.id, d.url, d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, b.added_at, b.favorited_at, b.updated_at, b.linked, b.favorite
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

func (r *BookmarkRepository) getCursorDate(ctx context.Context, id string, date string) (t time.Time, err error) {
	query := `
		SELECT b.%s
		FROM documents AS d
		INNER JOIN bookmarks AS b ON b.document_id = d.id
		WHERE d.id = ?
	`
	t = time.Now()
	if id == "" {
		return t, err
	}

	err = r.db.QueryRowContext(ctx, formatQuery(fmt.Sprintf(query, date)), id).Scan(&t)
	if err != nil && err != sql.ErrNoRows {
		return time.Now(), err
	}

	return t, nil
}

// BookmarkDocument the bookmark to the user
func (r *BookmarkRepository) BookmarkDocument(ctx context.Context, u *user.User, d *document.Document, isFavorite bool) error {
	var favoritedAt mysql.NullTime
	if isFavorite {
		favoritedAt.Valid = true
		favoritedAt.Time = time.Now()
	}

	query := `
		INSERT INTO bookmarks
		(user_id, document_id, added_at, favorited_at, updated_at, favorite, linked)
		VALUES
		(?, ?, ?, ?, ?, ?, 1)
		ON DUPLICATE KEY UPDATE updated_at = ?, linked = 1, favorite = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.ID,
		d.ID,
		time.Now(),
		favoritedAt,
		time.Now(),
		isFavorite,
		time.Now(),
		isFavorite,
	)

	return err
}

// ChangeFavoriteStatus mark or unmark a bookmark as favorite
func (r *BookmarkRepository) ChangeFavoriteStatus(ctx context.Context, u *user.User, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET favorite = ?, favorited_at = ?, updated_at = ?
		WHERE user_id = ? AND document_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		b.IsFavorite,
		b.FavoritedAt,
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
		SET linked = ?, updated_at = ?
		WHERE user_id = ? AND document_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
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
	var favoritedAt mysql.NullTime

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
		&favoritedAt,
		&b.UpdatedAt,
		&b.IsLinked,
		&b.IsFavorite,
	)

	if favoritedAt.Valid {
		b.FavoritedAt = favoritedAt.Time
	}

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
