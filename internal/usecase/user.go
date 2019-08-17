package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github/mickaelvieira/taipan/internal/domain/password"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Users use cases errors
var (
	ErrUserDoesNotExist         = errors.New("User does not exist")
	ErrWeakPassword             = errors.New("Your password must be at least 10 characters long")
	ErrInvalidEmail             = errors.New("Your email does not seem to be valid")
	ErrEmailDoesNotExist        = errors.New("Email does not exist")
	ErrEmailIsAlreadyUsed       = errors.New("There is already an account associated to this email address")
	ErrInvalidCreds             = errors.New("Email or password does not match any records in our database")
	ErrInvalidPassword          = errors.New("Your password is not correct") // only for password change
	ErrPrimaryEmailDeletion     = errors.New("You cannot delete your primary email address")
	ErrPrimaryEmailNotConfirmed = errors.New("You cannot add new email address before confirming your primary email")
	ErrEmailNotConfirmed        = errors.New("You cannot use this address as your primary email since it hasn't been unconfirmed yet")
	ErrInvalidResetToken        = errors.New("Your reset token does seem to be valid")
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
	if !user.IsEmailValid(e) {
		return nil, ErrInvalidEmail
	}

	if !password.IsValid(p) {
		return nil, ErrWeakPassword
	}

	h, err := password.Hash(p)
	if err != nil {
		return nil, err
	}

	// can we find the user with the same email?
	_, err = repos.Emails.GetEmail(ctx, e)
	if err == nil {
		return nil, ErrEmailIsAlreadyUsed
	}

	// Any other errors?
	if err != sql.ErrNoRows {
		return nil, err
	}

	ID, err := repos.Users.CreateUser(ctx, h)
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

	emails, err := repos.Emails.GetUserEmails(ctx, u)
	if err != nil {
		return nil, err
	}

	u.Emails = emails

	return u, nil
}

// ForgotPassword --
func ForgotPassword(ctx context.Context, repos *repository.Repositories, e string) error {
	if !user.IsEmailValid(e) {
		return ErrInvalidEmail
	}

	// can we find the user?
	u, err := repos.Users.GetByPrimaryEmail(ctx, e)
	if err != nil {
		return ErrInvalidCreds
	}

	// is there an active token?
	t, err := repos.ResetToken.FindUserActiveToken(ctx, u)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		err = repos.ResetToken.Create(ctx, password.NewResetPasswordToken(u))
		if err != nil {
			return err
		}
	} else {
		logger.Warn("Re-using active token")
		log.Println(t)
	}

	return nil
}

// ResetPassword --
func ResetPassword(ctx context.Context, repos *repository.Repositories, t string, p string) error {
	if !password.IsValid(p) {
		return ErrWeakPassword
	}

	h, err := password.Hash(p)
	if err != nil {
		return err
	}

	e, err := repos.ResetToken.GetToken(ctx, t)
	if err != nil {
		return ErrInvalidResetToken
	}

	if e.IsExpired() || e.IsUsed {
		return ErrInvalidResetToken
	}

	u, err := repos.Users.GetByID(ctx, e.UserID)
	if err != nil {
		return ErrInvalidResetToken
	}

	err = repos.Users.UpdatePassword(ctx, u, h)
	if err != nil {
		return err
	}

	e.IsUsed = true
	e.UsedAt = time.Now()

	err = repos.ResetToken.UpdateUsage(ctx, e)
	if err != nil {
		return err
	}

	return nil
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

// ChangePassword --
func ChangePassword(ctx context.Context, repos *repository.Repositories, usr *user.User, o string, n string) error {
	// can we find the user's password?
	p, err := repos.Users.GetPassword(ctx, usr.ID)
	if err != nil {
		return ErrInvalidPassword
	}

	// is the old password correct?
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(o)); err != nil {
		return ErrInvalidPassword
	}

	if !password.IsValid(n) {
		return ErrWeakPassword
	}

	h, err := password.Hash(n)
	if err != nil {
		return err
	}

	err = repos.Users.UpdatePassword(ctx, usr, h)
	if err != nil {
		return err
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
	_, err := repos.Emails.GetEmail(ctx, v)
	if err == nil {
		return ErrEmailIsAlreadyUsed
	}

	emails, err := repos.Emails.GetUserEmails(ctx, usr)
	if err != nil {
		return err
	}

	if len(emails) == 1 && !emails[0].IsConfirmed {
		return ErrPrimaryEmailNotConfirmed
	}

	err = repos.Emails.CreateUserEmail(ctx, usr, user.NewEmail(v))
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserEmail --
func DeleteUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	e, err := repos.Emails.GetUserEmail(ctx, usr, v)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrEmailDoesNotExist
		}
		return err
	}

	if e.IsPrimary {
		return ErrPrimaryEmailDeletion
	}

	err = repos.Emails.DeleteUserEmail(ctx, usr, e)
	if err != nil {
		return err
	}

	return nil
}

// PrimaryUserEmail --
func PrimaryUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	emails, err := repos.Emails.GetUserEmails(ctx, usr)
	if err != nil {
		return err
	}

	found := false
	for _, e := range emails {
		if e.Value == v {
			if !e.IsConfirmed {
				return ErrEmailNotConfirmed
			}
			found = true
			e.IsPrimary = true
			e.UpdatedAt = time.Now()
		} else {
			e.IsPrimary = false
			e.UpdatedAt = time.Now()
		}
	}

	if !found {
		return ErrEmailDoesNotExist
	}

	for _, e := range emails {
		err := repos.Emails.PrimaryUserEmail(ctx, usr, e)
		if err != nil {
			return err
		}
	}

	return nil
}
