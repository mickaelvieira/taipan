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
)

// SubscriptionRepository the Feed repository
type SubscriptionRepository struct {
	db *sql.DB
}

// FindAll find newest entries
func (r *SubscriptionRepository) FindAll(ctx context.Context, u *user.User, cursor int32, limit int32) ([]*subscription.Subscription, error) {
	var results []*subscription.Subscription

	query := `
		SELECT sy.id, sy.url, sy.title, sy.type, su.subscribed, su.created_at, su.updated_at
		FROM subscriptions AS su
		INNER JOIN syndication AS sy ON sy.id = su.source_id
		WHERE %s
		ORDER BY su.created_at DESC
		LIMIT ?, ?
	`
	var where []string
	var args []interface{}

	where = append(where, "sy.deleted = 0")
	where = append(where, "su.user_id = ?")
	query = fmt.Sprintf(query, strings.Join(where, " AND "))

	args = append(args, u.ID)
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
func (r *SubscriptionRepository) GetTotal(ctx context.Context, u *user.User) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(sy.id) as total
		FROM subscriptions AS su
		INNER JOIN syndication AS sy ON sy.id = su.source_id
		WHERE sy.deleted = 0 AND su.user_id = ?
		ORDER BY su.created_at DESC
	`

	err := r.db.QueryRowContext(ctx, formatQuery(query), u.ID).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *SubscriptionRepository) GetByURL(ctx context.Context, usr *user.User, u *url.URL) (*subscription.Subscription, error) {
	query := `
		SELECT sy.id, sy.url, sy.title, sy.type, su.subscribed, su.created_at, su.updated_at
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

	err := rows.Scan(
		&s.ID,
		&s.URL,
		&s.Title,
		&s.Type,
		&s.Subscribed,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
