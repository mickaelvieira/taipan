package resolvers

import (
	"context"
	"database/sql"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/graphql/loaders"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// SyndicationResolver syndication's root resolver
type SyndicationResolver struct {
	repositories *repository.Repositories
}

// SourceCollectionResolver resolver
type SourceCollectionResolver struct {
	Results []*SourceResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// SourceResolver resolves the bookmark entity
type SourceResolver struct {
	*syndication.Source
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *SourceResolver) ID() gql.ID {
	return gql.ID(r.Source.ID)
}

// URL resolves the URL
func (r *SourceResolver) URL() scalars.URL {
	return scalars.URL{URL: r.Source.URL}
}

// Title resolves the Title field
func (r *SourceResolver) Title() string {
	return r.Source.Title
}

// Type resolves the Type field
func (r *SourceResolver) Type() string {
	return string(r.Source.Type)
}

// Status resolves the Status field
func (r *SourceResolver) Status() string {
	return string(r.Source.Status)
}

// IsPaused resolves the IsPaused field
func (r *SourceResolver) IsPaused() bool {
	return r.Source.IsPaused
}

// CreatedAt resolves the CreatedAt field
func (r *SourceResolver) CreatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.CreatedAt}
}

// UpdatedAt resolves the UpdatedAt field
func (r *SourceResolver) UpdatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.UpdatedAt}
}

// ParsedAt resolves the ParsedAt field
func (r *SourceResolver) ParsedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Source.ParsedAt}
}

// LogEntries returns the document's parser log
func (r *SourceResolver) LogEntries(ctx context.Context) (*[]*HTTPClientLogResolver, error) {
	data, err := r.getLogsLoader().Load(ctx, dataloader.StringKey(r.Source.URL.String()))()
	if err != nil {
		return nil, err
	}
	results, ok := data.([]*http.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}
	var resolvers []*HTTPClientLogResolver
	for _, result := range results {
		resolvers = append(resolvers, &HTTPClientLogResolver{Result: result})
	}
	return &resolvers, nil
}

func (r *SourceResolver) getLogsLoader() *dataloader.Loader {
	return loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
}

// Source adds a syndication source
func (r *SyndicationResolver) Source(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.URL
	if syndication.IsBlacklisted(u.String()) {
		return nil, fmt.Errorf("URL %s is blacklisted", args.URL)
	}

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		if err == sql.ErrNoRows {
			s = syndication.NewSource(u, "", "")
			err = r.repositories.Syndication.Insert(ctx, s)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// @TODO push URLs to the queue
	_, err = usecase.ParseSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Disable disables a syndication source
func (r *SyndicationResolver) Disable(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.URL

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	err = usecase.DisableSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Enable enables a syndication source
func (r *SyndicationResolver) Enable(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.URL

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	err = usecase.EnableSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Delete deletes a syndication source
func (r *SyndicationResolver) Delete(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.URL

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	err = usecase.DeleteSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Sources resolves the query
func (r *SyndicationResolver) Sources(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     searchSourcesInput
}) (*SourceCollectionResolver, error) {
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)
	paused := args.Search.IsPaused

	results, err := r.repositories.Syndication.FindAll(ctx, paused, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Syndication.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var sources []*SourceResolver
	for _, result := range results {
		sources = append(sources, &SourceResolver{
			Source:       result,
			repositories: r.repositories,
		})
	}

	reso := SourceCollectionResolver{
		Results: sources,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}
