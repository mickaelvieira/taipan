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

// FindLatestByURI find the latest log entry for a given URI
func (r *BotlogRepository) FindLatestByURI(ctx context.Context, URI string) (*client.Result, error) {
	query := `
		SELECT l.id, HEX(l.checksum), content_type, l.response_status_code, l.response_reason_phrase, l.response_headers, l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY l.created_at DESC
		LIMIT 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), URI)
	log, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return log, nil
}

// FindByURI finds the log entries for a give URI
func (r *BotlogRepository) FindByURI(ctx context.Context, URI string) ([]*client.Result, error) {
	var logs []*client.Result

	query := `
		SELECT l.id, HEX(l.checksum), content_type, l.response_status_code, l.response_reason_phrase, l.response_headers, l.request_uri, l.request_method, l.request_headers, l.created_at
		FROM bot_logs AS l
		WHERE l.request_uri = ?
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), URI)
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
