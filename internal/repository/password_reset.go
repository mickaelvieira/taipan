package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/password"
	"github/mickaelvieira/taipan/internal/domain/user"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// PasswordResetRepository --
type PasswordResetRepository struct {
	db *sql.DB
}

// Create --
func (r *PasswordResetRepository) Create(ctx context.Context, pr *password.ResetToken) error {
	query := `
		INSERT INTO password_reset
		(token, user_id, used, expired_at, created_at)
		VALUES
		(?, ?, ?, ?, ?)
`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		pr.Token,
		pr.UserID,
		pr.IsUsed,
		pr.ExpiredAt,
		pr.CreatedAt,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// FindUserActiveToken find a single entry
func (r *PasswordResetRepository) FindUserActiveToken(ctx context.Context, usr *user.User) (*password.ResetToken, error) {
	query := `
		SELECT t.token, t.user_id, t.used, t.expired_at, t.created_at, t.used_at
		FROM password_reset as t
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
func (r *PasswordResetRepository) GetToken(ctx context.Context, v string) (*password.ResetToken, error) {
	query := `
		SELECT t.token, t.user_id, t.used, t.expired_at, t.created_at, t.used_at
		FROM password_reset as t
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
func (r *PasswordResetRepository) UpdateUsage(ctx context.Context, t *password.ResetToken) error {
	query := `
		UPDATE password_reset
		SET used = ?, used_at = ?
		WHERE token = ? AND user_id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		t.IsUsed,
		t.UsedAt,
		t.Token,
		t.UserID,
	)

	return errors.Wrap(err, "execute")
}

func (r *PasswordResetRepository) scan(rows Scanable) (*password.ResetToken, error) {
	var pr password.ResetToken
	var usedAt mysql.NullTime

	err := rows.Scan(
		&pr.Token,
		&pr.UserID,
		&pr.IsUsed,
		&pr.ExpiredAt,
		&pr.CreatedAt,
		&usedAt,
	)

	if usedAt.Valid {
		pr.UsedAt = usedAt.Time
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	return &pr, nil
}
