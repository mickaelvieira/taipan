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
func (r *BookmarkRepository) GetByID(ctx context.Context, id string) (*bookmark.Bookmark, error) {
	var bookmark bookmark.Bookmark

	query := `
		SELECT id, url, charset, language, title, description, image_url, status, created_at, updated_at
		FROM bookmarks
		WHERE id = ?
	`
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(
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
		)

	if err != nil {
		return nil, err
	}

	return &bookmark, nil
}

// GetByURL find a single entry
func (r *BookmarkRepository) GetByURL(ctx context.Context, URL string) (*bookmark.Bookmark, error) {
	var bookmark bookmark.Bookmark

	query := `
		SELECT id, url, charset, language, title, description, image_url, status, created_at, updated_at
		FROM bookmarks
		WHERE url = ?
	`
	err := r.db.QueryRowContext(ctx, query, URL).
		Scan(
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
		)

	if err != nil {
		return nil, err
	}

	return &bookmark, nil
}

// GetByIDs find all entries
func (r *BookmarkRepository) GetByIDs(ctx context.Context, ids []string) ([]*bookmark.Bookmark, error) {
	var bookmarks []*bookmark.Bookmark

	params := make([]interface{}, len(ids))
	for i := range ids {
		params[i] = ids[i]
	}

	query := `
		SELECT id, url, charset, language, title, description, image_url, status, created_at, updated_at
		FROM bookmarks
		WHERE id IN (?%s)
	`
	query = fmt.Sprintf(query, strings.Repeat(",?", len(ids)-1))
	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookmark bookmark.Bookmark
		if err := rows.
			Scan(
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

// Insert creates a new bookmark in the DB
func (r *BookmarkRepository) Insert(ctx context.Context, b *bookmark.Bookmark) error {
	query := `
		INSERT INTO bookmarks
		(id, url, charset, language, title, description, image_url, status, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		b.ID,
		b.URL,
		b.Charset,
		b.Lang,
		b.Title,
		b.Description,
		b.Image,
		b.Status,
		b.CreatedAt,
		b.UpdatedAt,
	)

	return err
}

// Update updates a bookmark in the DB
func (r *BookmarkRepository) Update(ctx context.Context, b *bookmark.Bookmark) error {
	query := `
		UPDATE bookmarks
		SET charset = ?, language = ?, title = ?, description = ?, image_url = ?, status = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		b.Charset,
		b.Lang,
		b.Title,
		b.Description,
		b.Image,
		b.Status,
		b.UpdatedAt,
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
