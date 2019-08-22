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
	log          *http.Result
	repositories *repository.Repositories
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
	return gql.ID(r.log.ID)
}

// Checksum resolves the Checksum
func (r *Log) Checksum() string {
	return r.log.Checksum.String()
}

// ContentType resolves the ContentType field
func (r *Log) ContentType() string {
	return r.log.ContentType
}

// StatusCode resolves the StatusCode field
func (r *Log) StatusCode() int32 {
	return int32(r.log.RespStatusCode)
}

// RequestURI resolves the RequestURI field
func (r *Log) RequestURI() scalars.URL {
	return scalars.NewURL(r.log.ReqURI)
}

// RequestMethod resolves the RequestMethod field
func (r *Log) RequestMethod() string {
	return r.log.ReqMethod
}

// HasFailed resolves the HasFailed field
func (r *Log) HasFailed() bool {
	return !r.log.RequestWasSuccessful()
}

// FailureReason resolves the FailureReason field
func (r *Log) FailureReason() string {
	if !r.log.RequestWasSuccessful() {
		return r.log.GetFailureReason()
	}
	return ""
}

// FinalURI resolves the FinalURI field
// func (r *LogResolver) FinalURI() scalars.URL {
// 	return scalars.NewURL(r.l.FinalURI)
// }

// CreatedAt resolves the CreatedAt field
func (r *Log) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.log.CreatedAt)
}

// Logs --
func (r *LogRootResolver) Logs(ctx context.Context, args struct {
	URL        scalars.URL
	Pagination OffsetPaginationInput
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
