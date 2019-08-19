package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

	gql "github.com/graph-gophers/graphql-go"
)

// LogRootResolver documents' root resolver
type LogRootResolver struct {
	repositories *repository.Repositories
}

// Log resolves the bookmark's image entity
type Log struct {
	l *http.Result
	r *repository.Repositories
}

// LogCollection resolver
type LogCollection struct {
	Results []*Log
	Total   int32
	Offset  int32
	Limit   int32
}

// ID resolves the ID
func (r *Log) ID() gql.ID {
	return gql.ID(r.l.ID)
}

// Checksum resolves the Checksum
func (r *Log) Checksum() string {
	return r.l.Checksum.String()
}

// ContentType resolves the ContentType field
func (r *Log) ContentType() string {
	return r.l.ContentType
}

// StatusCode resolves the StatusCode field
func (r *Log) StatusCode() int32 {
	return int32(r.l.RespStatusCode)
}

// RequestURI resolves the RequestURI field
func (r *Log) RequestURI() scalars.URL {
	return scalars.NewURL(r.l.ReqURI)
}

// RequestMethod resolves the RequestMethod field
func (r *Log) RequestMethod() string {
	return r.l.ReqMethod
}

// HasFailed resolves the HasFailed field
func (r *Log) HasFailed() bool {
	return r.l.RequestHasFailed()
}

// FailureReason resolves the FailureReason field
func (r *Log) FailureReason() string {
	if r.l.RequestHasFailed() {
		return r.l.GetFailureReason()
	}
	return ""
}

// FinalURI resolves the FinalURI field
// func (r *LogResolver) FinalURI() scalars.URL {
// 	return scalars.NewURL(r.l.FinalURI)
// }

// CreatedAt resolves the CreatedAt field
func (r *Log) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.l.CreatedAt)
}

// Logs --
func (r *LogRootResolver) Logs(ctx context.Context, args struct {
	URL        scalars.URL
	Pagination offsetPaginationInput
}) (*LogCollection, error) {
	u := args.URL.ToDomain()
	fromArgs := getOffsetBasedPagination(10)
	offset, limit := fromArgs(args.Pagination)

	results, err := r.repositories.Botlogs.FindAll(ctx, u, offset, limit)
	if err != nil {
		return nil, err
	}

	var total int32
	total, err = r.repositories.Botlogs.CountAll(ctx, u)
	if err != nil {
		return nil, err
	}

	res := LogCollection{
		Results: resolve(r.repositories).logs(results),
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
