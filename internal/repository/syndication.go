package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"log"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

// SyndicationRepository the Feed repository
type SyndicationRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *SyndicationRepository) GetByID(ctx context.Context, id string) (*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.title, s.type, s.status, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM feeds AS s
		WHERE s.id = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), id)
	s, err := r.scan(rows)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetOutdatedSources returns the sources which have been last updated more than 24 hrs
func (r *SyndicationRepository) GetOutdatedSources(ctx context.Context, f http.Frequency) ([]*syndication.Source, error) {
	var results []*syndication.Source
	query := `
		SELECT s.id, s.url, s.title, s.type, s.status, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM feeds AS s
		WHERE s.deleted = 0 AND s.paused = 0 AND s.frequency = ? AND (s.parsed_at IS NULL OR s.parsed_at < DATE_SUB(NOW(), INTERVAL %s))
		ORDER BY s.parsed_at ASC
		LIMIT ?;
		`
	query = fmt.Sprintf(query, f.SQLInterval())
	rows, err := r.db.QueryContext(ctx, formatQuery(query), f, 50)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		s, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

// FindAll find newest entries
func (r *SyndicationRepository) FindAll(ctx context.Context, isPaused bool, cursor int32, limit int32) ([]*syndication.Source, error) {
	var results []*syndication.Source

	query := `
		SELECT s.id, s.url, s.title, s.type, s.status, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM feeds AS s
		WHERE %s
		ORDER BY s.created_at DESC
		LIMIT ?, ?
	`
	var where []string
	var args []interface{}

	where = append(where, "s.deleted = 0")
	// where = append(where, "s.paused = ?")
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

	// args = append(args, isPaused)
	args = append(args, cursor)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
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
func (r *SyndicationRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(s.id) as total FROM feeds AS s WHERE s.deleted = 0
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query)).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *SyndicationRepository) GetByURL(ctx context.Context, u *url.URL) (*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.title, s.type, s.status, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM feeds AS s
		WHERE s.url = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), u.UnescapeString())
	s, err := r.scan(rows)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ExistWithURL checks whether a source already exists this the same URL
func (r *SyndicationRepository) ExistWithURL(ctx context.Context, u *url.URL) (bool, error) {
	_, err := r.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

// Insert creates a new source in the DB
func (r *SyndicationRepository) Insert(ctx context.Context, s *syndication.Source) error {
	query := `
		INSERT INTO feeds
		(url, title, type, status, created_at, updated_at, deleted, paused, frequency)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.URL,
		s.Title,
		s.Type,
		s.Status,
		s.CreatedAt,
		s.UpdatedAt,
		s.Deleted,
		s.IsPaused,
		s.Frequency,
	)

	if err == nil {
		var ID int64
		ID, err = result.LastInsertId()
		if err == nil {
			s.ID = strconv.FormatInt(ID, 10)
		}
	}

	return err
}

// Update updates a source in the DB
func (r *SyndicationRepository) Update(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE feeds
		SET type = ?, title = ?, status = ?, updated_at = ?, parsed_at = ?, deleted = ?, paused = ?, frequency = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.Type,
		s.Title,
		s.Status,
		s.UpdatedAt,
		s.ParsedAt,
		s.Deleted,
		s.IsPaused,
		s.Frequency,
		s.ID,
	)

	return err
}

// UpdateURL updates the source's URL and UpdatedAt fields
func (r *SyndicationRepository) UpdateURL(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE feeds
		SET url = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.URL,
		s.UpdatedAt,
		s.ID,
	)

	return err
}

// Delete soft deletes the source
func (r *SyndicationRepository) Delete(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE feeds
		SET deleted = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.Deleted,
		s.UpdatedAt,
		s.ID,
	)

	return err
}

// UpdateStatus changes the source status enabled/disabled
func (r *SyndicationRepository) UpdateStatus(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE feeds
		SET paused = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.Deleted,
		s.UpdatedAt,
		s.ID,
	)

	return err
}

// InsertIfNotExists stores the source in the database if there is none with the same URL
func (r *SyndicationRepository) InsertIfNotExists(ctx context.Context, s *syndication.Source) error {
	source, err := r.GetByURL(ctx, s.URL)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.Insert(ctx, s)
		}
	} else {
		s.ID = source.ID
	}
	return err
}

// InsertAllIfNotExists stores sources in the database if there are none with the same URL
func (r *SyndicationRepository) InsertAllIfNotExists(ctx context.Context, sources []*syndication.Source) error {
	var err error
	for _, source := range sources {
		err = r.InsertIfNotExists(ctx, source)
		if err != nil {
			break
		}
	}
	return err
}

func (r *SyndicationRepository) scan(rows Scanable) (*syndication.Source, error) {
	var s syndication.Source
	var parsedAt mysql.NullTime

	err := rows.Scan(
		&s.ID,
		&s.URL,
		&s.Title,
		&s.Type,
		&s.Status,
		&s.CreatedAt,
		&s.UpdatedAt,
		&parsedAt,
		&s.Deleted,
		&s.IsPaused,
		&s.Frequency,
	)

	if parsedAt.Valid {
		s.ParsedAt = parsedAt.Time
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}
