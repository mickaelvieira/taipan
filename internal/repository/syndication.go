package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/db"
	"github.com/mickaelvieira/taipan/internal/domain/http"
	"github.com/mickaelvieira/taipan/internal/domain/syndication"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// SyndicationRepository the Feed repository
type SyndicationRepository struct {
	db *sql.DB
}

type SyndicationSearchParams struct {
	Terms  []string
	Tags   []string
	Hidden bool
	Paused bool
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
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE s.deleted = 0 AND s.paused = 0 AND s.frequency = ? AND (s.parsed_at IS NULL OR s.parsed_at < DATE_SUB(NOW(), INTERVAL 1 HOUR))
		ORDER BY s.parsed_at ASC
		LIMIT ?;
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), f, 50)
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

// GetUndiscoveredSources returns the sources that have been parsed and are paused
func (r *SyndicationRepository) GetUndiscoveredSources(ctx context.Context) ([]*syndication.Source, error) {
	query := `
		SELECT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		WHERE s.deleted = 0 AND s.paused = 1 AND s.parsed_at IS NULL
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

func getSyndicationFilters(hidden bool, paused bool) string {
	var where []string
	if hidden {
		where = append(where, "s.deleted = 1")
	} else {
		where = append(where, "s.deleted = 0")
	}
	if paused {
		where = append(where, "s.paused = 1")
	} else {
		where = append(where, "s.paused = 0")
	}
	return strings.Join(where, " AND ")
}

// FindAll find newest entries
func (r *SyndicationRepository) FindAll(ctx context.Context, search *SyndicationSearchParams, paging *OffsetPagination, sort *Sorting) ([]*syndication.Source, error) {
	query := `
		SELECT DISTINCT s.id, s.url, s.domain, s.title, s.type, s.created_at, s.updated_at, s.parsed_at, s.deleted, s.paused, s.frequency
		FROM syndication AS s
		%s
		WHERE %s
		ORDER BY s.%s %s
		LIMIT ?, ?
	`
	var args []interface{}

	var t string
	if len(search.Tags) > 0 {
		p := getMultiInsertPlacements(1, len(search.Tags))
		t = fmt.Sprintf("INNER JOIN syndication_tags_relation AS r ON r.source_id = s.id AND tag_id IN %s", p)
		for _, a := range search.Tags {
			args = append(args, a)
		}
	}

	var w []string
	f := getSyndicationFilters(search.Hidden, search.Paused)
	w = append(w, f)

	var s string
	var a []interface{}
	if len(search.Terms) > 0 {
		s, a = getSyndicationSearch(search.Terms)
		w = append(w, s)
		args = append(args, a...)
	}

	args = append(args, paging.Offset)
	args = append(args, paging.Limit)

	query = formatQuery(fmt.Sprintf(query, t, strings.Join(w, " AND "), sort.By, sort.Dir))

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
func (r *SyndicationRepository) GetTotal(ctx context.Context, search *SyndicationSearchParams) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(DISTINCT s.id) as total
		FROM syndication AS s
		%s
		WHERE %s
	`
	var args []interface{}

	var t string
	if len(search.Tags) > 0 {
		p := getMultiInsertPlacements(1, len(search.Tags))
		t = fmt.Sprintf("INNER JOIN syndication_tags_relation AS r ON r.source_id = s.id AND tag_id IN %s", p)
		for _, a := range search.Tags {
			args = append(args, a)
		}
	}

	var w []string
	f := getSyndicationFilters(search.Hidden, search.Paused)
	w = append(w, f)

	var s string
	var a []interface{}
	if len(search.Terms) > 0 {
		s, a = getSyndicationSearch(search.Terms)
		w = append(w, s)
		args = append(args, a...)
	}

	query = formatQuery(fmt.Sprintf(query, t, strings.Join(w, " AND ")))

	err := r.db.QueryRowContext(ctx, formatQuery(query), args...).Scan(&total)
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

// GetActiveTagIDs --
func (r *SyndicationRepository) GetActiveTagIDs(ctx context.Context) ([]string, error) {
	query := `
		SELECT tag_id FROM syndication_tags_relation GROUP BY tag_id;
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query))
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var ids []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	if err != nil {
		return nil, err
	}

	return ids, nil
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
