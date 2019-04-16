package graphql

import (
	"context"

	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

var repo = repository.NewGpxRepository()

// GetActivityLoader get the loader
func GetActivityLoader() *dataloader.Loader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		var data = repo.FindOne()
		var result = &dataloader.Result{Data: data}

		results = append(results, result)

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	return loader
}
