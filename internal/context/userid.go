package userid

import (
	"context"
)

type key int

// https://blog.golang.org/context#TOC_3.2.
const (
	UserIDKey key = iota
)

// NewContext creates a new context with the userID attached to it
func NewContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// FromContext retrieves the userID from the context
func FromContext(ctx context.Context) string {
	userIP, ok := ctx.Value(UserIDKey).(string)

	if !ok {
		userIP = ""
	}

	return userIP
}
