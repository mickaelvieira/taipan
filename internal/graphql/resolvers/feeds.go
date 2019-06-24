package resolvers

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/graphql/loaders"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// FeedCollectionResolver resolver
type FeedCollectionResolver struct {
	Results []*FeedResolver
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
func (r *FeedResolver) ID() gql.ID {
	return gql.ID(r.Feed.ID)
}

// URL resolves the URL
func (r *FeedResolver) URL() scalars.URL {
	return scalars.URL{URL: r.Feed.URL}
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
func (r *FeedResolver) CreatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Feed.CreatedAt}
}

// UpdatedAt resolves the UpdatedAt field
func (r *FeedResolver) UpdatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Feed.UpdatedAt}
}

// ParsedAt resolves the ParsedAt field
func (r *FeedResolver) ParsedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Feed.ParsedAt}
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

// Feed adds a feed
func (r *RootResolver) Feed(ctx context.Context, args struct {
	URL scalars.URL
}) (*FeedResolver, error) {
	u := args.URL.URL
	if feed.IsBlacklisted(u.String()) {
		return nil, fmt.Errorf("URL %s is blacklisted", args.URL)
	}

	f, err := r.repositories.Feeds.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			f = feed.New(u, "", "")
			err = r.repositories.Feeds.Insert(ctx, f)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// @TODO push URLs to the queue
	_, err = usecase.ParseFeed(ctx, f, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &FeedResolver{Feed: f}

	return res, nil
}

// Feeds resolves the query
func (r *RootResolver) Feeds(ctx context.Context, args struct {
	Pagination OffsetPaginationInput
}) (*FeedCollectionResolver, error) {
	fromArgs := GetOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)

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
		Results: feeds,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}
