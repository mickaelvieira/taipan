package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// GetBookmarksLoader get the loader
func GetBookmarksLoader() *dataloader.Loader {
	var repository = repository.NewBookmarkRepository()

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {

		log.Println("BATCH")
		var results []*dataloader.Result
		var bookmarks = repository.FindAll(ctx, keys.Keys())

		log.Println(keys.Keys())

		log.Println(len(bookmarks))

		for _, bookmark := range bookmarks {
			var result = &dataloader.Result{Data: bookmark}
			results = append(results, result)
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	return loader
}
