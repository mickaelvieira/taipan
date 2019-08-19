package user

import (
	"errors"
	"github/mickaelvieira/taipan/internal/nonce"
	"strings"
	"time"
)

// User domain errors
var (
	ErrDoesNotExist                 = errors.New("User does not exist")
	ErrWeakPassword                 = errors.New("Your password must be at least 10 characters long")
	ErrEmailIsNotValid              = errors.New("Your email does not seem to be valid")
	ErrEmailDoesNotExist            = errors.New("Email does not exist")
	ErrEmailIsAlreadyUsed           = errors.New("There is already an account associated to this email address")
	ErrCredentialsAreNotValid       = errors.New("Email or password does not match any records in our database")
	ErrPasswordIsNotValid           = errors.New("Your password is not correct") // only for password change
	ErrPrimaryEmailDeletion         = errors.New("You cannot delete your primary email address")
	ErrPrimaryEmailIsNotConfirmed   = errors.New("You cannot add new email address before confirming your primary email")
	ErrEmailIsNotConfirmed          = errors.New("You cannot use this address as your primary email since it hasn't been unconfirmed yet")
	ErrPasswordResetTokenIsNotValid = errors.New("Your reset token does seem to be valid")
	ErrEmailConfirmTokenIsNotValid  = errors.New("Your reset token does seem to be valid")
)

// User represents a single user within the application
type User struct {
	ID        string
	Emails    []*Email
	Firstname string
	Lastname  string
	Image     *Image
	Theme     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HasImage determine whether the user has an image associated to it
func (u *User) HasImage() bool {
	return u.Image != nil
}

// Image represents a user's avatar
type Image struct {
	Name   string
	Width  int32
	Height int32
	Format string
}

// SetDimensions image's information
func (i *Image) SetDimensions(w int, h int) {
	i.Width = int32(w)
	i.Height = int32(h)
}

// NewImage returns a document's image
func NewImage(name string, width int32, height int32, format string) *Image {
	return &Image{
		Name:   name,
		Width:  width,
		Height: height,
		Format: format,
	}
}

// Email user's email
type Email struct {
	ID          string
	Value       string
	IsPrimary   bool
	IsConfirmed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ConfirmedAt time.Time
}

// IsEmailValid is the email valid?
func IsEmailValid(e string) bool {
	return len(e) > 0 && strings.Contains(e, "@")
}

// NewEmail creates a new user email address
func NewEmail(value string) *Email {
	return &Email{
		Value:     value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// EmailConfirmToken reset password entry
type EmailConfirmToken struct {
	Token     string
	IsUsed    bool
	UserID    string
	EmailID   string
	CreatedAt time.Time
	ExpiredAt time.Time
	UsedAt    time.Time
}

// IsExpired is the roken expired?
func (r *EmailConfirmToken) IsExpired() bool {
	now := time.Now().UTC()
	return now.After(r.ExpiredAt)
}

const tokenNonceLen = 30
const tokenValidityInHours = 168 // 7 days

// getExpireAt calculates the date when the token will be expired
func getExpireAt() time.Time {
	start := time.Now().UTC()
	return start.Add(time.Hour * tokenValidityInHours)
}

// NewEmailConfirmToken create a new password reset entry
func NewEmailConfirmToken(u *User, e *Email) *EmailConfirmToken {
	return &EmailConfirmToken{
		Token:     nonce.Generate(tokenNonceLen),
		UserID:    u.ID,
		EmailID:   e.ID,
		CreatedAt: time.Now(),
		ExpiredAt: getExpireAt(),
	}
}
