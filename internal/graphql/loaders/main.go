package loaders

import (
	"context"

	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
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

// GetHTTPClientLogEntriesLoader get the loader
func GetHTTPClientLogEntriesLoader(repository *repository.BotlogRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		var entries []*http.Result
		for _, key := range keys {
			url, err := url.FromRawURL(key.String())
			if err != nil {
				return nil
			}
			entries, err = repository.FindByURL(ctx, url)
			if err != nil {
				return nil
			}
			results = append(results, &dataloader.Result{Data: entries})
		}
		return results
	})
}
