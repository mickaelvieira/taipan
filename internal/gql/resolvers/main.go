package resolvers

import (
	"context"
	userid "github/mickaelvieira/taipan/internal/context"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/repository"
)

// Resolvers resolvers
type Resolvers struct {
	Dataloaders  *loaders.Loaders
	Repositories *repository.Repositories
}

func (r *Resolvers) getUser(ctx context.Context) (*user.User, error) {
	userID := userid.FromContext(ctx)
	return r.Repositories.Users.GetByID(ctx, userID)
}
