package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/graphql/loaders"
	"github/mickaelvieira/taipan/internal/graphql/scalars"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graph-gophers/graphql-go"
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

// DocumentResolver resolves the bookmark entity
type DocumentResolver struct {
	*document.Document
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *DocumentResolver) ID() gql.ID {
	return gql.ID(r.Document.ID)
}

// URL resolves the URL
func (r *DocumentResolver) URL() scalars.URL {
	return scalars.URL{URL: r.Document.URL}
}

// Image resolves the Image field
func (r *DocumentResolver) Image() *BookmarkImageResolver {
	if r.Document.Image == nil || r.Document.Image.Name == "" {
		return nil
	}

	return &BookmarkImageResolver{
		Image: r.Document.Image,
	}
}

// Lang resolves the Lang field
func (r *DocumentResolver) Lang() string {
	return r.Document.Lang
}

// Charset resolves the Charset field
func (r *DocumentResolver) Charset() string {
	return r.Document.Charset
}

// Title resolves the Title field
func (r *DocumentResolver) Title() string {
	return r.Document.Title
}

// Description resolves the Description field
func (r *DocumentResolver) Description() string {
	return r.Document.Description
}

// CreatedAt resolves the CreatedAt field
func (r *DocumentResolver) CreatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Document.CreatedAt}
}

// UpdatedAt resolves the UpdatedAt field
func (r *DocumentResolver) UpdatedAt() scalars.DateTime {
	return scalars.DateTime{Time: r.Document.UpdatedAt}
}

// LogEntries returns the document's parser log
func (r *DocumentResolver) LogEntries(ctx context.Context) (*[]*HTTPClientLogResolver, error) {
	data, err := r.getLogsLoader().Load(ctx, dataloader.StringKey(r.Document.URL.String()))()
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

func (r *DocumentResolver) getLogsLoader() *dataloader.Loader {
	return loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
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

	first, last := getDocumentsBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Documents.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var documents = make([]*DocumentResolver, 0)
	for _, result := range results {
		documents = append(documents, &DocumentResolver{
			Document:     result,
			repositories: r.repositories,
		})
	}

	reso := DocumentCollectionResolver{
		Results: documents,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &reso, nil
}
