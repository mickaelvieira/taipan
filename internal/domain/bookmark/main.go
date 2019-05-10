package bookmark

import (
	"github/mickaelvieira/taipan/internal/domain/types"
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

// https://husobee.github.io/golang/database/2015/06/12/scanner-valuer.html

// Bookmark struct represents what is a bookmark within the application
type Bookmark struct {
	ID          string
	Checksum    types.Checksum
	URL         *types.URI
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *Image
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Image represents a bookmark's image
type Image struct {
	Name   string
	URL    *types.URI
	Width  int32
	Height int32
	Format string
}

func (i *Image) String() string {
	if i.URL == nil {
		return ""
	}
	return i.URL.String()
}

// UserBookmark struct represents what is a bookmark from a user's perspective
type UserBookmark struct {
	ID          string
	URL         *types.URI
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *Image
	AddedAt     time.Time
	UpdatedAt   time.Time
	IsRead      bool
	IsLinked    bool
}

// New creates a new Bookmark with a UUID
func New(url *types.URI, lang string, charset string, title string, desc string, image *Image) *Bookmark {
	return &Bookmark{
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
