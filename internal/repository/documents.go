package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/db"
	"github.com/mickaelvieira/taipan/internal/domain/checksum"
	"github.com/mickaelvieira/taipan/internal/domain/document"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"strings"

	"github.com/pkg/errors"
)

// DocumentRepository the User Bookmark repository
type DocumentRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *DocumentRepository) GetByID(ctx context.Context, id string) (*document.Document, error) {
	query := `
		SELECT d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), id)
	d, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// GetByURL find a single entry
func (r *DocumentRepository) GetByURL(ctx context.Context, u *url.URL) (*document.Document, error) {
	query := `
		SELECT d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE url = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), u.UnescapeString())
	d, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// ExistWithURL checks whether a document already exists this the same URL
func (r *DocumentRepository) ExistWithURL(ctx context.Context, u *url.URL) (bool, error) {
	_, err := r.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

// GetByChecksum find a single entry
func (r *DocumentRepository) GetByChecksum(ctx context.Context, c checksum.Checksum) (*document.Document, error) {
	query := `
		SELECT d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE d.checksum = UNHEX(?)
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), c.String())
	d, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// GetByIDs find all entries
func (r *DocumentRepository) GetByIDs(ctx context.Context, ids []string) ([]*document.Document, error) {
	var results []*document.Document

	args := make([]interface{}, len(ids))
	for i := range ids {
		args[i] = ids[i]
	}

	query := `
		SELECT d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE id IN (?%s)
	`
	query = fmt.Sprintf(query, strings.Repeat(",?", len(ids)-1))
	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var d *document.Document
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

func (r *DocumentRepository) getPagination(fromID string, toID string) (where []string, args []interface{}) {
	var clause string
	if fromID != "" && toID != "" {
		clause = "d.id < ? AND d.id > ?"
		args = append(args, fromID)
		args = append(args, toID)
	} else if fromID != "" && toID == "" {
		clause = "d.id < ?"
		args = append(args, fromID)
	} else if fromID == "" && toID != "" {
		clause = "d.id > ?"
		args = append(args, toID)
	} else {
		return
	}

	where = append(where, clause)

	return
}

func getPagination(fromID string, toID string) (string, []interface{}) {
	var c string
	var a []interface{}

	if fromID != "" && toID != "" {
		c = "d.id < ? AND d.id > ?"
		a = append(a, fromID)
		a = append(a, toID)
	} else if fromID != "" && toID == "" {
		c = "d.id < ?"
		a = append(a, fromID)
	} else if fromID == "" && toID != "" {
		c = "d.id > ?"
		a = append(a, toID)
	}

	if c != "" {
		c = fmt.Sprintf("AND %s", c)
	}

	return c, a
}

// FindAll find bookmarks
func (r *DocumentRepository) FindAll(ctx context.Context, u *user.User, terms []string, paging *OffsetPagination) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT
		d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM newsfeed AS nf
		INNER JOIN documents AS d ON nf.document_id = d.id
		LEFT JOIN bookmarks AS b ON b.document_id = d.id
		WHERE
		nf.user_id = ? AND
		(b.user_id IS NULL OR b.user_id != ?)
		%s
		ORDER BY d.id DESC
		LIMIT ?, ?
	`

	s, a := getDocumentSearch(terms)

	var args []interface{}
	args = append(args, u.ID)
	args = append(args, u.ID)
	args = append(args, a...)
	args = append(args, paging.Offset)
	args = append(args, paging.Limit)

	query = fmt.Sprintf(formatQuery(query), s)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var d *document.Document
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

// CountAll --
func (r *DocumentRepository) CountAll(ctx context.Context, u *user.User, terms []string) (int32, error) {
	var total int32

	query := `
		SELECT
		COUNT(d.id) as total
		FROM newsfeed AS nf
		INNER JOIN documents AS d ON nf.document_id = d.id
		LEFT JOIN bookmarks AS b ON b.document_id = d.id
		WHERE
		nf.user_id = ? AND
		(b.user_id IS NULL OR b.user_id != ?)
		%s
		ORDER BY d.id DESC
	`
	s, a := getDocumentSearch(terms)

	var args []interface{}
	args = append(args, u.ID)
	args = append(args, u.ID)
	args = append(args, a...)

	query = fmt.Sprintf(formatQuery(query), s)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan rows")
	}
	return total, nil
}

// GetNews find newest entries
func (r *DocumentRepository) GetNews(ctx context.Context, u *user.User, paging *CursorPagination, isDescending bool) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT
		d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM newsfeed AS nf
		INNER JOIN documents AS d ON nf.document_id = d.id
		LEFT JOIN bookmarks AS b ON b.document_id = d.id
		WHERE
		nf.user_id = ? AND
		(b.user_id IS NULL OR b.user_id != ?)
		%s
		ORDER BY d.id %s
		LIMIT ?
	`
	dir := "ASC"
	if isDescending {
		dir = "DESC"
	}

	var args []interface{}
	where, bounds := getPagination(paging.From, paging.To)
	query = fmt.Sprintf(query, where, dir)

	args = append(args, u.ID, u.ID)
	args = append(args, bounds...)
	args = append(args, paging.Limit)

	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var d *document.Document
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

// GetTotalLatestNews returns the number of latest news
func (r *DocumentRepository) GetTotalLatestNews(ctx context.Context, u *user.User, paging *CursorPagination, isDescending bool) (int32, error) {
	var total int32
	query := `
		SELECT COUNT(d.id) AS total
		FROM newsfeed AS nf
		INNER JOIN documents AS d ON nf.document_id = d.id
		LEFT JOIN bookmarks AS b ON b.document_id = d.id
		WHERE
		nf.user_id = ? AND
		(b.user_id IS NULL OR b.user_id != ?)
		%s
		ORDER BY d.id %s
	`
	dir := "ASC"
	if isDescending {
		dir = "DESC"
	}

	var args []interface{}
	where, bounds := getPagination(paging.From, paging.To)
	query = fmt.Sprintf(query, where, dir)

	args = append(args, u.ID, u.ID)
	args = append(args, bounds...)

	err := r.db.QueryRowContext(ctx, formatQuery(query), args...).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan rows")
	}

	return total, nil
}

// GetTotalNews returns the total of new documents
func (r *DocumentRepository) GetTotalNews(ctx context.Context, u *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) AS total
		FROM newsfeed AS nf
		INNER JOIN documents AS d ON nf.document_id = d.id
		LEFT JOIN bookmarks AS b ON b.document_id = d.id
		WHERE
		nf.user_id = ? AND
		(b.user_id IS NULL OR b.user_id != ?)
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query), u.ID, u.ID).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan rows")
	}

	return total, nil
}

// GetDocuments returns the paginated documents
func (r *DocumentRepository) GetDocuments(ctx context.Context, paging *CursorPagination) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT d.id, d.source_id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
		d.image_url, d.image_name, d.image_width, d.image_height, d.image_format,
		d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE %s
		ORDER BY d.id DESC
		LIMIT ?
	`
	where, args := r.getPagination(paging.From, paging.To)
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

	args = append(args, paging.Limit)
	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var d *document.Document
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
func (r *DocumentRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total FROM documents AS d
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query)).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "execute")
	}

	return total, nil
}

