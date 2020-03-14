package usecase

import (
	"context"
	"database/sql"
	"github.com/mickaelvieira/taipan/internal/domain/errors"
	"github.com/mickaelvieira/taipan/internal/domain/password"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"github.com/mickaelvieira/taipan/internal/logger"
	"github.com/mickaelvieira/taipan/internal/repository"
	"strings"
	"time"

	liberr "github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

// Signin --
func Signin(ctx context.Context, repos *repository.Repositories, e string, p string) (*user.User, error) {
	e = strings.Trim(e, " ")
	p = strings.Trim(p, " ")

	// can we find the user?
	u, err := repos.Users.GetByPrimaryEmail(ctx, e)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, err
	}

	// can we find the user's password?
	d, err := repos.Users.GetPassword(ctx, u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, err
	}

	// do the password match?
	if err := bcrypt.CompareHashAndPassword([]byte(d), []byte(p)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, liberr.Wrap(err, "compare password")
	}

	return u, nil
}

// Signup --
func Signup(ctx context.Context, repos *repository.Repositories, e string, p string) (*user.User, error) {
	e = strings.Trim(e, " ")
	p = strings.Trim(p, " ")

	if !user.IsEmailValid(e) {
		return nil, errors.New(user.ErrEmailIsNotValid, nil)
	}

	if !password.IsValid(p) {
		return nil, errors.New(user.ErrWeakPassword, nil)
	}

	h, err := password.Hash(p)
	if err != nil {
		return nil, err
	}

	// can we find the user with the same email?
	_, err = repos.Emails.GetEmail(ctx, e)
	if err == nil {
		return nil, errors.New(user.ErrEmailIsAlreadyUsed, nil)
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

	if err := repos.Emails.CreateUserEmail(ctx, u, email); err != nil {
		return nil, err
	}

	if err := CreateTokenAndSendConfirmationEmail(ctx, repos, u, email); err != nil {
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
	e = strings.Trim(e, " ")

	if !user.IsEmailValid(e) {
		return errors.New(user.ErrEmailIsNotValid, nil)
	}

	// can we find the user?
	u, err := repos.Users.GetByPrimaryEmail(ctx, e)
	if err != nil {
		if err == sql.ErrNoRows {
			// For security reason, we just ignore the request
			// if we can't find an account linked to the email address.
			// Hackers don't need to know whether or not the account exists
			return nil
		}
		return err
	}

	// is there an active token?
	_, err = repos.PasswordReset.FindUserActiveToken(ctx, u)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		err = repos.PasswordReset.Create(ctx, password.NewResetPasswordToken(u))
		if err != nil {
			return err
		}
	}

	return nil
}

// ResetPassword --
func ResetPassword(ctx context.Context, repos *repository.Repositories, t string, p string) error {
	t = strings.Trim(t, " ")
	p = strings.Trim(p, " ")

	if !password.IsValid(p) {
		return errors.New(user.ErrWeakPassword, nil)
	}

	h, err := password.Hash(p)
	if err != nil {
		return err
	}

	e, err := repos.PasswordReset.GetToken(ctx, t)
	if err != nil {
		return errors.New(user.ErrPasswordResetTokenIsNotValid, err)
	}

	if e.IsExpired() || e.IsUsed {
		return errors.New(user.ErrPasswordResetTokenIsNotValid, nil)
	}

	u, err := repos.Users.GetByID(ctx, e.UserID)
	if err != nil {
		return errors.New(user.ErrPasswordResetTokenIsNotValid, err)
	}

	if err = repos.Users.UpdatePassword(ctx, u, h); err != nil {
		return err
	}

	e.IsUsed = true
	e.UsedAt = time.Now()

	if err := repos.PasswordReset.UpdateUsage(ctx, e); err != nil {
		return err
	}

	return nil
}

// CreateTokenAndSendConfirmationEmail --
func CreateTokenAndSendConfirmationEmail(ctx context.Context, repos *repository.Repositories, u *user.User, e *user.Email) error {
	var t *user.EmailConfirmToken
	// is there an active token?
	t, err := repos.EmailsConfirm.FindUserActiveToken(ctx, u, e)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		t = user.NewEmailConfirmToken(u, e)
		if err := repos.EmailsConfirm.Create(ctx, user.NewEmailConfirmToken(u, e)); err != nil {
			return err
		}
	}

	logger.Debug(t)

	// @TODO send confirmation email

	return nil
}

