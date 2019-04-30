package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"log"
)

// FeedRepository the Feed repository
type FeedRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *FeedRepository) GetByID(ctx context.Context, id string) *feed.Feed {
	var feed feed.Feed

	query := `#
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE id = ?
		`
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(
			&feed.ID,
			&feed.URL,
			&feed.Title,
			&feed.Type,
			&feed.Status,
			&feed.CreatedAt,
			&feed.UpdatedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return &feed
		}
		log.Fatal(err)
	}

	return &feed
}

// GetNewFeeds returns the new created feed entries
func (r *FeedRepository) GetNewFeeds(ctx context.Context) []*feed.Feed {
	var feeds []*feed.Feed

	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE status = ?
		`
	rows, err := r.db.QueryContext(ctx, query, feed.NEW)

	for rows.Next() {
		var feed feed.Feed
		if err := rows.Scan(
			&feed.ID,
			&feed.URL,
			&feed.Title,
			&feed.Type,
			&feed.Status,
			&feed.CreatedAt,
			&feed.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}
		feeds = append(feeds, &feed)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return feeds
}

// GetByURL find a single entry by URL and returns its ID
func (r *FeedRepository) GetByURL(ctx context.Context, URL string) string {
	var id string

	query := "SELECT id FROM feeds WHERE url = ?"
	err := r.db.QueryRowContext(ctx, query, URL).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return id
		}
		log.Fatal(err)
	}

	return id
}

// Insert creates a new feed in the DB
func (r *FeedRepository) Insert(ctx context.Context, f *feed.Feed) *feed.Feed {
	query := `
		INSERT INTO feeds
		(id, url, title, type, status, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)
		`
	_, err := r.db.ExecContext(
		ctx,
		query,
		f.ID,
		f.URL,
		f.Title,
		f.Type,
		f.Status,
		f.CreatedAt,
		f.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
