package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/domain/newsfeed"

	"github.com/pkg/errors"
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

// AddEntries adds an entries to the newsfeeds
func (r *NewsFeedRepository) AddEntries(ctx context.Context, entries []*newsfeed.Entry) error {
	query := `
		INSERT INTO newsfeed
		(user_id, document_id, created_at)
		VALUES
		%s
	`
	p := getMultiInsertPlacements(len(entries), 3)
	a := getFeedEntriesParameters(entries)

	_, err := r.db.ExecContext(ctx, formatQuery(fmt.Sprintf(query, p)), a...)
	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}