// ConfirmEmail --
func ConfirmEmail(ctx context.Context, repos *repository.Repositories, v string, loggedInUserID string) error {
	v = strings.Trim(v, " ")

	t, err := repos.EmailsConfirm.GetToken(ctx, v)
	if err != nil {
		return errors.New(user.ErrEmailConfirmTokenIsNotValid, err)
	}

	if t.IsExpired() || t.IsUsed || (loggedInUserID != "" && t.UserID != loggedInUserID) {
		return errors.New(user.ErrEmailConfirmTokenIsNotValid, nil)
	}

	u, err := repos.Users.GetByID(ctx, t.UserID)
	if err != nil {
		return errors.New(user.ErrEmailConfirmTokenIsNotValid, err)
	}

	e, err := repos.Emails.GetUserEmailByID(ctx, u, t.EmailID)
	if err != nil {
		return errors.New(user.ErrEmailConfirmTokenIsNotValid, err)
	}

	if e.IsConfirmed {
		return nil
	}

	e.IsConfirmed = true
	e.ConfirmedAt = time.Now()

	if err := repos.Emails.ConfirmUserEmail(ctx, u, e); err != nil {
		return err
	}

	t.IsUsed = true
	t.UsedAt = time.Now()

	if err := repos.EmailsConfirm.UpdateUsage(ctx, t); err != nil {
		return err
	}

	return nil
}

// UpdateUser usecase
func UpdateUser(ctx context.Context, repos *repository.Repositories, usr *user.User, f string, l, i string) error {
	usr.Firstname = strings.Trim(f, " ")
	usr.Lastname = strings.Trim(l, " ")
	usr.UpdatedAt = time.Now()

	if err := repos.Users.Update(ctx, usr); err != nil {
		return err
	}

	if i != "" {
		if err := HandleAvatar(ctx, repos, usr, i); err != nil {
			return err
		}
	}
	return nil
}

// ChangePassword --
func ChangePassword(ctx context.Context, repos *repository.Repositories, usr *user.User, o string, n string) error {
	o = strings.Trim(o, " ")
	n = strings.Trim(n, " ")

	// can we find the user's password?
	p, err := repos.Users.GetPassword(ctx, usr.ID)
	if err != nil {
		return errors.New(user.ErrPasswordIsNotValid, nil)
	}

	// is the old password correct?
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(o)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New(user.ErrPasswordIsNotValid, err)
		}
		return liberr.Wrap(err, "compare password")
	}

	if !password.IsValid(n) {
		return errors.New(user.ErrWeakPassword, nil)
	}

	h, err := password.Hash(n)
	if err != nil {
		return err
	}

	if err = repos.Users.UpdatePassword(ctx, usr, h); err != nil {
		return err
	}

	return nil
}

// UpdateTheme usecase
func UpdateTheme(ctx context.Context, repos *repository.Repositories, usr *user.User, t string) error {
	usr.Theme = t
	usr.UpdatedAt = time.Now()

	if err := repos.Users.UpdateTheme(ctx, usr); err != nil {
		return err
	}

	return nil
}

// CreateUserEmail --
func CreateUserEmail(ctx context.Context, repos *repository.Repositories, u *user.User, v string) error {
	v = strings.Trim(v, " ")

	if !user.IsEmailValid(v) {
		return errors.New(user.ErrEmailIsNotValid, nil)
	}

	_, err := repos.Emails.GetEmail(ctx, v)
	if err == nil {
		return errors.New(user.ErrEmailIsAlreadyUsed, nil)
	}

	emails, err := repos.Emails.GetUserEmails(ctx, u)
	if err != nil {
		return err
	}

	if len(emails) == 1 && !emails[0].IsConfirmed {
		return errors.New(user.ErrPrimaryEmailIsNotConfirmed, nil)
	}

	email := user.NewEmail(v)
	if err := repos.Emails.CreateUserEmail(ctx, u, email); err != nil {
		return err
	}

	if err := CreateTokenAndSendConfirmationEmail(ctx, repos, u, email); err != nil {
		return err
	}

	return nil
}

// DeleteUserEmail --
func DeleteUserEmail(ctx context.Context, repos *repository.Repositories, u *user.User, v string) error {
	v = strings.Trim(v, " ")

	e, err := repos.Emails.GetUserEmailByValue(ctx, u, v)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(user.ErrEmailDoesNotExist, nil)
		}
		return err
	}

	if e.IsPrimary {
		return errors.New(user.ErrPrimaryEmailDeletion, nil)
	}

	if err := repos.Emails.DeleteUserEmail(ctx, u, e); err != nil {
		return err
	}

	return nil
}

// PrimaryUserEmail --
func PrimaryUserEmail(ctx context.Context, repos *repository.Repositories, u *user.User, v string) error {
	v = strings.Trim(v, " ")

	emails, err := repos.Emails.GetUserEmails(ctx, u)
	if err != nil {
		return err
	}

	found := false
	for _, e := range emails {
		if e.Value == v {
			if !e.IsConfirmed {
				return errors.New(user.ErrEmailIsNotConfirmed, nil)
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
		return errors.New(user.ErrEmailDoesNotExist, nil)
	}

	for _, e := range emails {
		if err := repos.Emails.PrimaryUserEmail(ctx, u, e); err != nil {
			return err
		}
	}

	return nil
}
