package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// SyndicationRepository the Feed repository
type SyndicationRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *SyndicationRepository) GetByID(ctx context.Context, id string) (*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE s.id = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), id)
	s, err := r.scan(rows)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetByIDs find a multiple entries by their ID
func (r *SyndicationRepository) GetByIDs(ctx context.Context, ids []string) ([]*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE s.id IN %s
	`
	args := make([]interface{}, len(ids))
	for i, a := range ids {
		args[i] = a
	}

	p := getMultiInsertPlacements(1, len(ids))
	rows, err := r.db.QueryContext(ctx, formatQuery(fmt.Sprintf(query, p)), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*syndication.Source
	for rows.Next() {
		s, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetDocumentSource retrieve the soource (if any) from which the document was created from
func (r *SyndicationRepository) GetDocumentSource(ctx context.Context, id string) (*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM documents AS d
		INNER JOIN syndication AS s ON d.source_id = s.id
		WHERE d.id = ?
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
	// query := `
	// 	SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
	// 	FROM syndication AS s
	// 	WHERE s.deleted = 0 AND s.paused = 0 AND s.frequency = ? AND (s.parsed_at IS NULL OR s.parsed_at < DATE_SUB(NOW(), INTERVAL %s))
	// 	ORDER BY s.parsed_at ASC
	// 	LIMIT ?;
	// 	`
	// query = fmt.Sprintf(query, f.SQLInterval())
	// rows, err := r.db.QueryContext(ctx, formatQuery(query), f, 50)

	// @TODO we ignore the frequency for now to how whether the calculation is actually working but
	// we will use the query abose later on
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE s.deleted = 0 AND s.paused = 0 AND (s.parsed_at IS NULL OR s.parsed_at < DATE_SUB(NOW(), INTERVAL 1 HOUR))
		ORDER BY s.parsed_at ASC
		LIMIT ?;
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), 50)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*syndication.Source
	for rows.Next() {
		s, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return results, nil
}

func getSyndicationFilters(showDeleted bool, pausedOnly bool) string {
	var where []string
	if showDeleted {
		where = append(where, "s.deleted = 1")
	} else {
		where = append(where, "s.deleted = 0")
	}
	if pausedOnly {
		where = append(where, "s.paused = 1")
	}
	return strings.Join(where, " AND ")
}

// FindAll find newest entries
func (r *SyndicationRepository) FindAll(ctx context.Context, terms []string, showDeleted bool, pausedOnly bool, offset int32, limit int32) ([]*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE %s
		ORDER BY s.title ASC
		LIMIT ?, ?
	`
	var args []interface{}

	var w []string
	f := getSyndicationFilters(showDeleted, pausedOnly)
	w = append(w, f)

	var s string
	var a []interface{}
	if len(terms) > 0 {
		s, a = getSyndicationSearch(terms)
		w = append(w, s)
		args = append(args, a...)
	}

	args = append(args, offset)
	args = append(args, limit)

	query = formatQuery(fmt.Sprintf(query, strings.Join(w, " AND ")))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*syndication.Source
	for rows.Next() {
		d, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return results, nil
}

// GetTotal count latest entries
func (r *SyndicationRepository) GetTotal(ctx context.Context, terms []string, showDeleted bool, pausedOnly bool) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(s.id) as total FROM syndication AS s WHERE %s
	`

	var w []string
	f := getSyndicationFilters(showDeleted, pausedOnly)
	w = append(w, f)

	var s string
	var a []interface{}
	if len(terms) > 0 {
		s, a = getSyndicationSearch(terms)
		w = append(w, s)
	}

	query = formatQuery(fmt.Sprintf(query, strings.Join(w, " AND ")))

	err := r.db.QueryRowContext(ctx, formatQuery(query), a...).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan")
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *SyndicationRepository) GetByURL(ctx context.Context, u *url.URL) (*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
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
		INSERT INTO syndication
		(url, domain, title, type, created_at, updated_at, deleted, paused, frequency)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.URL,
		s.Domain,
		s.Title,
		s.Type,
		s.CreatedAt,
		s.UpdatedAt,
		s.IsDeleted,
		s.IsPaused,
		s.Frequency,
	)

	if err == nil {
		s.ID = db.GetLastInsertID(res)
	}

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Update updates a source in the DB
func (r *SyndicationRepository) Update(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE syndication
		SET domain = ?, type = ?, title = ?, updated_at = ?, parsed_at = ?, deleted = ?, paused = ?, frequency = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.Domain,
		s.Type,
		s.Title,
		s.UpdatedAt,
		s.ParsedAt,
		s.IsDeleted,
		s.IsPaused,
		s.Frequency,
		s.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateURL updates the source's URL and UpdatedAt fields
func (r *SyndicationRepository) UpdateURL(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE syndication
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateVisibility soft deletes the source
func (r *SyndicationRepository) UpdateVisibility(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE syndication
		SET deleted = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.IsDeleted,
		s.UpdatedAt,
		s.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateStatus changes whether the source should be parsed or it is paused
func (r *SyndicationRepository) UpdateStatus(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE syndication
		SET paused = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.IsPaused,
		s.UpdatedAt,
		s.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateTitle changes the source title
func (r *SyndicationRepository) UpdateTitle(ctx context.Context, s *syndication.Source) error {
	query := `
		UPDATE syndication
		SET title = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.Title,
		s.UpdatedAt,
		s.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
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

// Tag tags a syndication source
func (r *SyndicationRepository) Tag(ctx context.Context, s *syndication.Source, t *syndication.Tag) error {
	query := `
		INSERT INTO syndication_tags_relation
		(source_id, tag_id)
		VALUES
		(?, ?)
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.ID,
		t.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Untag remove tag
func (r *SyndicationRepository) Untag(ctx context.Context, s *syndication.Source, t *syndication.Tag) error {
	query := `
		DELETE FROM syndication_tags_relation
		WHERE source_id = ? AND tag_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		s.ID,
		t.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

func (r *SyndicationRepository) scan(rows Scanable) (*syndication.Source, error) {
	var s syndication.Source
	var parsedAt mysql.NullTime

	err := rows.Scan(
		&s.ID,
		&s.URL,
		&s.Domain,
		&s.Title,
		&s.Type,
		&s.CreatedAt,
		&s.UpdatedAt,
		&parsedAt,
		&s.IsDeleted,
		&s.IsPaused,
		&s.Frequency,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	if parsedAt.Valid {
		s.ParsedAt = parsedAt.Time
	}

	return &s, nil
}
