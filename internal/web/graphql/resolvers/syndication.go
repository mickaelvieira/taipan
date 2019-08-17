package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/graphql/loaders"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

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

// URL resolves the URL field
func (r *SourceResolver) URL() scalars.URL {
	return scalars.NewURL(r.Source.URL)
}

// Domain resolves the Domain field
func (r *SourceResolver) Domain() *scalars.URL {
	d := scalars.NewURL(r.Source.Domain)
	return &d
}

// Title resolves the Title field
func (r *SourceResolver) Title() string {
	return r.Source.Title
}

// Type resolves the Type field
func (r *SourceResolver) Type() string {
	return string(r.Source.Type)
}

// Frequency resolves the Frequency field
func (r *SourceResolver) Frequency() string {
	return string(r.Source.Frequency)
}

// IsPaused resolves the IsPaused field
func (r *SourceResolver) IsPaused() bool {
	return r.Source.IsPaused
}

// IsDeleted resolves the IsDeleted field
func (r *SourceResolver) IsDeleted() bool {
	return r.Source.IsDeleted
}

// CreatedAt resolves the CreatedAt field
func (r *SourceResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.Source.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *SourceResolver) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.Source.UpdatedAt)
}

// ParsedAt resolves the ParsedAt field
func (r *SourceResolver) ParsedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.Source.ParsedAt)
	return &t
}

// LogEntries returns the document's parser log
func (r *SourceResolver) LogEntries(ctx context.Context) (*[]*LogResolver, error) {
	data, err := r.getLogsLoader().Load(ctx, dataloader.StringKey(r.Source.URL.String()))()
	if err != nil {
		return nil, err
	}
	results, ok := data.([]*http.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}
	var resolvers []*LogResolver
	for _, result := range results {
		resolvers = append(resolvers, &LogResolver{Result: result})
	}
	return &resolvers, nil
}

func (r *SourceResolver) getLogsLoader() *dataloader.Loader {
	return loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
}

// Source returns the syndication source
func (r *SyndicationResolver) Source(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.ToDomain()

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s, repositories: r.repositories}

	return res, nil
}

// Create adds a syndication source
func (r *SyndicationResolver) Create(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	u := args.URL.ToDomain()

	s, err := usecase.CreateSyndicationSource(ctx, r.repositories, u)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// UpdateTitle adds a syndication source
func (r *SyndicationResolver) UpdateTitle(ctx context.Context, args struct {
	URL   scalars.URL
	Title string
}) (*SourceResolver, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.UpdateSourceTitle(ctx, r.repositories, s, args.Title)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Pause disables a syndication source
func (r *SyndicationResolver) Pause(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.PauseSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	res := &SourceResolver{Source: s}

	return res, nil
}

// Resume enables a syndication source
func (r *SyndicationResolver) Resume(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.ResumeSyndicationSource(ctx, r.repositories, s)
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
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
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

// Disable disables a syndication source
func (r *SyndicationResolver) Disable(ctx context.Context, args struct {
	URL scalars.URL
}) (*SourceResolver, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
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

	res := SourceCollectionResolver{
		Results: sources,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
