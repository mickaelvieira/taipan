package bookmark

import (
	"github/mickaelvieira/taipan/internal/domain/uuid"
	"time"
)

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
type ReadStatus bool

// see https://gist.github.com/husobee/cac9cddbaacc1d3a7ae1
// ReadStatus values
const (
	UNREAD ReadStatus = false
	READ   ReadStatus = true
)

// Bookmark struct represents what is a bookmark within the application
type Bookmark struct {
	ID          string
	URL         string
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// UserBookmark struct represents what is a bookmark from a user's perspective
type UserBookmark struct {
	ID          string
	URL         string
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       string
	AddedAt     time.Time
	UpdatedAt   time.Time
	IsRead      bool
	IsLinked    bool
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

// New creates a new Bookmark with a UUID
func New(url string, lang string, charset string, title string, desc string, image string) *Bookmark {
	return &Bookmark{
		ID:          uuid.New(),
		URL:         url,
		Lang:        lang,
		Charset:     charset,
		Title:       title,
		Description: desc,
		Image:       image,
		Status:      FETCHED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
