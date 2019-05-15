package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

// FeedCollectionResolver resolver
type FeedCollectionResolver struct {
	Results *[]*FeedResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// FeedResolver resolves the bookmark entity
type FeedResolver struct {
	*feed.Feed
	logsLoader *dataloader.Loader
}

// ID resolves the ID field
func (r *FeedResolver) ID() graphql.ID {
	return graphql.ID(r.Feed.ID)
}

// URL resolves the URL
func (r *FeedResolver) URL() string {
	return r.Feed.URL.String()
}

// Title resolves the Title field
func (r *FeedResolver) Title() string {
	return r.Feed.Title
}

// Type resolves the Type field
func (r *FeedResolver) Type() string {
	return string(r.Feed.Type)
}

// Status resolves the Status field
func (r *FeedResolver) Status() string {
	return string(r.Feed.Status)
}

// CreatedAt resolves the CreatedAt field
func (r *FeedResolver) CreatedAt() string {
	return r.Feed.CreatedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *FeedResolver) UpdatedAt() string {
	return r.Feed.UpdatedAt.Format(time.RFC3339)
}

// LogEntries returns the document's parser log
func (r *FeedResolver) LogEntries(ctx context.Context) (*[]*HTTPClientLogResolver, error) {
	data, err := r.logsLoader.Load(ctx, dataloader.StringKey(r.Feed.URL.String()))()
	if err != nil {
		return nil, err
	}
	results, ok := data.([]*client.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}
	var resolvers []*HTTPClientLogResolver
	for _, result := range results {
		resolvers = append(resolvers, &HTTPClientLogResolver{Result: result})
	}
	return &resolvers, nil
}

// Feeds resolves the query
func (r *Resolvers) Feeds(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*FeedCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(10)
	offset, limit := fromArgs(args.Offset, args.Limit)

	results, err := r.repositories.Feeds.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Feeds.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var feeds []*FeedResolver
	var logsLoader = loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
	for _, result := range results {
		feeds = append(feeds, &FeedResolver{
			Feed:       result,
			logsLoader: logsLoader,
		})
	}

	reso := FeedCollectionResolver{
		Results: &feeds,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}
