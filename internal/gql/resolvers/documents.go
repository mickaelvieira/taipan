package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/document"
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
