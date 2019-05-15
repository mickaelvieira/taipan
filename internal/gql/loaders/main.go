package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// Loaders helps interact the various dataloaders
type Loaders struct {
	Repositories   *repository.Repositories
	Documents      *dataloader.Loader
	DocumentsFeeds *dataloader.Loader
}

// GetDocumentLoader get the loader
func GetDocumentLoader(repository *repository.DocumentRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		documents, err := repository.GetByIDs(ctx, keys.Keys())
		if err != nil {
			return nil
		}
		for _, document := range documents {
			results = append(results, &dataloader.Result{Data: document})
		}
		return results
	})
}

// GetDocumentsFeedsLoader get the loader
func GetDocumentsFeedsLoader(repository *repository.FeedRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		log.Println("run batch feeds")
		var results []*dataloader.Result
		for _, key := range keys {
			doc, ok := key.Raw().(*document.Document)
			if !ok {
				return nil
			}
			feeds, err := repository.GetDocumentFeeds(ctx, doc)
			if err != nil {
				return nil
			}
			results = append(results, &dataloader.Result{Data: feeds})
		}
		return results
	})
}

// GetHTTPClientLogEntriesLoader get the loader
func GetHTTPClientLogEntriesLoader(repository *repository.BotlogRepository) *dataloader.Loader {
	log.Println("=>>>>> GetHTTPClientLogEntriesLoader")
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		log.Println("run batch logs")
		log.Println(keys.Keys())
		var results []*dataloader.Result
		for _, key := range keys {
			entries, err := repository.FindByURI(ctx, key.String())
			if err != nil {
				return nil
			}
			results = append(results, &dataloader.Result{Data: entries})
		}
		return results
	})
}
