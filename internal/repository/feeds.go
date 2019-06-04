package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/url"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

// @TODO add the ability to soft delete a feed

// FeedRepository the Feed repository
type FeedRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *FeedRepository) GetByID(ctx context.Context, id string) (*feed.Feed, error) {
	query := `
		SELECT f.id, f.url, f.title, f.type, f.status, f.created_at, f.updated_at, f.parsed_at, f.deleted
		FROM feeds AS f
		WHERE f.id = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), id)
	f, err := r.scan(rows)

	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetOutdatedFeeds returns the feeds which have been last updated more than 24 hrs
func (r *FeedRepository) GetOutdatedFeeds(ctx context.Context) ([]*feed.Feed, error) {
	var results []*feed.Feed
	query := `
		SELECT f.id, f.url, f.title, f.type, f.status, f.created_at, f.updated_at, f.parsed_at, f.deleted
		FROM feeds AS f
		WHERE f.deleted = 0 AND (f.parsed_at IS NULL OR f.parsed_at < DATE_SUB(NOW(), INTERVAL 1 HOUR))
		LIMIT ?;
		`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), 10)
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

// FindAll find newest entries
func (r *FeedRepository) FindAll(ctx context.Context, cursor int32, limit int32) ([]*feed.Feed, error) {
	var results []*feed.Feed

	query := `
		SELECT f.id, f.url, f.title, f.type, f.status, f.created_at, f.updated_at, f.parsed_at, f.deleted
		FROM feeds AS f
		ORDER BY f.updated_at DESC
		LIMIT ?, ?
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), cursor, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// GetTotal count latest entries
func (r *FeedRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(f.id) as total FROM feeds AS f
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query)).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *FeedRepository) GetByURL(ctx context.Context, u *url.URL) (*feed.Feed, error) {
	query := `
		SELECT f.id, f.url, f.title, f.type, f.status, f.created_at, f.updated_at, f.parsed_at, f.deleted
		FROM feeds AS f
		WHERE f.url = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), u.String())
	f, err := r.scan(rows)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Insert creates a new feed in the DB
func (r *FeedRepository) Insert(ctx context.Context, f *feed.Feed) error {
	query := `
		INSERT INTO feeds
		(url, title, type, status, created_at, updated_at, deleted)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		f.URL,
		f.Title,
		f.Type,
		f.Status,
		f.CreatedAt,
		f.UpdatedAt,
		f.Deleted,
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

// Update updates a feed in the DB
func (r *FeedRepository) Update(ctx context.Context, f *feed.Feed) error {
	query := `
		UPDATE feeds
		SET type = ?, title = ?, status = ?, updated_at = ?, parsed_at = ?, deleted = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		f.Type,
		f.Title,
		f.Status,
		f.UpdatedAt,
		f.ParsedAt,
		f.Deleted,
		f.ID,
	)

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

func (r *FeedRepository) scan(rows Scanable) (*feed.Feed, error) {
	var feed feed.Feed
	var parsedAt mysql.NullTime

	err := rows.Scan(
		&feed.ID,
		&feed.URL,
		&feed.Title,
		&feed.Type,
		&feed.Status,
		&feed.CreatedAt,
		&feed.UpdatedAt,
		&parsedAt,
		&feed.Deleted,
	)

	if parsedAt.Valid {
		feed.ParsedAt = parsedAt.Time
	}

	if err != nil {
		return nil, err
	}

	return &feed, nil
}
