package resolvers

import (
	"github/mickaelvieira/taipan/internal/repository"
)

// Resolvers resolvers
type Resolvers struct {
	repositories *repository.Repositories
}

// GetRootResolver returns the root resolver. Queries and mutations are methods of this resolver
func GetRootResolver(repositories *repository.Repositories) *Resolvers {
	return &Resolvers{
		repositories: repositories,
	}
}
