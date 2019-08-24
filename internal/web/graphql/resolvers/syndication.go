package resolvers

import (
	"context"
	"database/sql"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/graphql/loaders"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"
	"time"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
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

// Source resolves the syndication source entity
type Source struct {
	source       *syndication.Source
	repositories *repository.Repositories
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

// Tags resolves the Tags field
func (r *Source) Tags(ctx context.Context) ([]*Tag, error) {
	l := loaders.FromContext(ctx)

	ids, err := r.repositories.SyndicationTags.GetSourceTagIDs(ctx, r.source)
	if err != nil {
		return nil, err
	}

	var keys = make([]dataloader.Key, len(ids))
	for i, id := range ids {
		keys[i] = dataloader.StringKey(id)
	}

	future := l.SyndicationTag.LoadMany(ctx, keys)
	data, e := future()
	if len(e) > 0 {
		return nil, e[0]
	}

	var resolver = resolve(r.repositories)
	var tags = make([]*Tag, len(data))

	for i, datum := range data {
		result, ok := datum.(*syndication.Tag)
		if !ok {
			return nil, ErrDataTypeIsNotValid
		}
		tags[i] = resolver.tag(result)
	}

	return tags, nil
}

// LogEntries returns the document's parser log
func (r *Source) LogEntries(ctx context.Context) (*[]*Log, error) {
	l := loaders.FromContext(ctx)
	if l == nil {
		return nil, ErrLoadersNotFound
	}

	d, err := l.Logs.Load(ctx, r.source.URL)()
	if err != nil {
		return nil, err
	}

	results, ok := d.([]*http.Result)
	if !ok {
		return nil, ErrDataTypeIsNotValid
	}

	res := resolve(r.repositories).logs(results)

	return &res, nil
}

// TagCollection resolver
type TagCollection struct {
	Results []*Tag
	Total   int32
}

// Tag resolves the syndication tag entity
type Tag struct {
	tag *syndication.Tag
}

// ID resolves the ID field
func (r *Tag) ID() gql.ID {
	return gql.ID(r.tag.ID)
}

// Label resolves the Tag field
func (r *Tag) Label() string {
	return r.tag.Label
}

// CreatedAt resolves the CreatedAt field
func (r *Tag) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.tag.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *Tag) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.tag.UpdatedAt)
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
	URL  scalars.URL
	Tags []string
}) (*Source, error) {
	u := args.URL.ToDomain()

	s, err := usecase.CreateSyndicationSource(ctx, r.repositories, u, true)
	if err != nil {
		return nil, err
	}

	for _, id := range args.Tags {
		t, err := r.repositories.SyndicationTags.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if err := r.repositories.Syndication.Tag(ctx, s, t); err != nil {
			return nil, err
		}
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
	Pagination OffsetPaginationInput
	Search     SyndicationSearchInput
}) (*SourceCollection, error) {
	page := offsetPagination(10)(args.Pagination)
	sort := sorting("title", "ASC", []string{"title", "updated_at"})(args.Search.Sort)

	results, err := r.repositories.Syndication.FindAll(ctx, args.Search.Terms, args.Search.Tags, args.Search.Hidden, args.Search.Paused, page, sort)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Syndication.GetTotal(ctx, args.Search.Terms, args.Search.Tags, args.Search.Hidden, args.Search.Paused)
	if err != nil {
		return nil, err
	}

	res := SourceCollection{
		Results: resolve(r.repositories).sources(results),
		Total:   total,
		Offset:  page.Offset,
		Limit:   page.Limit,
	}

	return &res, nil
}

// Tags resolves the query
func (r *SyndicationRootResolver) Tags(ctx context.Context) (*TagCollection, error) {
	results, err := r.repositories.SyndicationTags.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.SyndicationTags.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	res := TagCollection{
		Results: resolve(r.repositories).tags(results),
		Total:   total,
	}

	return &res, nil
}

// Tag --
func (r *SyndicationRootResolver) Tag(ctx context.Context, args struct {
	SourceID string
	TagID    string
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByID(ctx, args.SourceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("source not found")
		}
		return nil, err
	}

	t, err := r.repositories.SyndicationTags.GetByID(ctx, args.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	if err := r.repositories.Syndication.Tag(ctx, s, t); err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// Untag --
func (r *SyndicationRootResolver) Untag(ctx context.Context, args struct {
	SourceID string
	TagID    string
}) (*Source, error) {
	s, err := r.repositories.Syndication.GetByID(ctx, args.SourceID)
	if err != nil {
		return nil, err
	}

	t, err := r.repositories.SyndicationTags.GetByID(ctx, args.TagID)
	if err != nil {
		return nil, err
	}

	if err := r.repositories.Syndication.Untag(ctx, s, t); err != nil {
		return nil, err
	}

	return resolve(r.repositories).source(s), nil
}

// CreateTag --
func (r *SyndicationRootResolver) CreateTag(ctx context.Context, args struct{ Label string }) (*Tag, error) {
	if args.Label == "" {
		return nil, errors.New("Tab must have a label")
	}

	t := &syndication.Tag{
		Label:     args.Label,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.repositories.SyndicationTags.Insert(ctx, t); err != nil {
		return nil, err
	}

	return resolve(r.repositories).tag(t), nil
}

// UpdateTag --
func (r *SyndicationRootResolver) UpdateTag(ctx context.Context, args struct {
	ID    string
	Label string
}) (*Tag, error) {
	if args.Label == "" {
		return nil, errors.New("Tab must have a label")
	}

	t, err := r.repositories.SyndicationTags.GetByID(ctx, args.ID)
	if err != nil {
		return nil, err
	}

	t.Label = args.Label
	t.UpdatedAt = time.Now()

	if err := r.repositories.SyndicationTags.Update(ctx, t); err != nil {
		return nil, err
	}

	return resolve(r.repositories).tag(t), nil
}

// DeleteTag --
func (r *SyndicationRootResolver) DeleteTag(ctx context.Context, args struct {
	ID string
}) (bool, error) {
	t, err := r.repositories.SyndicationTags.GetByID(ctx, args.ID)
	if err != nil {
		return false, err
	}

	if err := r.repositories.SyndicationTags.Delete(ctx, t); err != nil {
		return false, err
	}

	return true, nil
}
