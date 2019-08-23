package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/errors"
	"github/mickaelvieira/taipan/internal/logger"
	"log"
)

func handleError(err error) error {
	if err, ok := err.(errors.DomainError); ok {
		if err.HasReason() {
			logger.Debug(err.Reason())
		}
		return err.Domain()
	}
	logger.Error(err)
	return err
}

const maxLimit = 100

// getOffsetBasedPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		fromArgs := getOffsetBasedPagination(10)
// 		offset, limit := fromArgs(args.Offset, args.Limit)
func getOffsetBasedPagination(defLimit int32) func(OffsetPaginationInput) (int32, int32) {
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

// getCursorBasedPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		fromArgs := getCursorBasedPagination(10)
// 		from, to, limit := fromArgs(args.First, args.Last, args.Limit)
func getCursorBasedPagination(defLimit int32) func(i CursorPaginationInput) (string, string, int32) {
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

// getBookmarksBoundaryIDs returns the first and last ID in the results set
func getBookmarksBoundaryIDs(results []*bookmark.Bookmark) (first string, last string) {
	if len(results) > 0 {
		first = results[0].ID
		last = results[len(results)-1].ID
	}
	return
}

// getDocumentsBoundaryIDs returns the first and last ID in the results set
func getDocumentsBoundaryIDs(results []*document.Document) (first string, last string) {
	if len(results) > 0 {
		first = results[0].ID
		last = results[len(results)-1].ID
	}
	return
}
