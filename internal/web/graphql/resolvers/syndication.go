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
	s  *syndication.Source
	r  *repository.Repositories
	ll *dataloader.Loader
}

// ID resolves the ID field
func (r *SourceResolver) ID() gql.ID {
	return gql.ID(r.s.ID)
}

// URL resolves the URL field
func (r *SourceResolver) URL() scalars.URL {
	return scalars.NewURL(r.s.URL)
}

// Domain resolves the Domain field
func (r *SourceResolver) Domain() *scalars.URL {
	d := scalars.NewURL(r.s.Domain)
	return &d
}

// Title resolves the Title field
func (r *SourceResolver) Title() string {
	return r.s.Title
}

// Type resolves the Type field
func (r *SourceResolver) Type() string {
	return string(r.s.Type)
}

// Frequency resolves the Frequency field
func (r *SourceResolver) Frequency() string {
	return string(r.s.Frequency)
}

// IsPaused resolves the IsPaused field
func (r *SourceResolver) IsPaused() bool {
	return r.s.IsPaused
}

// IsDeleted resolves the IsDeleted field
func (r *SourceResolver) IsDeleted() bool {
	return r.s.IsDeleted
}

// CreatedAt resolves the CreatedAt field
func (r *SourceResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.s.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *SourceResolver) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.s.UpdatedAt)
}

// ParsedAt resolves the ParsedAt field
func (r *SourceResolver) ParsedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.s.ParsedAt)
	return &t
}

// LogEntries returns the document's parser log
func (r *SourceResolver) LogEntries(ctx context.Context) (*[]*LogResolver, error) {
	data, err := r.ll.Load(ctx, dataloader.StringKey(r.s.URL.String()))()
	if err != nil {
		return nil, err
	}

	results, ok := data.([]*http.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}

	res := resolve(r.r).logs(results)

	return &res, nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	return resolve(r.repositories).source(s), nil
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

	res := SourceCollectionResolver{
		Results: resolve(r.repositories).sources(results),
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
