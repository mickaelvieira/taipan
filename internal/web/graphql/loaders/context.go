package loaders

import (
	"context"
	"github/mickaelvieira/taipan/internal/config"
)

// NewContext creates a new context containing the dataloaders
func NewContext(ctx context.Context, loaders *Loaders) context.Context {
	return context.WithValue(ctx, config.LoadersContextKey, loaders)
}

// FromContext retrieves the dataloaders from the context
func FromContext(ctx context.Context) *Loaders {
	loaders, ok := ctx.Value(config.LoadersContextKey).(*Loaders)
	if !ok {
		return nil
	}
	return loaders
}