// Insert creates a new document in the DB
func (r *DocumentRepository) Insert(ctx context.Context, d *document.Document) error {
	query := `
		INSERT INTO documents
		(url, checksum, charset, language, title, description, created_at, updated_at, deleted)
		VALUES
		(?, UNHEX(?), ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.URL,
		d.Checksum,
		d.Charset,
		d.Lang,
		d.Title,
		d.Description,
		d.CreatedAt,
		d.UpdatedAt,
		d.Deleted,
	)

	if err == nil {
		d.ID = db.GetLastInsertID(res)
	}

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Update updates a document in the DB
func (r *DocumentRepository) Update(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET checksum = UNHEX(?), charset = ?, language = ?, title = ?, description = ?, updated_at = ?, deleted = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.Checksum,
		d.Charset,
		d.Lang,
		d.Title,
		d.Description,
		d.UpdatedAt,
		d.Deleted,
		d.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateURL updates the document's URL and UpdatedAt fields
func (r *DocumentRepository) UpdateURL(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET url = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.URL,
		d.UpdatedAt,
		d.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateImage updates a document's image in the DB
func (r *DocumentRepository) UpdateImage(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET image_url = ?, image_name = ?, image_width = ?, image_height = ?, image_format = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.Image.URL.String(),
		d.Image.Name,
		d.Image.Width,
		d.Image.Height,
		d.Image.Format,
		d.UpdatedAt,
		d.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdateSource --
func (r *DocumentRepository) UpdateSource(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET source_id = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.SourceID,
		d.UpdatedAt,
		d.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// Upsert insert the document or update if there is already one with the same URL
func (r *DocumentRepository) Upsert(ctx context.Context, d *document.Document) error {
	e, err := r.GetByURL(ctx, d.URL)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return r.Insert(ctx, d)
	}

	d.ID = e.ID

	// Populate the document with the existing image to avoid refetching it
	if d.Image != nil && e.Image != nil && d.Image.URL.String() == e.Image.URL.String() {
		d.Image = e.Image
	}

	return r.Update(ctx, d)
}

// Delete soft deletes the document
func (r *DocumentRepository) Delete(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET deleted = ? updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		d.Deleted,
		d.UpdatedAt,
		d.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

func (r *DocumentRepository) scan(rows Scanable) (*document.Document, error) {
	var d document.Document
	var sourceID sql.NullString
	var imageURL, imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&d.ID,
		&sourceID,
		&d.URL,
		&d.Checksum,
		&d.Charset,
		&d.Lang,
		&d.Title,
		&d.Description,
		&imageURL,
		&imageName,
		&imageWidth,
		&imageHeight,
		&imageFormat,
		&d.CreatedAt,
		&d.UpdatedAt,
		&d.Deleted,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	if sourceID.Valid {
		d.SourceID = sourceID.String
	}

	if imageURL != "" {
		i, err := document.NewImage(imageURL, imageName, imageWidth, imageHeight, imageFormat)
		if err != nil {
			return nil, errors.Wrap(err, "cannot create image")
		}
		d.Image = i
	}

	return &d, nil
}
