package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/db"
	"github/mickaelvieira/taipan/internal/domain/user"
)

// UserRepository the Bookmark repository
type UserRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `
		SELECT u.id, u.username, u.firstname, u.lastname, u.status, u.theme,
		u.image_name, u.image_width, u.image_height, u.image_format,
		u.created_at, u.updated_at
		FROM users as u
		WHERE u.id = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), id)
	u, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	emails, err := r.GetUserEmails(ctx, id)
	if err != nil {
		return nil, err
	}

	u.Emails = emails

	return u, nil
}

// GetByUsername find a single entry
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	query := `
		SELECT u.id, u.username, u.firstname, u.lastname, u.status, u.theme,
		u.image_name, u.image_width, u.image_height, u.image_format,
		u.created_at, u.updated_at
		FROM users as u
		INNER JOIN users_emails as e ON u.id = e.user_id
		WHERE e.value = ? AND e.primary = 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), username)
	u, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	emails, err := r.GetUserEmails(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	u.Emails = emails

	return u, nil
}

// GetPassword find a single entry
func (r *UserRepository) GetPassword(ctx context.Context, id string) (string, error) {
	var password string
	query := `
		SELECT u.password
		FROM users as u
		WHERE u.id = ?
	`
	err := r.db.QueryRowContext(ctx, formatQuery(query), id).Scan(&password)
	if err != nil {
		return password, err
	}

	return password, nil
}

// CreateUserEmail --
func (r *UserRepository) CreateUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
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
func (r *UserRepository) DeleteUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
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
func (r *UserRepository) PrimaryUserEmail(ctx context.Context, u *user.User, e *user.Email) error {
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
func (r *UserRepository) GetUserEmails(ctx context.Context, id string) ([]*user.Email, error) {
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
		var e user.Email
		if err := rows.Scan(&e.ID, &e.Value, &e.IsPrimary, &e.IsConfirmed, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		emails = append(emails, &e)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}

// Update update a user
func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET firstname = ?, lastname = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.Firstname,
		u.Lastname,
		u.UpdatedAt,
		u.ID,
	)

	return err
}

// UpdateTheme update a user's theme
func (r *UserRepository) UpdateTheme(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET theme = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.Theme,
		u.UpdatedAt,
		u.ID,
	)

	return err
}

// UpdateImage updates the user's image
func (r *UserRepository) UpdateImage(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET image_name = ?, image_width = ?, image_height = ?, image_format = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.Image.Name,
		u.Image.Width,
		u.Image.Height,
		u.Image.Format,
		u.UpdatedAt,
		u.ID,
	)

	return err
}

func (r *UserRepository) scan(rows Scanable) (*user.User, error) {
	var u user.User
	var imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&u.ID,
		&u.Username,
		&u.Firstname,
		&u.Lastname,
		&u.Status,
		&u.Theme,
		&imageName,
		&imageWidth,
		&imageHeight,
		&imageFormat,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if imageName != "" {
		u.Image = user.NewImage(
			imageName,
			imageWidth,
			imageHeight,
			imageFormat,
		)
	}

	return &u, nil
}
