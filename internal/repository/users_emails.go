package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/user"
)

// UserEmailRepository the User Email repository
type UserEmailRepository struct {
	db *sql.DB
}

// GetEmail retrieves an address email
func (r *UserEmailRepository) GetEmail(ctx context.Context, email string) (*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at
		FROM users_emails as e
		WHERE e.value = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), email)
	u, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return u, nil
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
		SET primary = ?, updated_at = ?
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

// GetUserEmails returns user's emails
func (r *UserEmailRepository) GetUserEmails(ctx context.Context, id string) ([]*user.Email, error) {
	query := `
		SELECT e.id, e.value, e.primary, e.confirmed, e.created_at, e.updated_at
		FROM users_emails as e
		WHERE e.user_id = ?
		`
	rows, err := r.db.QueryContext(ctx, formatQuery(query), id)
	if err != nil {
		return nil, err
	}

	var emails []*user.Email
	for rows.Next() {
		e, err := r.scan(rows)
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

	err := rows.Scan(
		&e.ID,
		&e.Value,
		&e.IsPrimary,
		&e.IsConfirmed,
		&e.CreatedAt,
		&e.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &e, nil
}
