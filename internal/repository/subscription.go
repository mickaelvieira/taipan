package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

// SubscriptionRepository the Feed repository
type SubscriptionRepository struct {
	db *sql.DB
}

func getSubscriptionSearch(terms []string) (string, []interface{}) {
	var search string
	var args []interface{}

	if len(terms) > 0 {
		var clause = make([]string, len(terms)*2)
		j := 0
		for _, t := range terms {
			clause[j] = "sy.url LIKE ?"
			clause[j+1] = "sy.title LIKE ?"
			args = append(args, "%"+t+"%")
			args = append(args, "%"+t+"%")
			j = j + 2
		}

		search = fmt.Sprintf("(%s)", strings.Join(clause, " OR "))
	} else {
		search = "su.subscribed = 1"
	}

	return search, args
}

func getWhere(showDeleted bool, pausedOnly bool) string {
	var where []string
	if showDeleted {
		where = append(where, "sy.deleted = 1")
	} else {
		where = append(where, "sy.deleted = 0")
	}
	if pausedOnly {
		where = append(where, "sy.paused = 1")
	}
	return strings.Join(where, " AND ")
}

// FindSubscribersIDs find users who have subscribed to the syndication source
func (r *SubscriptionRepository) FindSubscribersIDs(ctx context.Context, sourceID string) ([]string, error) {
	query := `
		SELECT su.user_id
		FROM subscriptions AS su
		WHERE su.source_id = ? AND su.subscribed = 1
	`

	rows, err := r.db.QueryContext(ctx, formatQuery(query), sourceID)
	if err != nil {
		return nil, err
	}

	var subscribers []string
	for rows.Next() {
		var userID string
		err := rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, userID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}

// FindAll --
func (r *SubscriptionRepository) FindAll(ctx context.Context, u *user.User, terms []string, showDeleted bool, pausedOnly bool, offset int32, limit int32) ([]*subscription.Subscription, error) {
	query := `
		SELECT sy.id, sy.url, sy.domain, sy.title, sy.type, su.subscribed, sy.frequency, su.created_at, su.updated_at
		FROM syndication AS sy
		LEFT JOIN subscriptions AS su ON sy.id = su.source_id
		WHERE (su.user_id = ? OR su.user_id IS NULL) AND %s AND %s
		ORDER BY sy.title ASC
		LIMIT ?, ?
	`
	var args []interface{}

	search, t := getSubscriptionSearch(terms)
	where := getWhere(showDeleted, pausedOnly)

	args = append(args, u.ID)
	args = append(args, t...)
	args = append(args, offset)
	args = append(args, limit)

	query = formatQuery(fmt.Sprintf(query, where, search))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var results []*subscription.Subscription
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
func (r *SubscriptionRepository) GetTotal(ctx context.Context, u *user.User, terms []string, showDeleted bool, pausedOnly bool) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(sy.id) as total
		FROM syndication AS sy
		LEFT JOIN subscriptions AS su ON sy.id = su.source_id
		WHERE (su.user_id = ? OR su.user_id IS NULL) AND %s AND %s
	`
	var args []interface{}

	search, t := getSubscriptionSearch(terms)
	where := getWhere(showDeleted, pausedOnly)

	args = append(args, u.ID)
	args = append(args, t...)

	query = formatQuery(fmt.Sprintf(query, where, search))

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *SubscriptionRepository) GetByURL(ctx context.Context, usr *user.User, u *url.URL) (*subscription.Subscription, error) {
	query := `
		SELECT sy.id, sy.url, sy.domain, sy.title, sy.type, su.subscribed, sy.frequency, su.created_at, su.updated_at
		FROM subscriptions AS su
		INNER JOIN syndication AS sy ON sy.id = su.source_id
		WHERE su.user_id = ? AND sy.url = ?
	`
	rows := r.db.QueryRowContext(ctx, formatQuery(query), usr.ID, u.UnescapeString())
	s, err := r.scan(rows)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ExistWithURL checks whether a subscriptions already exists with the same URL
func (r *SubscriptionRepository) ExistWithURL(ctx context.Context, usr *user.User, u *url.URL) (bool, error) {
	_, err := r.GetByURL(ctx, usr, u)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

// Subscribe subscribes a user to a web syndication source
func (r *SubscriptionRepository) Subscribe(ctx context.Context, u *user.User, s *syndication.Source) error {
	query := `
		INSERT INTO subscriptions
		(user_id, source_id, subscribed, created_at, updated_at)
		VALUES
		(?, ?, 1, ?, ?)
		ON DUPLICATE KEY UPDATE updated_at = ?, subscribed = 1
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.ID,
		s.ID,
		time.Now(),
		time.Now(),
		time.Now(),
	)

	return err
}

// Unsubscribe unsubscribes user from a web syndication source
func (r *SubscriptionRepository) Unsubscribe(ctx context.Context, u *user.User, s *syndication.Source) error {
	query := `
		UPDATE subscriptions
		SET subscribed = 0, updated_at = ?
		WHERE user_id = ? AND source_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		time.Now(),
		u.ID,
		s.ID,
	)

	return err
}

func (r *SubscriptionRepository) scan(rows Scanable) (*subscription.Subscription, error) {
	var s subscription.Subscription
	var subscribed sql.NullBool
	var createdAt mysql.NullTime
	var updatedAt mysql.NullTime

	err := rows.Scan(
		&s.ID,
		&s.URL,
		&s.Domain,
		&s.Title,
		&s.Type,
		&subscribed,
		&s.Frequency,
		&createdAt,
		&updatedAt,
	)

	if createdAt.Valid {
		s.CreatedAt = createdAt.Time
	}

	if updatedAt.Valid {
		s.CreatedAt = updatedAt.Time
	}

	if subscribed.Valid {
		s.Subscribed = subscribed.Bool
	} else {
		s.Subscribed = false
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}
