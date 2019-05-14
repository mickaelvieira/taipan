package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"log"
	"strconv"
)

// FeedRepository the Feed repository
type FeedRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *FeedRepository) GetByID(ctx context.Context, id string) (*feed.Feed, error) {
	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE id = ?
	`
	rows := r.db.QueryRowContext(ctx, query, id)
	f, err := r.scan(rows)

	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetDocumentFeeds returns the document's feeds
func (r *FeedRepository) GetDocumentFeeds(ctx context.Context, d *document.Document) ([]*feed.Feed, error) {
	var results []*feed.Feed
	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE document_id = ?
	`
	rows, err := r.db.QueryContext(ctx, query, d.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

// GetNewFeeds returns the new created feed entries
func (r *FeedRepository) GetNewFeeds(ctx context.Context) ([]*feed.Feed, error) {
	var results []*feed.Feed
	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE status = ? AND url != "http://1001days.london/comments/feed/"
		LIMIT 1
	`
	rows, err := r.db.QueryContext(ctx, query, feed.NEW)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

// GetByURL find a single entry by URL and returns its ID
func (r *FeedRepository) GetByURL(ctx context.Context, u *uri.URI) (*feed.Feed, error) {
	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE url = ?
	`
	rows := r.db.QueryRowContext(ctx, query, u.String())
	f, err := r.scan(rows)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Insert creates a new feed in the DB
func (r *FeedRepository) Insert(ctx context.Context, f *feed.Feed, d *document.Document) error {
	query := `
		INSERT INTO feeds
		(document_id, url, title, type, status, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		d.ID,
		f.URL,
		f.Title,
		f.Type,
		f.Status,
		f.CreatedAt,
		f.UpdatedAt,
	)

	if err == nil {
		var ID int64
		ID, err = result.LastInsertId()
		if err == nil {
			f.ID = strconv.FormatInt(ID, 10)
		}
	}

	return err
}

// InsertIfNotExists stores the feed in the database if there is none with the same URL
func (r *FeedRepository) InsertIfNotExists(ctx context.Context, f *feed.Feed, d *document.Document) error {
	feed, err := r.GetByURL(ctx, f.URL)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.Insert(ctx, f, d)
		}
	} else {
		f.ID = feed.ID
	}
	return err
}

// InsertAllIfNotExists stores feeds in the database if there are none with the same URL
func (r *FeedRepository) InsertAllIfNotExists(ctx context.Context, feeds []*feed.Feed, d *document.Document) error {
	var err error
	for _, feed := range feeds {
		err = r.InsertIfNotExists(ctx, feed, d)
		if err != nil {
			break
		}
	}
	return err
}

func (r *FeedRepository) scan(rows Scanable) (*feed.Feed, error) {
	var feed feed.Feed
	err := rows.Scan(
		&feed.ID,
		&feed.URL,
		&feed.Title,
		&feed.Type,
		&feed.Status,
		&feed.CreatedAt,
		&feed.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &feed, nil
}
