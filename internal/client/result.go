package client

import (
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"time"
)

// Result represents an entry in the history logs
type Result struct {
	ID               string
	Checksum         checksum.Checksum
	ContentType      string
	ReqURI           string
	ReqMethod        string
	ReqHeaders       string
	RespStatusCode   int
	RespReasonPhrase string
	RespHeaders      string
	CreatedAt        time.Time
}

// IsContentDifferent have we fetched a new document
func (r *Result) IsContentDifferent(prev *Result) bool {
	return prev != nil && prev.Checksum.String() == r.Checksum.String()
}
