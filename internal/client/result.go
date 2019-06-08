package client

import (
	"bytes"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"time"
)

// Result represents an entry in the history logs
type Result struct {
	ID               string
	WasRedirected    bool
	Checksum         checksum.Checksum
	ContentType      string
	ReqURI           string
	FinalURI         string
	ReqMethod        string
	ReqHeaders       string
	RespStatusCode   int
	RespReasonPhrase string
	RespHeaders      string
	CreatedAt        time.Time
	Content          *bytes.Reader
}

// IsContentDifferent have we fetched a new document
func (r *Result) IsContentDifferent(prev *Result) bool {
	return prev == nil || prev.Checksum.String() != r.Checksum.String()
}
