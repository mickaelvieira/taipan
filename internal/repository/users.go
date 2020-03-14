package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mickaelvieira/taipan/internal/db"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"time"

	"github.com/pkg/errors"
)

// UserRepository the Bookmark repository
type UserRepository struct {
	db *sql.DB
}

// CreateUser --
func (r *UserRepository) CreateUser(ctx context.Context, hash string) (string, error) {
	query := `
		INSERT INTO users
		(password, created_at, updated_at)
		VALUES
		(?, ?, ?)
	`
	res, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		hash,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	return db.GetLastInsertID(res), nil
}

// GetByID find a single entry
func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `
		SELECT u.id, u.firstname, u.lastname, u.theme,
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

	return u, nil
}

// GetByIDs find a multiple entries by their ID
func (r *UserRepository) GetByIDs(ctx context.Context, ids []string) ([]*user.User, error) {
	query := `
		SELECT u.id, u.firstname, u.lastname, u.theme,
		u.image_name, u.image_width, u.image_height, u.image_format,
		u.created_at, u.updated_at
		FROM users as u
		WHERE u.id IN %s
	`
	args := make([]interface{}, len(ids))
	for i, a := range ids {
		args[i] = a
	}

	p := getMultiInsertPlacements(1, len(ids))
	rows, err := r.db.QueryContext(ctx, formatQuery(fmt.Sprintf(query, p)), args...)
	if err != nil {
		return nil, errors.Wrap(err, "execute")
	}

	var results []*user.User
	for rows.Next() {
		var s *user.User
		s, err = r.scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "scan rows")
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetByPrimaryEmail find a single entry
func (r *UserRepository) GetByPrimaryEmail(ctx context.Context, v string) (*user.User, error) {
	query := `
		SELECT u.id, u.firstname, u.lastname, u.theme,
		u.image_name, u.image_width, u.image_height, u.image_format,
		u.created_at, u.updated_at
		FROM users as u
		INNER JOIN users_emails as e ON u.id = e.user_id
		WHERE e.value = ? AND e.primary = 1
	`
	row := r.db.QueryRowContext(ctx, formatQuery(query), v)
	u, err := r.scan(row)

	if err != nil {
		return nil, err
	}

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
		if err == sql.ErrNoRows {
			return password, err
		}
		return password, errors.Wrap(err, "scan")
	}

	return password, nil
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

// UpdatePassword update a user's password
func (r *UserRepository) UpdatePassword(ctx context.Context, u *user.User, h string) error {
	query := `
		UPDATE users
		SET password = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(
		ctx,
		formatQuery(query),
		h,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
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

	if err != nil {
		return errors.Wrap(err, "execute")
	}

	return nil
}

func (r *UserRepository) scan(rows Scanable) (*user.User, error) {
	var u user.User
	var imageName, imageFormat string
	var imageWidth, imageHeight int32

	err := rows.Scan(
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Theme,
		&imageName,
		&imageWidth,
		&imageHeight,
		&imageFormat,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "scan")
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
