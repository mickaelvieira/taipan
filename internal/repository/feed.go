package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"log"
	"strconv"
)

// FeedRepository the Feed repository
type FeedRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *FeedRepository) GetByID(ctx context.Context, id string) (*feed.Feed, error) {
	var feed feed.Feed

	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE id = ?
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
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

// GetNewFeeds returns the new created feed entries
func (r *FeedRepository) GetNewFeeds(ctx context.Context) ([]*feed.Feed, error) {
	var feeds []*feed.Feed

	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE status = ?
	`
	rows, err := r.db.QueryContext(ctx, query, feed.NEW)
	if err != nil {
		return nil, err
	}

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
			break
		}
		feeds = append(feeds, &feed)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return feeds, nil
}

// GetByURL find a single entry by URL and returns its ID
func (r *FeedRepository) GetByURL(ctx context.Context, URL string) (*feed.Feed, error) {
	var feed feed.Feed

	query := `
		SELECT id, url, title, type, status, created_at, updated_at
		FROM feeds
		WHERE url = ?
	`
	err := r.db.QueryRowContext(ctx, query, URL).Scan(
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

// Insert creates a new feed in the DB
func (r *FeedRepository) Insert(ctx context.Context, f *feed.Feed) error {
	query := `
		INSERT INTO feeds
		(url, title, type, status, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
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
func (r *FeedRepository) InsertIfNotExists(ctx context.Context, f *feed.Feed) error {
	feed, err := r.GetByURL(ctx, f.URL)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.Insert(ctx, f)
		}
	} else {
		f.ID = feed.ID
	}
	return err
}

// InsertAllIfNotExists stores feeds in the database if there are none with the same URL
func (r *FeedRepository) InsertAllIfNotExists(ctx context.Context, feeds []*feed.Feed) error {
	var err error
	for _, feed := range feeds {
		err = r.InsertIfNotExists(ctx, feed)
		if err != nil {
			break
		}
	}
	return err
}
