package auth

import (
	"context"
	"database/sql"
	"errors"
	"github/mickaelvieira/taipan/internal/config"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"net/http"

	"github.com/go-session/session"
)

// Authentication errors
var (
	ErrorInvalidCreds = errors.New("Username or password does not match any records in our database")
	ErrorServerIssue  = errors.New("Something went wrong. Please try again later")
)

// Credentials user's credential necessary to log into the application
type Credentials struct {
	Username string
	Password string
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

// WithUser stores the user ID in the context
func WithUser(next http.HandlerFunc, repository *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get the current context or create a new one
		ctx := req.Context()
		store, err := session.Start(ctx, w, req)
		if err != nil {
			log.Fatal(err)
		}

		value, ok := store.Get("user_id")
		if !ok {
			w.WriteHeader(401)
			return
		}

		userID, ok := value.(string)
		if !ok {
			w.WriteHeader(401)
			return
		}

		user, err := repository.GetByID(ctx, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(401)
			} else {
				log.Fatal(err)
			}
		} else {
			ctx = NewContext(ctx, user) // @TODO not sure it is a good idea to keep the user in the context
			next.ServeHTTP(w, req.WithContext(ctx))
		}
	}
}
