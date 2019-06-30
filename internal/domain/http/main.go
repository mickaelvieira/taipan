package http

import (
	"bytes"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"time"
)

// Result represents an entry in the history logs
type Result struct {
	ID               string
	WasRedirected    bool
	Checksum         checksum.Checksum
	ContentType      string
	ReqURI           *url.URL
	FinalURI         *url.URL
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

// ByCreatedAt sorts Results by ascending creation date
type ByCreatedAt []*Result

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }
