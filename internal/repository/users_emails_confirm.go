package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/user"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// UserEmailConfirmRepository --
type UserEmailConfirmRepository struct {
	db *sql.DB
}

// Create --
func (r *UserEmailConfirmRepository) Create(ctx context.Context, t *user.EmailConfirmToken) error {
	query := `
		INSERT INTO users_emails_confirm
		(token, user_id, email_id, used, expired_at, created_at)
		VALUES
		(?, ?, ?, ?, ?, ?)
`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.Token,
		t.UserID,
		t.EmailID,
		t.IsUsed,
		t.ExpiredAt,
		t.CreatedAt,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// FindUserActiveToken find a single entry
func (r *UserEmailConfirmRepository) FindUserActiveToken(ctx context.Context, usr *user.User) (*user.EmailConfirmToken, error) {
	query := `
		SELECT t.token, t.user_id, t.email_id, t.used, t.expired_at, t.created_at, t.used_at
		FROM users_emails_confirm as t
		WHERE t.user_id = ? AND t.used = 0 AND t.used_at IS NULL AND t.expired_at > NOW()
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), usr.ID)
	pr, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

// GetToken find a single entry
func (r *UserEmailConfirmRepository) GetToken(ctx context.Context, v string) (*user.EmailConfirmToken, error) {
	query := `
		SELECT t.token, t.user_id, t.email_id, t.used, t.expired_at, t.created_at, t.used_at
		FROM users_emails_confirm as t
		WHERE t.token = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), v)
	pr, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

// UpdateUsage --
func (r *UserEmailConfirmRepository) UpdateUsage(ctx context.Context, t *user.EmailConfirmToken) error {
	query := `
		UPDATE users_emails_confirm
		SET used = ?, used_at = ?
		WHERE token = ? AND user_id = ? AND email_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.IsUsed,
		t.UsedAt,
		t.Token,
		t.UserID,
		t.EmailID,
	)

	return errors.Wrap(err, "execute")
}

func (r *UserEmailConfirmRepository) scan(rows Scanable) (*user.EmailConfirmToken, error) {
	var t user.EmailConfirmToken
	var usedAt mysql.NullTime

	err := rows.Scan(
		&t.Token,
		&t.UserID,
		&t.EmailID,
		&t.IsUsed,
		&t.ExpiredAt,
		&t.CreatedAt,
		&usedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	if usedAt.Valid {
		t.UsedAt = usedAt.Time
	}

	return &t, nil
}
