package usecase

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
)

// Users use cases errors
var (
	ErrUserDoesNotExist = errors.New("User does not exist")
)

// UpdateUser usecase
func UpdateUser(ctx context.Context, repos *repository.Repositories, u *user.User, f string, l, i string) (*user.User, error) {
	u.Firstname = f
	u.Lastname = l

	err := repos.Users.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	if i != "" {
		err = HandleAvatar(ctx, repos, u, i)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}
