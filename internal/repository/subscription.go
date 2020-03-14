package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/domain/subscription"
	"github.com/mickaelvieira/taipan/internal/domain/syndication"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// SubscriptionRepository the Feed repository
type SubscriptionRepository struct {
	db *sql.DB
}

type SubscriptionSearchParams struct {
	Terms []string
	Tags  []string
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
		return nil, errors.Wrap(err, "execute")
	}

	var subscribers []string
	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			return nil, errors.Wrap(err, "scan")
		}
		subscribers = append(subscribers, userID)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return subscribers, nil
}

// FindAll --
func (r *SubscriptionRepository) FindAll(ctx context.Context, u *user.User, search *SubscriptionSearchParams, paging *OffsetPagination) ([]*subscription.Subscription, error) {
	query := `
		SELECT DISTINCT s.id, su.user_id, s.url, s.domain, s.title, s.type, su.subscribed, s.frequency, su.created_at, su.updated_at
		FROM syndication AS s
		%s
		LEFT JOIN subscriptions AS su ON s.id = su.source_id AND su.user_id = ?
		%s
		ORDER BY s.title ASC
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

	args = append(args, u.ID)

	var s string
	if len(search.Terms) > 0 {
		var a []interface{}
		s, a = getSyndicationSearch(search.Terms)
		args = append(args, a...)
	}

	if len(search.Terms) == 0 && len(search.Tags) == 0 {
		s = "su.subscribed = 1"
	}

	if s != "" {
		// @TODO I need to improve the contrustion of those queries
		s = fmt.Sprintf("WHERE %s", s) // #nosec
	}

	args = append(args, paging.Offset)
	args = append(args, paging.Limit)

	query = formatQuery(fmt.Sprintf(query, t, s))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*subscription.Subscription
	for rows.Next() {
		var d *subscription.Subscription
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
func (r *SubscriptionRepository) GetTotal(ctx context.Context, u *user.User, search *SubscriptionSearchParams) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(DISTINCT s.id) as total
		FROM syndication AS s
		%s
		LEFT JOIN subscriptions AS su ON s.id = su.source_id AND su.user_id = ?
		%s
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

	args = append(args, u.ID)

	var s string
	if len(search.Terms) > 0 {
		var a []interface{}
		s, a = getSyndicationSearch(search.Terms)
		args = append(args, a...)
	}

	if len(search.Terms) == 0 && len(search.Tags) == 0 {
		s = "su.subscribed = 1"
	}

	if s != "" {
		// @TODO I need to improve the contrustion of those queries
		s = fmt.Sprintf("WHERE %s", s) // #nosec
	}

	query = formatQuery(fmt.Sprintf(query, t, s))

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan")
	}

	return total, nil
}

// CountUserSubscription --
func (r *SubscriptionRepository) CountUserSubscription(ctx context.Context, u *user.User) (int32, error) {
	query := `
		SELECT COUNT(s.id) as total
		FROM syndication AS s
		INNER JOIN subscriptions AS su ON s.id = su.source_id
		WHERE su.user_id = ? AND su.subscribed = 1
	`
	var total int32
	err := r.db.QueryRowContext(ctx, formatQuery(query), u.ID).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan")
	}

	return total, nil
}

// GetByURL find a single entry by URL
func (r *SubscriptionRepository) GetByURL(ctx context.Context, usr *user.User, u *url.URL) (*subscription.Subscription, error) {
	query := `
		SELECT s.id, su.user_id, s.url, s.domain, s.title, s.type, su.subscribed, s.frequency, su.created_at, su.updated_at
		FROM subscriptions AS su
		INNER JOIN syndication AS s ON s.id = su.source_id
		WHERE su.user_id = ? AND s.url = ?
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

func (r *SubscriptionRepository) scan(rows Scanable) (*subscription.Subscription, error) {
	var s subscription.Subscription
	var subscribed sql.NullBool
	var createdAt mysql.NullTime
	var updatedAt mysql.NullTime
	var userID sql.NullString

	err := rows.Scan(
		&s.ID,
		&userID,
		&s.URL,
		&s.Domain,
		&s.Title,
		&s.Type,
		&subscribed,
		&s.Frequency,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	if userID.Valid {
		s.UserID = userID.String
	}

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

	return &s, nil
}
