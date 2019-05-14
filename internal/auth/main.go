package auth

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
	"net/http"
)

type key int

// https://blog.golang.org/context#TOC_3.2.
const (
	UserKey key = iota
)

// NewContext creates a new context with the userID attached to it
func NewContext(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

// FromContext retrieves the userID from the context
func FromContext(ctx context.Context) *user.User {
	user, ok := ctx.Value(UserKey).(*user.User)
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
		// @TODO I need to implement authentication
		user, err := repository.GetByID(ctx, "1")
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
