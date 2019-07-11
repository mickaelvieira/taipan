package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/user"
)

// UserRepository the Bookmark repository
type UserRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `
		SELECT u.id, u.username, u.firstname, u.lastname, u.status, u.image_name, u.image_width, u.image_height, u.image_format, u.created_at, u.updated_at
		FROM users as u
		WHERE u.id = ?
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), id)
	u, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Update update a user
func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET firstname = ?, lastname = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		u.Firstname,
		u.Lastname,
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
