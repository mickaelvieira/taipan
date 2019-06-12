package bookmark

import (
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/url"
	"time"
)

// Bookmark struct represents what is a bookmark from a user's perspective
type Bookmark struct {
	ID          string
	URL         *url.URL
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
