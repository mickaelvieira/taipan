package resolvers

import (
	"github/mickaelvieira/taipan/internal/client"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// HTTPClientLogResolver resolves the bookmark's image entity
type HTTPClientLogResolver struct {
	*client.Result
}

// ID resolves the ID
func (r *HTTPClientLogResolver) ID() graphql.ID {
	return graphql.ID(r.Result.ID)
}

// Checksum resolves the Checksum
func (r *HTTPClientLogResolver) Checksum() string {
	return r.Result.Checksum.String()
}

// ContentType resolves the ContentType field
func (r *HTTPClientLogResolver) ContentType() string {
	return r.Result.ContentType
}

// StatusCode resolves the StatusCode field
func (r *HTTPClientLogResolver) StatusCode() int32 {
	return int32(r.Result.RespStatusCode)
}

// RequestURI resolves the RequestURI field
func (r *HTTPClientLogResolver) RequestURI() string {
	return r.Result.ReqURI
}

// CreatedAt resolves the CreatedAt field
func (r *HTTPClientLogResolver) CreatedAt() string {
	return r.Result.CreatedAt.Format(time.RFC3339)
}
