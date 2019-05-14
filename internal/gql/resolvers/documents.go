package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/repository"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// DocumentCollectionResolver resolver
type DocumentCollectionResolver struct {
	Results *[]*DocumentResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// DocumentResolver resolves the bookmark entity
type DocumentResolver struct {
	*document.Document
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
	if r.Document.Image == nil {
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

// Feeds returns the document's feeds
func (r *DocumentResolver) Feeds(ctx context.Context) (*[]*FeedResolver, error) {
	results, err := r.repositories.Feeds.GetDocumentFeeds(ctx, r.Document)
	if err != nil {
		return nil, err
	}

	var feeds []*FeedResolver
	for _, result := range results {
		res := FeedResolver{Feed: result}
		feeds = append(feeds, &res)
	}

	return &feeds, nil
}

// GetLatestDocuments resolves the query
func (r *Resolvers) GetLatestDocuments(ctx context.Context, args struct {
	Offset *int32
	Limit  *int32
}) (*DocumentCollectionResolver, error) {
	fromArgs := GetBoundariesFromArgs(10)
	offset, limit := fromArgs(args.Offset, args.Limit)
	user := auth.FromContext(ctx)

	results, err := r.Repositories.Documents.FindNew(ctx, user, offset, limit)
	if err != nil {
		return nil, err
	}

	total, err := r.Repositories.Documents.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	var documents []*DocumentResolver
	for _, result := range results {
		res := DocumentResolver{
			Document:     result,
			repositories: r.Repositories,
		}
		documents = append(documents, &res)
	}

	reso := DocumentCollectionResolver{
		Results: &documents,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &reso, nil
}
