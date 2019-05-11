package bookmark

import (
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"time"
)

// ReadStatus are the values defineing whether or not a bookmark has been read
type ReadStatus bool

// ReadStatus values
const (
	UNREAD ReadStatus = false
	READ   ReadStatus = true
)

// Bookmark struct represents what is a bookmark from a user's perspective
type Bookmark struct {
	ID          string
	URL         *uri.URI
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *image.Image
	AddedAt     time.Time
	UpdatedAt   time.Time
	IsRead      bool
	IsLinked    bool
}
