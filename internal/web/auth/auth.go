package auth

import (
	"context"
	"errors"
	"github.com/mickaelvieira/taipan/internal/config"
	"github.com/mickaelvieira/taipan/internal/domain/user"
)

// Authentication errors
var (
	ErrorServerIssue = errors.New("Something went wrong. Please try again later")
)

// Credentials user's credential necessary to log into the application
type Credentials struct {
	Email    string
	Password string
}

type ResetPassword struct {
	Token    string
	Password string
}

type ConfirmEmail struct {
	Token string
}

// NewContext creates a new context with the userID attached to it
func NewContext(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, config.UserContextKey, user)
}

// FromContext retrieves the userID from the context
func FromContext(ctx context.Context) *user.User {
	user, ok := ctx.Value(config.UserContextKey).(*user.User)
	if !ok {
		user = nil
	}
	return user
}
