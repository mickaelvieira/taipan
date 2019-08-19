package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/graphql/scalars"

	gql "github.com/graph-gophers/graphql-go"
)

// BotResolver documents' root resolver
type BotResolver struct {
	repositories *repository.Repositories
}

// LogResolver resolves the bookmark's image entity
type LogResolver struct {
	l *http.Result
	r *repository.Repositories
}

// LogCollectionResolver resolver
type LogCollectionResolver struct {
	Results []*LogResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// ID resolves the ID
func (r *LogResolver) ID() gql.ID {
	return gql.ID(r.l.ID)
}

// Checksum resolves the Checksum
func (r *LogResolver) Checksum() string {
	return r.l.Checksum.String()
}

// ContentType resolves the ContentType field
func (r *LogResolver) ContentType() string {
	return r.l.ContentType
}

// StatusCode resolves the StatusCode field
func (r *LogResolver) StatusCode() int32 {
	return int32(r.l.RespStatusCode)
}

// RequestURI resolves the RequestURI field
func (r *LogResolver) RequestURI() scalars.URL {
	return scalars.NewURL(r.l.ReqURI)
}

// RequestMethod resolves the RequestMethod field
func (r *LogResolver) RequestMethod() string {
	return r.l.ReqMethod
}

// HasFailed resolves the HasFailed field
func (r *LogResolver) HasFailed() bool {
	return r.l.RequestHasFailed()
}

// FailureReason resolves the FailureReason field
func (r *LogResolver) FailureReason() string {
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
func (r *LogResolver) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.l.CreatedAt)
}

// Logs --
func (r *BotResolver) Logs(ctx context.Context, args struct {
	URL        scalars.URL
	Pagination offsetPaginationInput
}) (*LogCollectionResolver, error) {
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

	res := LogCollectionResolver{
		Results: resolve(r.repositories).logs(results),
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}

	return &res, nil
}
