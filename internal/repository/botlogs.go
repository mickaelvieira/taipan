package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/client"
)

// BotlogRepository the Bot logs repository
type BotlogRepository struct {
	db *sql.DB
}

// Insert saves an entry in the bookmark log
func (r *BotlogRepository) Insert(ctx context.Context, l *client.Result) error {
	query := `
		INSERT INTO bot_logs
		(checksum, content_type, response_status_code, response_reason_phrase, response_headers, request_uri, request_method, request_headers, created_at)
		VALUES
		(UNHEX(?), ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		l.Checksum,
		l.ContentType,
		l.RespStatusCode,
		l.RespReasonPhrase,
		l.ReqHeaders,
		l.ReqURI,
		l.ReqMethod,
		l.ReqHeaders,
		l.CreatedAt,
	)

	return err
}

// FindLatestByURL find the latest log entry for a given URL
func (r *BotlogRepository) FindLatestByURL(ctx context.Context, URL string) (*client.Result, error) {
	query := `
		SELECT l.id, HEX(l.checksum), content_type, l.response_status_code, l.response_reason_phrase, l.response_headers, l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY l.created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), URL)
	log, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return log, nil
}

// FindByURL finds the log entries for a given URL
func (r *BotlogRepository) FindByURL(ctx context.Context, URL string) ([]*client.Result, error) {
	var logs []*client.Result

	query := `
		SELECT l.id, HEX(l.checksum), content_type, l.response_status_code, l.response_reason_phrase, l.response_headers, l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log *client.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

// FindByURLAndStatus finds the log entries for a given URL and status
func (r *BotlogRepository) FindByURLAndStatus(ctx context.Context, URL string, status int) ([]*client.Result, error) {
	var logs []*client.Result

	query := `
		SELECT l.id, HEX(l.checksum), content_type, l.response_status_code, l.response_reason_phrase, l.response_headers, l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ? AND l.response_status_code = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URL, status)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log *client.Result
		log, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *BotlogRepository) scan(rows Scanable) (*client.Result, error) {
	var l client.Result

	err := rows.Scan(
		&l.ID,
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
		return nil, err
	}

	return &l, nil
}
