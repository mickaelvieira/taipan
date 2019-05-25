package document

import (
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/uri"
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

// Document struct represents a web document
type Document struct {
	ID          string
	Checksum    checksum.Checksum
	URL         *uri.URI
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *image.Image
	Feeds       []*feed.Feed
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
}

func (d *Document) String() string {
	return d.ID
}

// Raw returns the key raws data
func (d *Document) Raw() interface{} {
	return d
}

// New creates a new document
func New(url *uri.URI, lang string, charset string, title string, desc string, image *image.Image, feeds []*feed.Feed) *Document {
	return &Document{
		URL:         url,
		Lang:        lang,
		Charset:     charset,
		Title:       title,
		Description: desc,
		Image:       image,
		Feeds:       feeds,
		Status:      FETCHED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
