package usecase

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/errors"
	"github/mickaelvieira/taipan/internal/domain/password"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"time"

	liberr "github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

// Signin --
func Signin(ctx context.Context, repos *repository.Repositories, e string, pwd string) (*user.User, error) {
	// can we find the user?
	u, err := repos.Users.GetByPrimaryEmail(ctx, e)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, err
	}

	// can we find the user's password?
	p, err := repos.Users.GetPassword(ctx, u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, err
	}

	// do the password match?
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(pwd)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.New(user.ErrCredentialsAreNotValid, err)
		}
		return nil, liberr.Wrap(err, "compare password")
	}

	return u, nil
}

// Signup --
func Signup(ctx context.Context, repos *repository.Repositories, e string, p string) (*user.User, error) {
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
	_, err = repos.ResetToken.FindUserActiveToken(ctx, u)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		err = repos.ResetToken.Create(ctx, password.NewResetPasswordToken(u))
		if err != nil {
			return err
		}
	}

	return nil
}

// ResetPassword --
func ResetPassword(ctx context.Context, repos *repository.Repositories, t string, p string) error {
	if !password.IsValid(p) {
		return errors.New(user.ErrWeakPassword, nil)
	}

	h, err := password.Hash(p)
	if err != nil {
		return err
	}

	e, err := repos.ResetToken.GetToken(ctx, t)
	if err != nil {
		return errors.New(user.ErrResetTokenIsNotValid, err)
	}

	if e.IsExpired() || e.IsUsed {
		return errors.New(user.ErrResetTokenIsNotValid, nil)
	}

	u, err := repos.Users.GetByID(ctx, e.UserID)
	if err != nil {
		return errors.New(user.ErrResetTokenIsNotValid, err)
	}

	if err = repos.Users.UpdatePassword(ctx, u, h); err != nil {
		return err
	}

	e.IsUsed = true
	e.UsedAt = time.Now()

	if err := repos.ResetToken.UpdateUsage(ctx, e); err != nil {
		return err
	}

	return nil
}

// UpdateUser usecase
func UpdateUser(ctx context.Context, repos *repository.Repositories, usr *user.User, f string, l, i string) error {
	usr.Firstname = f
	usr.Lastname = l
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
func CreateUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	_, err := repos.Emails.GetEmail(ctx, v)
	if err == nil {
		return errors.New(user.ErrEmailIsAlreadyUsed, nil)
	}

	emails, err := repos.Emails.GetUserEmails(ctx, usr)
	if err != nil {
		return err
	}

	if len(emails) == 1 && !emails[0].IsConfirmed {
		return errors.New(user.ErrPrimaryEmailIsNotConfirmed, nil)
	}

	if err := repos.Emails.CreateUserEmail(ctx, usr, user.NewEmail(v)); err != nil {
		return err
	}

	return nil
}

// DeleteUserEmail --
func DeleteUserEmail(ctx context.Context, repos *repository.Repositories, usr *user.User, v string) error {
	e, err := repos.Emails.GetUserEmail(ctx, usr, v)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New(user.ErrEmailDoesNotExist, nil)
		}
		return err
	}

	if e.IsPrimary {
		return errors.New(user.ErrPrimaryEmailDeletion, nil)
	}

	if err := repos.Emails.DeleteUserEmail(ctx, usr, e); err != nil {
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
		if err := repos.Emails.PrimaryUserEmail(ctx, usr, e); err != nil {
			return err
		}
	}

	return nil
}
