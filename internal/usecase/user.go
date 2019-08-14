package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Users use cases errors
var (
	ErrUserDoesNotExist = errors.New("User does not exist")
	ErrWeakPassword     = errors.New("Your password must be at least 10 characters long")
	ErrInvalidEmail     = errors.New("Your email does not seem to be valid")
	ErrNoEmail          = errors.New("Please provide your email")
	ErrNoPassword       = errors.New("Please provide a valid password")
	ErrEmailExists      = errors.New("There is already an account associated to this email address")
	ErrInvalidCreds     = errors.New("Email or password does not match any records in our database")
)

// Signin --
func Signin(ctx context.Context, repos *repository.Repositories, e string, pwd string) (*user.User, error) {
	// can we find the user?
	u, err := repos.Users.GetByPrimaryEmail(ctx, e)
	if err != nil {
		return nil, ErrInvalidCreds
	}

	// can we find the user's password?
	p, err := repos.Users.GetPassword(ctx, u.ID)
	if err != nil {
		return nil, ErrInvalidCreds
	}

	// do the password match?
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(pwd)); err != nil {
		return nil, ErrInvalidCreds
	}

	return u, nil
}

// Signup --
func Signup(ctx context.Context, repos *repository.Repositories, e string, p string) (*user.User, error) {
	if e == "" {
		return nil, ErrNoEmail
	}

	// @TODO can we do a something better here?
	if !strings.Contains(e, "@") {
		return nil, ErrInvalidEmail
	}

	if p == "" {
		return nil, ErrNoPassword
	}

	// @TODO can we do a something better here?
	if len(p) < 10 {
		return nil, ErrWeakPassword
	}

	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// can we find the user with the same email?
	_, err = repos.Emails.GetEmail(ctx, e)
	if err == nil {
		return nil, ErrEmailExists
	}

	// Any other errors?
	if err != sql.ErrNoRows {
		return nil, err
	}

	ID, err := repos.Users.CreateUser(ctx, string(h))
	if err != nil {
		return nil, err
	}

	u, err := repos.Users.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	email := user.NewEmail(e)
	email.IsPrimary = true // first email is always the primary

	err = repos.Emails.CreateUserEmail(ctx, u, email)
	if err != nil {
		return nil, err
	}

	emails, err := repos.Emails.GetUserEmails(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	u.Emails = emails

	return u, nil
}

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

	err := repos.Emails.CreateUserEmail(ctx, usr, e)
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

	err := repos.Emails.DeleteUserEmail(ctx, usr, email)
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
		err := repos.Emails.PrimaryUserEmail(ctx, usr, e)
		if err != nil {
			return err
		}
	}

	return nil
}
