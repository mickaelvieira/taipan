package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"strings"
)

// DocumentRepository the User Bookmark repository
type DocumentRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *DocumentRepository) GetByID(ctx context.Context, id string) (*document.Document, error) {
	query := `
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
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
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
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
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
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
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE id IN (?%s)
	`
	query = fmt.Sprintf(query, strings.Repeat(",?", len(ids)-1))
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

	return c, a
}

// GetNews find newest entries
func (r *DocumentRepository) GetNews(ctx context.Context, u *user.User, fromID string, toID string, limit int32, isDescending bool) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT
		d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description,
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
	where, bounds := getPagination(fromID, toID)
	query = fmt.Sprintf(query, where, dir)

	args = append(args, u.ID, u.ID)
	args = append(args, bounds...)
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

// GetTotalLatestNews returns the number of latest news
func (r *DocumentRepository) GetTotalLatestNews(ctx context.Context, u *user.User, fromID string, toID string, isDescending bool) (int32, error) {
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
	where, bounds := getPagination(fromID, toID)
	query = fmt.Sprintf(query, where, dir)

	args = append(args, u.ID, u.ID)
	args = append(args, bounds...)

	err := r.db.QueryRowContext(ctx, formatQuery(query), args...).Scan(&total)
	if err != nil {
		return total, err
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
		return total, err
	}

	return total, nil
}

// GetDocuments returns the paginated documents
func (r *DocumentRepository) GetDocuments(ctx context.Context, fromID string, toID string, limit int32) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
		FROM documents AS d
		WHERE %s
		ORDER BY d.id DESC
		LIMIT ?
	`
	where, args := r.getPagination(fromID, toID)
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

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
func (r *DocumentRepository) GetTotal(ctx context.Context) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(d.id) as total FROM documents AS d
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query)).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// Insert creates a new document in the DB
func (r *DocumentRepository) Insert(ctx context.Context, d *document.Document) error {
	query := `
		INSERT INTO documents
		(url, checksum, charset, language, title, description, status, created_at, updated_at, deleted)
		VALUES
		(?, UNHEX(?), ?, ?, ?, ?, ?, ?, ?, ?)
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
		d.Status,
		d.CreatedAt,
		d.UpdatedAt,
		d.Deleted,
	)

	if err == nil {
		d.ID = db.GetLastInsertID(res)
	}

	return err
}

// Update updates a document in the DB
func (r *DocumentRepository) Update(ctx context.Context, d *document.Document) error {
	query := `
		UPDATE documents
		SET checksum = UNHEX(?), charset = ?, language = ?, title = ?, description = ?, status = ?, updated_at = ?, deleted = ?
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
		d.Status,
		d.UpdatedAt,
		d.Deleted,
		d.ID,
	)

	return err
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

	return err
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

	return err
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

	return err
}

func (r *DocumentRepository) scan(rows Scanable) (*document.Document, error) {
	var d document.Document
	var imageURL, imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&d.ID,
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
		return nil, err
	}

	if imageURL != "" {
		i, err := document.NewImage(imageURL, imageName, imageWidth, imageHeight, imageFormat)
		if err != nil {
			return nil, err
		}
		d.Image = i
	}

	return &d, nil
}
