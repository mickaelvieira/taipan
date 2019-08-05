package usecase

import (
	"context"
	"errors"
	"fmt"
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

// CreateUserEmail --
func CreateUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	e := user.NewEmail(v)

	if len(usr.Emails) == 0 {
		e.IsPrimary = true
	}

	if len(usr.Emails) == 1 && !usr.Emails[0].IsConfirmed {
		return fmt.Errorf("You cannot add new email address before confirming your primary email")
	}

	err := repos.Users.CreateUserEmail(ctx, usr, e)
	if err != nil {
		return err
	}

	usr.Emails = append(usr.Emails, e)

	return nil
}

// DeleteUserEmail --
func DeleteUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	var email *user.Email
	var emails []*user.Email
	for _, e := range usr.Emails {
		if e.Value == v {
			email = e
		} else {
			emails = append(emails, e)
		}
	}

	if email == nil {
		return fmt.Errorf("Email %s does not exist", v)
	}

	if email.IsPrimary {
		return fmt.Errorf("Cannot delete primary email address %s", v)
	}

	err := repos.Users.DeleteUserEmail(ctx, usr, email)
	if err != nil {
		return err
	}

	usr.Emails = emails

	return nil
}

// PrimaryUserEmail --
func PrimaryUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	for _, e := range usr.Emails {
		if e.Value == v {
			if !e.IsConfirmed {
				return fmt.Errorf("Cannot mark unconfirmed email address %s as primary", v)
			}
			e.IsPrimary = true
			e.UpdatedAt = time.Now()
		} else {
			e.IsPrimary = false
			e.UpdatedAt = time.Now()
		}
	}

	for _, e := range usr.Emails {
		err := repos.Users.PrimaryUserEmail(ctx, usr, e)
		if err != nil {
			return err
		}
	}

	return nil
}
