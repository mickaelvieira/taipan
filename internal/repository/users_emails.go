package repository

import (
	"context"
	"database/sql"
	"github.com/mickaelvieira/taipan/internal/db"
	"github.com/mickaelvieira/taipan/internal/domain/user"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// UserEmailRepository the User Email repository
type UserEmailRepository struct {
	db *sql.DB
}

// GetEmail retrieves an address email
func (r *UserEmailRepository) GetEmail(ctx context.Context, v string) (*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at, e.confirmed_at
		FROM users_emails as e
		WHERE e.value = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), v)
	u, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserEmailByValue retrieves an address email
func (r *UserEmailRepository) GetUserEmailByValue(ctx context.Context, u *user.User, v string) (*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at, e.confirmed_at
		FROM users_emails as e
		WHERE e.value = ? AND e.user_id = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), v, u.ID)
	e, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// GetUserEmailByID retrieves an address email
func (r *UserEmailRepository) GetUserEmailByID(ctx context.Context, u *user.User, id string) (*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at, e.confirmed_at
		FROM users_emails as e
		WHERE e.id = ? AND e.user_id = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), id, u.ID)
	e, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// CreateUserEmail --
func (r *UserEmailRepository) CreateUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
	query := `
		INSERT INTO users_emails
		(user_id, value, 'primary', confirmed, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.ID,
		e.Value,
		e.IsPrimary,
		e.IsConfirmed,
		e.CreatedAt,
		e.UpdatedAt,
	)

	if err == nil {
		e.ID = db.GetLastInsertID(res)
	}

	return err
}

// DeleteUserEmail --
func (r *UserEmailRepository) DeleteUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
	query := `
		DELETE FROM users_emails WHERE user_id = ? AND id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.ID,
		e.ID,
	)

	return err
}

// PrimaryUserEmail --
func (r *UserEmailRepository) PrimaryUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
	query := `
		UPDATE users_emails
		SET 'primary' = ?, updated_at = ?
		WHERE user_id = ? AND id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		e.IsPrimary,
		e.UpdatedAt,
		u.ID,
		e.ID,
	)

	return err
}

// ConfirmUserEmail --
func (r *UserEmailRepository) ConfirmUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
	query := `
		UPDATE users_emails
		SET confirmed = ?, confirmed_at = ?, updated_at = ?
		WHERE user_id = ? AND id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		e.IsConfirmed,
		e.ConfirmedAt,
		e.UpdatedAt,
		u.ID,
		e.ID,
	)

	return err
}

// GetUserEmails returns user's emails
func (r *UserEmailRepository) GetUserEmails(ctx context.Context, u *user.User) ([]*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at, e.confirmed_at
		FROM users_emails as e
		WHERE e.user_id = ? ORDER BY e.primary DESC
	`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), u.ID)
	if err != nil {
		return nil, err
	}

	var emails []*user.Email
	for rows.Next() {
		var e *user.Email
		e, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		emails = append(emails, e)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}

func (r *UserEmailRepository) scan(rows Scanable) (*user.Email, error) {
	var e user.Email
	var confirmedAt mysql.NullTime

	err := rows.Scan(
		&e.ID,
		&e.Value,
		&e.IsPrimary,
		&e.IsConfirmed,
		&e.CreatedAt,
		&e.UpdatedAt,
		&confirmedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
	}

	if confirmedAt.Valid {
		e.ConfirmedAt = confirmedAt.Time
	}

	return &e, nil
}
