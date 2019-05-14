package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/feed"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// FeedCollectionResolver resolver
type FeedCollectionResolver struct {
	Results *[]*FeedResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// FeedResolver resolves the bookmark entity
type FeedResolver struct {
	*feed.Feed
}

// ID resolves the ID field
func (r *FeedResolver) ID() graphql.ID {
	return graphql.ID(r.Feed.ID)
}

// URL resolves the URL
func (r *FeedResolver) URL() string {
	return r.Feed.URL.String()
}

// Title resolves the Title field
func (r *FeedResolver) Title() string {
	return r.Feed.Title
}

// Type resolves the Type field
func (r *FeedResolver) Type() string {
	return string(r.Feed.Type)
}

// Status resolves the Status field
func (r *FeedResolver) Status() string {
	return string(r.Feed.Status)
}

// CreatedAt resolves the CreatedAt field
func (r *FeedResolver) CreatedAt() string {
	return r.Feed.CreatedAt.Format(time.RFC3339)
}

// UpdatedAt resolves the UpdatedAt field
func (r *FeedResolver) UpdatedAt() string {
	return r.Feed.UpdatedAt.Format(time.RFC3339)
}
