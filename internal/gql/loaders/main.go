package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// Loaders helps interact the various dataloaders
type Loaders struct {
	Repositories *repository.Repositories
}

// GetBookmarksLoader get the loader
func (l *Loaders) GetBookmarksLoader() *dataloader.Loader {
	log.Println("get bookmark loader")

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		documents, err := l.Repositories.Documents.GetByIDs(ctx, keys.Keys())

		if err != nil {
			log.Fatal(err)
		}

		for _, document := range documents {
			var result = &dataloader.Result{Data: document}
			results = append(results, result)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	return loader
}
