package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/graphql/scalars"

	gql "github.com/graph-gophers/graphql-go"
)

// HTTPClientLogResolver resolves the bookmark's image entity
type HTTPClientLogResolver struct {
	*http.Result
}

// ID resolves the ID
func (r *HTTPClientLogResolver) ID() gql.ID {
	return gql.ID(r.Result.ID)
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
func (r *HTTPClientLogResolver) RequestURI() scalars.URL {
	return scalars.NewURL(r.Result.ReqURI)
}

// RequestMethod resolves the RequestMethod field
func (r *HTTPClientLogResolver) RequestMethod() string {
	return r.Result.ReqMethod
}

// HasFailed resolves the HasFailed field
func (r *HTTPClientLogResolver) HasFailed() bool {
	return r.Result.RequestHasFailed()
}

// FailureReason resolves the FailureReason field
func (r *HTTPClientLogResolver) FailureReason() string {
	if r.Result.RequestHasFailed() {
		return r.Result.GetFailureReason()
	}
	return ""
}

// FinalURI resolves the FinalURI field
// func (r *HTTPClientLogResolver) FinalURI() scalars.URL {
// 	return scalars.NewURL(r.Result.FinalURI)
// }

// CreatedAt resolves the CreatedAt field
func (r *HTTPClientLogResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.Result.CreatedAt)
}
