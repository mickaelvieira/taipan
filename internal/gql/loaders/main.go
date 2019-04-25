package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// Loaders helps interact the various dataloaders
type Loaders struct{}

// GetBookmarksLoader get the loader
func (l *Loaders) GetBookmarksLoader() *dataloader.Loader {
	log.Println("get bookmark loader")
	var repository = repository.NewBookmarkRepository()

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		var bookmarks = repository.GetByIDs(ctx, keys.Keys())

		for _, bookmark := range bookmarks {
			var result = &dataloader.Result{Data: bookmark}
			results = append(results, result)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	return loader
}
