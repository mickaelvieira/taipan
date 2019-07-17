package bookmark

import (
	"github/mickaelvieira/taipan/internal/domain/document"
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
	Image       *document.Image
	AddedAt     time.Time
	UpdatedAt   time.Time
	IsFavorite  bool
	IsLinked    bool
}

// HasImage determine whether the bookmark has an image associated to it
func (b *Bookmark) HasImage() bool {
	return b.Image != nil && b.Image.URL != nil && b.Image.Name != ""
}
