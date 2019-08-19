package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"
	"log"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// DocumentsResolver documents' root resolver
type DocumentsResolver struct {
	repositories *repository.Repositories
}

// DocumentCollectionResolver resolver
type DocumentCollectionResolver struct {
	Results []*DocumentResolver
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// DocumentSearchResultsResolver resolver
type DocumentSearchResultsResolver struct {
	Results []*DocumentResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// DocumentResolver resolves the bookmark entity
type DocumentResolver struct {
	d  *document.Document
	r  *repository.Repositories
	sl *dataloader.Loader
	ll *dataloader.Loader
}

// ID resolves the ID field
func (r *DocumentResolver) ID() gql.ID {
	return gql.ID(r.d.ID)
}

// URL resolves the URL
func (r *DocumentResolver) URL() scalars.URL {
	return scalars.NewURL(r.d.URL)
}

// Image resolves the Image field
func (r *DocumentResolver) Image() *BookmarkImageResolver {
	if !r.d.HasImage() {
		return nil
	}
	return &BookmarkImageResolver{Image: r.d.Image}
}

// Lang resolves the Lang field
func (r *DocumentResolver) Lang() string {
	return r.d.Lang
}

// Charset resolves the Charset field
func (r *DocumentResolver) Charset() string {
	return r.d.Charset
}

// Title resolves the Title field
func (r *DocumentResolver) Title() string {
	return r.d.Title
}

// Description resolves the Description field
func (r *DocumentResolver) Description() string {
	return r.d.Description
}

// CreatedAt resolves the CreatedAt field
func (r *DocumentResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.d.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *DocumentResolver) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.d.UpdatedAt)
}

// Source resolves the Source field
func (r *DocumentResolver) Source(ctx context.Context) (*SourceResolver, error) {
	data, err := r.sl.Load(ctx, dataloader.StringKey(r.d.SourceID))()
	if err != nil {
		return nil, err
	}

	result, ok := data.(*syndication.Source)
	if !ok {
		return nil, errors.New("Loader returns incorrect type")
	}

	return resolve(r.r).source(result), nil
}

// Syndication resolves the Syndication field
func (r *DocumentResolver) Syndication() *[]*SourceResolver {
	res := resolve(r.r).sources(r.d.Feeds)
	return &res
}

// LogEntries returns the document's parser log
func (r *DocumentResolver) LogEntries(ctx context.Context) (*[]*LogResolver, error) {
	data, err := r.ll.Load(ctx, dataloader.StringKey(r.d.URL.String()))()
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

// DocumentEventResolver is a document event
type DocumentEventResolver struct {
	event        *publisher.Event
	repositories *repository.Repositories
}

// Item returns the event's message
func (r *DocumentEventResolver) Item() *DocumentResolver {
	d, ok := r.event.Payload.(*document.Document)
	if !ok {
		log.Fatalln("Cannot resolve item, payload is not a document")
	}

	return resolve(r.repositories).document(d)
}

// Emitter returns the event's emitter ID
func (r *DocumentEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *DocumentEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *DocumentEventResolver) Action() string {
	return string(r.event.Action)
}

// DocumentChanged --
func (r *RootResolver) DocumentChanged(ctx context.Context) <-chan *DocumentEventResolver {
	// @TODO better handle authentication
	c := make(chan *DocumentEventResolver)
	s := &documentSubscriber{events: c}
	r.publisher.Subscribe(publisher.TopicDocument, s, ctx.Done())
	return c
}

// Create creates a new document
func (r *DocumentsResolver) Create(ctx context.Context, args struct {
	URL scalars.URL
}) (*DocumentResolver, error) {
	user := auth.FromContext(ctx)
	d, err := usecase.Document(ctx, r.repositories, args.URL.ToDomain(), true)
	if err != nil {
		return nil, err
	}

	// excludes user's subscriptions
	var subscriptions []*syndication.Source
	for _, s := range d.Feeds {
		exists, err := r.repositories.Subscriptions.ExistWithURL(ctx, user, s.URL)
		if err != nil {
			return nil, err
		}
		if !exists {
			subscriptions = append(subscriptions, s)
		}
	}

	d.Feeds = subscriptions

	return resolve(r.repositories).document(d), nil
}

// Documents resolves the query
func (r *DocumentsResolver) Documents(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollectionResolver, error) {
	fromArgs := getCursorBasedPagination(10)
	from, to, limit := fromArgs(args.Pagination)

	results, err := r.repositories.Documents.GetDocuments(ctx, from, to, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.repositories.Documents.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	first, last := getDocumentsBoundaryIDs(results)

	res := DocumentCollectionResolver{
		Results: resolve(r.repositories).documents(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// Search --
func (r *DocumentsResolver) Search(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     bookmarkSearchInput
}) (*DocumentSearchResultsResolver, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)
	terms := args.Search.Terms

	var documents []*DocumentResolver
	var total int32

	if len(terms) > 0 {
		results, err := r.repositories.Documents.FindAll(ctx, user, terms, offset, limit)
		if err != nil {
			return nil, err
		}

		total, err = r.repositories.Documents.CountAll(ctx, user, terms)
		if err != nil {
			return nil, err
		}

		documents = resolve(r.repositories).documents(results)
	}

	res := DocumentSearchResultsResolver{
		Results: documents,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
