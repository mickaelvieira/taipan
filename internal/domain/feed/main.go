package feed

import "time"

// Type represents whether the feed is an atom or rss feed
type Type string

// Type values
const (
	ATOM Type = "application/atom+xml"
	RSS  Type = "application/rss+xml"
)

// Status represents the status of the feed during the fetching process
type Status string

// Status values
const (
	NEW      Status = "new"
	PENDING  Status = "pending"
	FETCHING Status = "fetching"
)

// Feed represents what is the feed within the application
type Feed struct {
	ID        string
	URL       string
	Hash      string
	Type      Type
	Title     string
	status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserFeed represents a feed from a user's prespective
type UserFeed struct {
	Feed
	AddedAt time.Time
}

// FetchingHistory represents an entry in the history logs
// @TODO I need to find a better name
type FetchingHistory struct {
	ID               string
	ReqURI           string
	ReqMethod        string
	ReqHeaders       string
	RespStatusCode   int
	RespReasonPhrase string
	RespHeaders      string
	CreatedAt        time.Time
}
