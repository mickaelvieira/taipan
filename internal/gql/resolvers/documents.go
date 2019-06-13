package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/repository"
	"time"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

// DocumentCollectionResolver resolver
type DocumentCollectionResolver struct {
	Results *[]*DocumentResolver
	Total   int32
	First   string
	Last    string
	Limit   int32
}

// DocumentResolver resolves the bookmark entity
type DocumentResolver struct {
	*document.Document
	logsLoader   *dataloader.Loader
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *DocumentResolver) ID() graphql.ID {
	return graphql.ID(r.Document.ID)
}

// URL resolves the URL
func (r *DocumentResolver) URL() string {
	return r.Document.URL.String()
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
func (r *DocumentResolver) CreatedAt() string {
	return r.Document.CreatedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *DocumentResolver) UpdatedAt() string {
	return r.Document.UpdatedAt.Format(time.RFC3339)
}

// LogEntries returns the document's parser log
func (r *DocumentResolver) LogEntries(ctx context.Context) (*[]*HTTPClientLogResolver, error) {
	data, err := r.logsLoader.Load(ctx, dataloader.StringKey(r.Document.URL.String()))()
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

// News resolves the query
func (r *Resolvers) News(ctx context.Context, args struct {
	From  *string
	To    *string
	Limit *int32
}) (*DocumentCollectionResolver, error) {
	fromArgs := GetCursorBasedPagination(10)
	from, to, limit := fromArgs(args.From, args.To, args.Limit)
	user := auth.FromContext(ctx)

	results, err := r.repositories.Documents.FindNew(ctx, user, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := GetDocumentsBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Documents.GetTotalNew(ctx, user)
	if err != nil {
		return nil, err
	}

	var logsLoader = loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
	var documents []*DocumentResolver
	for _, result := range results {
		documents = append(documents, &DocumentResolver{
			Document:     result,
			logsLoader:   logsLoader,
			repositories: r.repositories,
		})
	}

	reso := DocumentCollectionResolver{
		Results: &documents,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &reso, nil
}

// Documents resolves the query
func (r *Resolvers) Documents(ctx context.Context, args struct {
	From  *string
	To    *string
	Limit *int32
}) (*DocumentCollectionResolver, error) {
	fromArgs := GetCursorBasedPagination(10)
	from, to, limit := fromArgs(args.From, args.To, args.Limit)

	results, err := r.repositories.Documents.GetDocuments(ctx, from, to, limit)
	if err != nil {
		return nil, err
	}

	first, last := GetDocumentsBoundaryIDs(results)

	var total int32
	total, err = r.repositories.Documents.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var logsLoader = loaders.GetHTTPClientLogEntriesLoader(r.repositories.Botlogs)
	var documents []*DocumentResolver
	for _, result := range results {
		documents = append(documents, &DocumentResolver{
			Document:     result,
			logsLoader:   logsLoader,
			repositories: r.repositories,
		})
	}

	reso := DocumentCollectionResolver{
		Results: &documents,
		Total:   total,
		First:   first,
		Last:    last,
		Limit:   limit,
	}

	return &reso, nil
}
