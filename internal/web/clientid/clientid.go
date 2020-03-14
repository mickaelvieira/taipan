package clientid

import (
	"context"
	"github.com/mickaelvieira/taipan/internal/config"
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
