package http

import (
	"github.com/mickaelvieira/taipan/internal/domain/checksum"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"io"
	nethttp "net/http"
	"time"
)

// Result represents an entry in the history logs
type Result struct {
	ID               string
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
	Content          io.Reader
	Failed           bool
	FailureReason    string
}

// RequestWasRedirected is the final URL different from the requested URL
func (r *Result) RequestWasRedirected() bool {
	return r.ReqURI.String() != r.FinalURI.String()
}

// RequestWasSuccessful determines whether the request was successful
func (r *Result) RequestWasSuccessful() bool {
	return r.RespStatusCode >= nethttp.StatusOK && r.RespStatusCode < nethttp.StatusMultipleChoices || r.RespStatusCode == nethttp.StatusNotModified
}

// GetFailureReason returns the failure reason
func (r Result) GetFailureReason() string {
	if r.Failed {
		return r.FailureReason
	}
	return r.RespReasonPhrase
}

// IsContentDifferent have we fetched a new document
func (r *Result) IsContentDifferent(prev *Result) bool {
	return prev == nil || prev.Checksum.String() != r.Checksum.String()
}

// ByCreatedAt sorts Results by ascending creation date
type ByCreatedAt []*Result

func (a ByCreatedAt) Len() int {
	return len(a)
}

func (a ByCreatedAt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCreatedAt) Less(i, j int) bool {
	return a[i].CreatedAt.Before(a[j].CreatedAt)
}
