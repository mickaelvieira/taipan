package http

import (
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io"
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

// // ResultPipe build a pipe
// type ResultPipe []*Result

// // Sort it
// func (r *ResultPipe) Sort(T func(ResultPipe) sort.Interface) *ResultPipe {
// 	p := *r
// 	sort.Sort(T(p))
// 	*r = p
// 	return r
// }

// // Map it
// func (r *ResultPipe) Map(T func(ResultPipe) ResultPipe) {
// 	p := *r
// 	T(p)
// 	*r = p
// }

// // test
// func bla() {
// 	var r = make(ResultPipe, 0)
// 	r.
// 		Sort(func(r ResultPipe) sort.Interface {
// 			return ByCreatedAt(r)
// 		}).
// 		Map(func(r ResultPipe) ResultPipe {
// 			return r
// 		})
// }
