package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// GetDocumentLoader get the loader
func GetDocumentLoader(repository *repository.DocumentRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		documents, err := repository.GetByIDs(ctx, keys.Keys())
		if err != nil {
			log.Fatalln(err)
		}

		results := make([]*dataloader.Result, len(documents))
		for i, d := range documents {
			results[i] = &dataloader.Result{Data: d}
		}

		return results
	})
}

// GetSource get the syndication source loader
func GetSource(repository *repository.SyndicationRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		sources, err := repository.GetByIDs(ctx, keys.Keys())
		if err != nil {
			log.Fatalln(err)
		}

		results := make([]*dataloader.Result, len(sources))
		for i, s := range sources {
			results[i] = &dataloader.Result{Data: s}
		}

		return results
	})
}

// GetLogs get the loader
func GetLogs(repository *repository.BotlogRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		for _, key := range keys {
			url, err := url.FromRawURL(key.String())
			if err != nil {
				return nil
			}

			entries, err := repository.FindByURL(ctx, url)
			if err != nil {
				return nil
			}

			results = append(results, &dataloader.Result{Data: entries})
		}
		return results
	})
}
