package loaders

import (
	"context"
	"log"

	"github/mickaelvieira/taipan/internal/aggregator"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
)

// Loaders --
type Loaders struct {
	Users      *dataloader.Loader
	UsersStats *dataloader.Loader
	Emails     *dataloader.Loader
	Sources    *dataloader.Loader
	Logs       *dataloader.Loader
}

// NewDataloaders --
func NewDataloaders(r *repository.Repositories) *Loaders {
	return &Loaders{
		Users:      getUserLoaders(r.Users),
		UsersStats: getUserStatsLoaders(r),
		Emails:     getEmailsLoaders(r.Emails),
		Sources:    getSourcesLoaders(r.Syndication),
		Logs:       getLogsLoader(r.Botlogs),
	}
}

// getDocumentsLoader get the loader
func getDocumentsLoader(repository *repository.DocumentRepository) *dataloader.Loader {
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

// getSourcesLoaders get the syndication source loader
func getSourcesLoaders(repository *repository.SyndicationRepository) *dataloader.Loader {
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

// getEmailsLoaders loads user's emails
func getEmailsLoaders(repository *repository.UserEmailRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			u, ok := key.(*user.User)
			if !ok {
				log.Fatalln("Key must a user")
			}

			e, err := repository.GetUserEmails(ctx, u)
			if err != nil {
				log.Fatalln(err)
			}

			results[i] = &dataloader.Result{Data: e}
		}
		return results
	})
}

// getUserLoaders loads users
func getUserLoaders(repository *repository.UserRepository) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		users, err := repository.GetByIDs(ctx, keys.Keys())
		if err != nil {
			log.Fatalln(err)
		}

		results := make([]*dataloader.Result, len(users))
		for i, u := range users {
			results[i] = &dataloader.Result{Data: u}
		}

		return results
	})
}

// getEmailsLoaders loads user's emails
func getUserStatsLoaders(repositories *repository.Repositories) *dataloader.Loader {
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			k, ok := key.(*aggregator.LoaderKey)
			if !ok {
				log.Fatalln("Key must an aggregator key")
			}

			t, err := aggregator.Aggregate(ctx, repositories, k.User, k.Type)
			if err != nil {
				return nil
			}

			results[i] = &dataloader.Result{Data: t}
		}
		return results
	})
}

// getLogsLoader get the loader
func getLogsLoader(repository *repository.BotlogRepository) *dataloader.Loader {
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
