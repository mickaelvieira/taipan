package clientid

import (
	"context"
	"github/mickaelvieira/taipan/internal/config"
	"net/http"
	"os"
)

// @TODO sign and validate the client ID instead to strenghen security
// The resulting token should be forged with some kind of session ID and timestamp
// https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#synchronizer-token-pattern

// NewContext creates a new context with the userID attached to it
func NewContext(ctx context.Context, clientID string) context.Context {
	return context.WithValue(ctx, config.ClientIDContextKey, clientID)
}

// FromContext retrieves the userID from the context
func FromContext(ctx context.Context) string {
	clientID, ok := ctx.Value(config.ClientIDContextKey).(string)
	if !ok {
		clientID = ""
	}
	return clientID
}

// WithClientID stores the client (aka user agent) ID in the context
func WithClientID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		clientID := req.Header.Get(os.Getenv("APP_CLIENT_ID_HEADER"))
		ctx = NewContext(ctx, clientID)
		next.ServeHTTP(w, req.WithContext(ctx))
	}
}
