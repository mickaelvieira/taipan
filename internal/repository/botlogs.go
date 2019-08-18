package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"

	"github.com/pkg/errors"
)

// BotlogRepository the Bot logs repository
type BotlogRepository struct {
	db *sql.DB
}

// Insert saves an entry in the bookmark log
func (r *BotlogRepository) Insert(ctx context.Context, l *http.Result) error {
	query := `
		INSERT INTO bot_logs
		(failed, failure_reason, checksum, content_type, response_status_code, response_reason_phrase, response_headers, request_uri, request_method, request_headers, created_at)
		VALUES
		(?, ?, UNHEX(?), ?, ?, ?, ?, ?, ?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		query,
		l.Failed,
		l.FailureReason,
		l.Checksum,
		l.ContentType,
		l.RespStatusCode,
		l.RespReasonPhrase,
		l.RespHeaders,
		l.ReqURI,
		l.ReqMethod,
		l.ReqHeaders,
		l.CreatedAt,
	)

	if err == nil {
		l.ID = db.GetLastInsertID(res)
	}

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// FindAll finds the log entries for a given URL
func (r *BotlogRepository) FindAll(ctx context.Context, URL *url.URL, cursor int32, limit int32) ([]*http.Result, error) {
	var logs []*http.Result

	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY created_at DESC
		LIMIT ?, ?
	`
	var args []interface{}

	args = append(args, URL.UnescapeString())
	args = append(args, cursor)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, formatQuery(query), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var log *http.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return logs, nil
}

// CountAll --
func (r *BotlogRepository) CountAll(ctx context.Context, URL *url.URL) (int32, error) {
	var total int32

	query := `
		SELECT COUNT(l.id) as total
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY created_at DESC
	`
	var args []interface{}

	args = append(args, URL.UnescapeString())

	query = formatQuery(query)

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return total, errors.Wrap(err, "scan rows")
	}

	return total, nil
}

// FindLatestByURL find the latest log entry for a given URL
func (r *BotlogRepository) FindLatestByURL(ctx context.Context, URL *url.URL) (*http.Result, error) {
	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY l.created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), URL.String())
	log, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return log, nil
}

// FindPreviousByURL find the previous log entry for a given URL
func (r *BotlogRepository) FindPreviousByURL(ctx context.Context, URL *url.URL, c *http.Result) (*http.Result, error) {
	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ? AND l.id != ?
		ORDER BY l.created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), URL.String(), c.ID)
	log, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return log, nil
}

// FindByURL finds the log entries for a given URL
func (r *BotlogRepository) FindByURL(ctx context.Context, URL *url.URL) ([]*http.Result, error) {
	var logs []*http.Result

	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URL.String())
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var log *http.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return logs, nil
}

// FindByURLAndStatus finds the log entries for a given URL and status
func (r *BotlogRepository) FindByURLAndStatus(ctx context.Context, URL *url.URL, status int) ([]*http.Result, error) {
	var logs []*http.Result

	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ? AND l.response_status_code = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URL.String(), status)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var log *http.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return logs, nil
}

// FindFailureByURL finds failed http operations by URL
func (r *BotlogRepository) FindFailureByURL(ctx context.Context, URL *url.URL) ([]*http.Result, error) {
	var logs []*http.Result

	query := `
		SELECT l.id, l.failed, l.failure_reason, HEX(l.checksum), content_type,
		l.response_status_code, l.response_reason_phrase, l.response_headers,
		l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ? AND l.failed = 1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URL.String())
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	for rows.Next() {
		var log *http.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	return logs, nil
}

func (r *BotlogRepository) scan(rows Scanable) (*http.Result, error) {
	var l http.Result

	err := rows.Scan(
		&l.ID,
		&l.Failed,
		&l.FailureReason,
		&l.Checksum,
		&l.ContentType,
		&l.RespStatusCode,
		&l.RespReasonPhrase,
		&l.RespHeaders,
		&l.ReqURI,
		&l.ReqMethod,
		&l.ReqHeaders,
		&l.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	return &l, nil
}
