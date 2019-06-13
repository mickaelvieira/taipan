package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"log"
)

const maxLimit = 100

// CursorPaginationInput graphQL input to handle cursor based pagination
type CursorPaginationInput struct {
	From  *string
	To    *string
	Limit *int32
}

// OffsetPaginationInput graphQL input to handle offset based pagination
type OffsetPaginationInput struct {
	Offset *int32
	Limit  *int32
}

// GetOffsetBasedPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		fromArgs := GetOffsetBasedPagination(10)
// 		offset, limit := fromArgs(args.Offset, args.Limit)
func GetOffsetBasedPagination(defLimit int32) func(OffsetPaginationInput) (int32, int32) {
	if defLimit <= 0 {
		log.Fatal("the default limit must be greater than zero")
	}

	return func(i OffsetPaginationInput) (offset int32, limit int32) {
		if i.Offset != nil {
			offset = *i.Offset
		}

		if offset < 0 {
			offset = 0
		}

		if i.Limit != nil {
			limit = *i.Limit
		}

		if limit <= 0 || limit > maxLimit {
			limit = defLimit
		}

		return offset, limit
	}
}

// GetCursorBasedPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		fromArgs := GetCursorBasedPagination(10)
// 		from, to, limit := fromArgs(args.First, args.Last, args.Limit)
func GetCursorBasedPagination(defLimit int32) func(i CursorPaginationInput) (string, string, int32) {
	if defLimit <= 0 {
		log.Fatal("the default limit must be greater than zero")
	}

	return func(i CursorPaginationInput) (from string, to string, limit int32) {
		if i.From != nil {
			from = *i.From
		}

		if i.To != nil {
			to = *i.To
		}

		if i.Limit != nil {
			limit = *i.Limit
		}

		if limit <= 0 || limit > maxLimit {
			limit = defLimit
		}

		return from, to, limit
	}
}

// GetBookmarksBoundaryIDs returns the first and last ID in the results set
func GetBookmarksBoundaryIDs(results []*bookmark.Bookmark) (first string, last string) {
	if len(results) > 0 {
		first = results[0].ID
		last = results[len(results)-1].ID
	}
	return
}

// GetDocumentsBoundaryIDs returns the first and last ID in the results set
func GetDocumentsBoundaryIDs(results []*document.Document) (first string, last string) {
	if len(results) > 0 {
		first = results[0].ID
		last = results[len(results)-1].ID
	}
	return
}
