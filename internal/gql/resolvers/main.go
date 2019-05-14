package resolvers

import (
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/repository"
)

// Resolvers resolvers
type Resolvers struct {
	Dataloaders  *loaders.Loaders
	Repositories *repository.Repositories
}
