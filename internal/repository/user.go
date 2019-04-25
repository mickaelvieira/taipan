package repository

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/user"
)

// UserRepository the User Bookmark repository
type UserRepository struct {
	db *sql.DB
}

// GetByID find a single entry
func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	var user user.User

	query := "SELECT id, username, firstname, lastname, status, created_at, updated_at FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Status, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
