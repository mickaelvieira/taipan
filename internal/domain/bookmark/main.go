package bookmark

import "time"

// Status defines the status of a bookmark
type Status string

// Status values
const (
	FETCHED  Status = "fetched"
	PENDING  Status = "pending"
	FETCHING Status = "fetching"
	FAILED   Status = "failed"
)

// ReadStatus are the values defineing whether or not a bookmark has been read
type ReadStatus int

// ReadStatus values
const (
	UNREAD ReadStatus = iota
	READ
)

// Bookmark struct represents what is a bookmark within the application
type Bookmark struct {
	ID          string
	URL         string
	Hash        string
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// UserBookmark struct represents what is a bookmark from a user's perspective
type UserBookmark struct {
	Bookmark
	AddedAt    time.Time
	AccessedAt time.Time
	IsRead     ReadStatus
	IsLinked   bool
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
