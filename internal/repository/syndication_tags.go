package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/db"
	"github.com/mickaelvieira/taipan/internal/domain/syndication"

	"github.com/pkg/errors"
)

// SyndicationTagsRepository the Feed repository
type SyndicationTagsRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *SyndicationTagsRepository) GetByID(ctx context.Context, id string) (*syndication.Tag, error) {
	query := `
		SELECT t.id, t.label, t.created_at, t.updated_at
		FROM syndication_tags AS t
		WHERE t.id = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), id)
	s, err := r.scan(rows)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetByIDs find a multiple entries by their ID
func (r *SyndicationTagsRepository) GetByIDs(ctx context.Context, ids []string) ([]*syndication.Tag, error) {
	query := `
		SELECT t.id, t.label, t.created_at, t.updated_at
		FROM syndication_tags AS t
		WHERE t.id IN %s
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

	var results []*syndication.Tag
	for rows.Next() {
		var s *syndication.Tag
		s, err = r.scan(rows)
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

// GetSourceTagIDs find a multiple entries by their ID
func (r *SyndicationTagsRepository) GetSourceTagIDs(ctx context.Context, s *syndication.Source) ([]string, error) {
	query := `
		SELECT t.id
		FROM syndication_tags_relation AS r
		INNER JOIN syndication_tags AS t ON r.tag_id = t.id AND r.source_id = ?
		ORDER BY t.label ASC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), s.ID)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var ids []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
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

// FindAll find newest entries
func (r *SyndicationTagsRepository) FindAll(ctx context.Context) ([]*syndication.Tag, error) {
	query := `
		SELECT t.id, t.label, t.created_at, t.updated_at
		FROM syndication_tags AS t
		ORDER BY t.label ASC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query))
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*syndication.Tag
	for rows.Next() {
		var d *syndication.Tag
		d, err = r.scan(rows)
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
func (r *SyndicationTagsRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(t.id) as total FROM syndication_tags AS t
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query)).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan")
	}

	return total, nil
}

// Insert creates a new source in the DB
func (r *SyndicationTagsRepository) Insert(ctx context.Context, t *syndication.Tag) error {
	query := `
		INSERT INTO syndication_tags
		(label, created_at, updated_at)
		VALUES
		(?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.Label,
		t.CreatedAt,
		t.UpdatedAt,
	)

	if err == nil {
		t.ID = db.GetLastInsertID(res)
	}

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Update --
func (r *SyndicationTagsRepository) Update(ctx context.Context, t *syndication.Tag) error {
	query := `
		UPDATE syndication_tags
		SET label = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.Label,
		t.UpdatedAt,
		t.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Delete --
func (r *SyndicationTagsRepository) Delete(ctx context.Context, t *syndication.Tag) error {
	query := `
		DELETE FROM syndication_tags WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

func (r *SyndicationTagsRepository) scan(rows Scanable) (*syndication.Tag, error) {
	var t syndication.Tag

	err := rows.Scan(
		&t.ID,
		&t.Label,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	return &t, nil
}
