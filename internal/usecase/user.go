package usecase

import (
	"context"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"time"
)

// Users use cases errors
var (
	ErrUserDoesNotExist = errors.New("User does not exist")
)

// UpdateUser usecase
func UpdateUser(ctx context.Context, repos *repository.Repositories, usr *user.User, f string, l, i string) error {
	usr.Firstname = f
	usr.Lastname = l
	usr.UpdatedAt = time.Now()

	err := repos.Users.Update(ctx, usr)
	if err != nil {
		return err
	}

	if i != "" {
		err = HandleAvatar(ctx, repos, usr, i)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateTheme usecase
func UpdateTheme(ctx context.Context, repos *repository.Repositories, usr *user.User, t string) error {
	usr.Theme = t
	usr.UpdatedAt = time.Now()

	err := repos.Users.UpdateTheme(ctx, usr)
	if err != nil {
		return err
	}
	return nil
}
