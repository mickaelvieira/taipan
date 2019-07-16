package subscription

import (
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"time"
)

// Subscription represents a user'subscription to a syndication source
type Subscription struct {
	ID         string
	URL        *url.URL
	Type       syndication.Type
	Title      string
	Subscribed bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
