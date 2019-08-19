package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
)

// SyndicationRootResolver syndication's root resolver
type SyndicationRootResolver struct {
	repositories *repository.Repositories
}

// SourceCollection resolver
type SourceCollection struct {
	Results []*Source
	Total   int32
	Offset  int32
	Limit   int32
}

// Source resolves the bookmark entity
type Source struct {
	source     *syndication.Source
	repository *repository.Repositories
	logLoader  *dataloader.Loader
}

// ID resolves the ID field
func (r *Source) ID() gql.ID {
	return gql.ID(r.source.ID)
}

// URL resolves the URL field
func (r *Source) URL() scalars.URL {
	return scalars.NewURL(r.source.URL)
}

// Domain resolves the Domain field
func (r *Source) Domain() *scalars.URL {
	d := scalars.NewURL(r.source.Domain)
	return &d
}

// Title resolves the Title field
func (r *Source) Title() string {
	return r.source.Title
}

// Type resolves the Type field
func (r *Source) Type() string {
	return string(r.source.Type)
}

// Frequency resolves the Frequency field
func (r *Source) Frequency() string {
	return string(r.source.Frequency)
}

// IsPaused resolves the IsPaused field
func (r *Source) IsPaused() bool {
	return r.source.IsPaused
}

// IsDeleted resolves the IsDeleted field
func (r *Source) IsDeleted() bool {
	return r.source.IsDeleted
}

// CreatedAt resolves the CreatedAt field
func (r *Source) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.source.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *Source) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.source.UpdatedAt)
}

// ParsedAt resolves the ParsedAt field
func (r *Source) ParsedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.source.ParsedAt)
	return &t
}

// LogEntries returns the document's parser log
func (r *Source) LogEntries(ctx context.Context) (*[]*Log, error) {
	data, err := r.logLoader.Load(ctx, dataloader.StringKey(r.source.URL.String()))()
	if err != nil {
		return nil, err
	}

	results, ok := data.([]*http.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}

	res := resolve(r.repository).logs(results)

	return &res, nil
}

// Source returns the syndication source
func (r *SyndicationRootResolver) Source(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	u := args.URL.ToDomain()

	s, err := r.repositories.Syndication.GetByURL(ctx, u)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Create adds a syndication source
func (r *SyndicationRootResolver) Create(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	u := args.URL.ToDomain()

	s, err := usecase.CreateSyndicationSource(ctx, r.repositories, u, true)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// UpdateTitle adds a syndication source
func (r *SyndicationRootResolver) UpdateTitle(ctx context.Context, args struct {
	URL   scalars.URL
	Title string
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.UpdateSourceTitle(ctx, r.repositories, s, args.Title)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Pause disables a syndication source
func (r *SyndicationRootResolver) Pause(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.PauseSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Resume enables a syndication source
func (r *SyndicationRootResolver) Resume(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.ResumeSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Enable enables a syndication source
func (r *SyndicationRootResolver) Enable(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.EnableSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Disable disables a syndication source
func (r *SyndicationRootResolver) Disable(ctx context.Context, args struct {
	URL scalars.URL
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByURL(ctx, args.URL.ToDomain())
	if err != nil {
		return nil, err
	}

	err = usecase.DisableSyndicationSource(ctx, r.repositories, s)
	if err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Sources resolves the query
func (r *SyndicationRootResolver) Sources(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     searchSourcesInput
}) (*SourceCollection, error) {
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

	res := SourceCollection{
		Results: resolve(r.repositories).sources(results),
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
