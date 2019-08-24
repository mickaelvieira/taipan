package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/errors"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
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

// offsetPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		page := offsetPagination(10)(args.Offset, args.Limit)
func offsetPagination(defLimit int32) func(OffsetPaginationInput) *repository.OffsetPagination {
	if defLimit <= 0 {
		log.Fatal("the default limit must be greater than zero")
	}

	return func(i OffsetPaginationInput) *repository.OffsetPagination {
		p := &repository.OffsetPagination{}
		if i.Offset != nil {
			p.Offset = *i.Offset
		}

		if p.Offset < 0 {
			p.Offset = 0
		}

		if i.Limit != nil {
			p.Limit = *i.Limit
		}

		if p.Limit <= 0 || p.Limit > maxLimit {
			p.Limit = defLimit
		}

		return p
	}
}

// cursorPagination prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		page := cursorPagination(10)(args.First, args.Last, args.Limit)
func cursorPagination(l int32) func(i CursorPaginationInput) *repository.CursorPagination {
	if l <= 0 {
		log.Fatal("the default limit must be greater than zero")
	}

	return func(i CursorPaginationInput) *repository.CursorPagination {
		p := &repository.CursorPagination{}

		if i.From != nil {
			p.From = *i.From
		}

		if i.To != nil {
			p.To = *i.To
		}

		if i.Limit != nil {
			p.Limit = *i.Limit
		}

		if p.Limit <= 0 || p.Limit > maxLimit {
			p.Limit = l
		}

		return p
	}
}

func sorting(by string, dir string, allowed []string) func(i *SortingInput) *repository.Sorting {
	okBy := func(b *string) bool {
		if b == nil {
			return false
		}
		v := *b
		for _, a := range allowed {
			if a == v {
				return true
			}
		}
		return false
	}

	okDir := func(d *string) bool {
		if d == nil {
			return false
		}
		v := *d
		if v == "ASC" || v == "DESC" {
			return true
		}
		return false
	}

	return func(i *SortingInput) *repository.Sorting {
		s := &repository.Sorting{
			By:  by,
			Dir: dir,
		}
		if i != nil {
			if okBy(i.By) {
				s.By = *i.By
			}
			if okDir(i.Dir) {
				s.Dir = *i.Dir
			}
		}
		return s
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
