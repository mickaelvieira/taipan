package resolvers

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/graphql/loaders"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// SyndicationCollectionResolver resolver
type SyndicationCollectionResolver struct {
	Results []*SyndicationResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// SyndicationResolver resolves the bookmark entity
type SyndicationResolver struct {
	*syndication.Source
	logsLoader *dataloader.Loader
}

// ID resolves the ID field
func (r *SyndicationResolver) ID() gql.ID {
	return gql.ID(r.Source.ID)
}

// URL resolves the URL
func (r *SyndicationResolver) URL() scalars.URL {
	return scalars.URL{URL: r.Source.URL}
}

// Title resolves the Title field
func (r *SyndicationResolver) Title() string {
	return r.Source.Title
}

// Type resolves the Type field
func (r *SyndicationResolver) Type() string {
	return string(r.Source.Type)
}

// Status resolves the Status field
func (r *SyndicationResolver) Status() string {
	return string(r.Source.Status)
}

// CreatedAt resolves the CreatedAt field
func (r *SyndicationResolver) CreatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.CreatedAt}
}

// UpdatedAt resolves the UpdatedAt field
func (r *SyndicationResolver) UpdatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.UpdatedAt}
}

// ParsedAt resolves the ParsedAt field
func (r *SyndicationResolver) ParsedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.ParsedAt}
}

// LogEntries returns the document's parser log
func (r *SyndicationResolver) LogEntries(ctx context.Context) (*[]*HTTPClientLogResolver, error) {
	data, err := r.logsLoader.Load(ctx, dataloader.StringKey(r.Source.URL.String()))()
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
}) (*SyndicationResolver, error) {
	u := args.URL.URL
	if syndication.IsBlacklisted(u.String()) {
		return nil, fmt.Errorf("URL %s is blacklisted", args.URL)
	}

	f, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			f = syndication.NewSource(u, "", "")
			err = r.repositories.Syndication.Insert(ctx, f)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// @TODO push URLs to the queue
	_, err = usecase.ParseSyndicationSource(ctx, f, r.repositories)
	if err != nil {
		return nil, err
	}

	res := &SyndicationResolver{Source: f}

	return res, nil
}

// Feeds resolves the query
func (r *RootResolver) Feeds(ctx context.Context, args struct {
	Pagination OffsetPaginationInput
}) (*SyndicationCollectionResolver, error) {
	fromArgs := GetOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)

	results, err := r.repositories.Syndication.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Syndication.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var sources []*SyndicationResolver
	var logsLoader = loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
	for _, result := range results {
		sources = append(sources, &SyndicationResolver{
			Source:     result,
			logsLoader: logsLoader,
		})
	}

	reso := SyndicationCollectionResolver{
		Results: sources,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}
