package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/newsfeed"
	"github/mickaelvieira/taipan/internal/domain/user"
	"strings"
)

// NewsFeedRepository the NewsFeed repository
type NewsFeedRepository struct {
	db *sql.DB
}

func getFeedEntriesParameters(e []*newsfeed.Entry) []interface{} {
	a := make([]interface{}, len(e)*3)
	j := 0
	for i := 0; i < len(e); i++ {
		a[j] = e[i].UserID
		a[j+1] = e[i].DocumentID
		a[j+2] = e[i].CreatedAt
		j = j + 3
	}
	return a
}

func (r *NewsFeedRepository) getPagination(fromID string, toID string) (where []string, args []interface{}) {
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

// AddEntries adds an entries to the newsfeeds
func (r *NewsFeedRepository) AddEntries(ctx context.Context, entries []*newsfeed.Entry) error {
	query := `
		INSERT INTO newsfeed
		(user_id, document_id, created_at)
		VALUES
		%s
	`
	p := getMultiInsertPlacements(len(entries), 3)
	models := make([]ToParams, len(entries))
	for i, v := range entries {
		models[i] = ToParams(v)
	}

	a := getMultipleParameters(models, func(e ToParams) []interface{} {
		return e.Params()
	})

	_, err := r.db.ExecContext(ctx, formatQuery(fmt.Sprintf(query, p)), a...)

	return err
}

// GetNews find newest entries
func (r *NewsFeedRepository) GetNews(ctx context.Context, u *user.User, fromID string, toID string, limit int32, isDescending bool) ([]*document.Document, error) {
	var results []*document.Document

	query := `
		SELECT d.id, d.url, HEX(d.checksum), d.charset, d.language, d.title, d.description, d.image_url, d.image_name, d.image_width, d.image_height, d.image_format, d.created_at, d.updated_at, d.deleted
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

	where, args := r.getPagination(fromID, toID)
	// where = append(where, "nf.user_id = ?")
	// where = append(where, "(b.user_id IS NULL OR b.user_id != ?)")
	query = fmt.Sprintf(query, strings.Join(where, " AND "), dir)

	args = append(args, u.ID, u.ID, limit)

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
