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

// DocumentRootResolver documents' root resolver
type DocumentRootResolver struct {
	repositories *repository.Repositories
}

// DocumentCollection resolver
type DocumentCollection struct {
	Results []*Document
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// DocumentSearchResults resolver
type DocumentSearchResults struct {
	Results []*Document
	Total   int32
	Offset  int32
	Limit   int32
}

// Document resolves the bookmark entity
type Document struct {
	document     *document.Document
	repositories *repository.Repositories
	sourceLoader *dataloader.Loader
	logLoader    *dataloader.Loader
}

// ID resolves the ID field
func (r *Document) ID() gql.ID {
	return gql.ID(r.document.ID)
}

// URL resolves the URL
func (r *Document) URL() scalars.URL {
	return scalars.NewURL(r.document.URL)
}

// Image resolves the Image field
func (r *Document) Image() *BookmarkImage {
	if !r.document.HasImage() {
		return nil
	}
	return &BookmarkImage{Image: r.document.Image}
}

// Lang resolves the Lang field
func (r *Document) Lang() string {
	return r.document.Lang
}

// Charset resolves the Charset field
func (r *Document) Charset() string {
	return r.document.Charset
}

// Title resolves the Title field
func (r *Document) Title() string {
	return r.document.Title
}

// Description resolves the Description field
func (r *Document) Description() string {
	return r.document.Description
}

// CreatedAt resolves the CreatedAt field
func (r *Document) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.document.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *Document) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.document.UpdatedAt)
}

// Source resolves the Source field
func (r *Document) Source(ctx context.Context) (*Source, error) {
	data, err := r.sourceLoader.Load(ctx, dataloader.StringKey(r.document.SourceID))()
	if err != nil {
		return nil, err
	}

	result, ok := data.(*syndication.Source)
	if !ok {
		return nil, errors.New("Loader returns incorrect type")
	}

	return resolve(r.repositories).source(result), nil
}

// Syndication resolves the Syndication field
func (r *Document) Syndication() *[]*Source {
	res := resolve(r.repositories).sources(r.document.Feeds)
	return &res
}

// LogEntries returns the document's parser log
func (r *Document) LogEntries(ctx context.Context) (*[]*Log, error) {
	data, err := r.logLoader.Load(ctx, dataloader.StringKey(r.document.URL.String()))()
	if err != nil {
		return nil, err
	}

	results, ok := data.([]*http.Result)
	if !ok {
		return nil, fmt.Errorf("Invalid data")
	}

	res := resolve(r.repositories).logs(results)

	return &res, nil
}

// DocumentEvent is a document event
type DocumentEvent struct {
	event        *publisher.Event
	repositories *repository.Repositories
}

// Item returns the event's message
func (r *DocumentEvent) Item() *Document {
	d, ok := r.event.Payload.(*document.Document)
	if !ok {
		log.Fatalln("Cannot resolve item, payload is not a document")
	}

	return resolve(r.repositories).document(d)
}

// Emitter returns the event's emitter ID
func (r *DocumentEvent) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *DocumentEvent) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *DocumentEvent) Action() string {
	return string(r.event.Action)
}

// DocumentChanged --
func (r *RootResolver) DocumentChanged(ctx context.Context) <-chan *DocumentEvent {
	// @TODO better handle authentication
	c := make(chan *DocumentEvent)
	s := &documentSubscriber{events: c}
	r.publisher.Subscribe(publisher.TopicDocument, s, ctx.Done())
	return c
}

// Create creates a new document
func (r *DocumentRootResolver) Create(ctx context.Context, args struct {
	URL scalars.URL
}) (*Document, error) {
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
func (r *DocumentRootResolver) Documents(ctx context.Context, args struct {
	Pagination cursorPaginationInput
}) (*DocumentCollection, error) {
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

	res := DocumentCollection{
		Results: resolve(r.repositories).documents(results),
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &res, nil
}

// Search --
func (r *DocumentRootResolver) Search(ctx context.Context, args struct {
	Pagination offsetPaginationInput
	Search     bookmarkSearchInput
}) (*DocumentSearchResults, error) {
	user := auth.FromContext(ctx)
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)
	terms := args.Search.Terms

	var documents []*Document
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

	res := DocumentSearchResults{
		Results: documents,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
