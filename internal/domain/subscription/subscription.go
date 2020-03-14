package subscription

import (
	"github.com/mickaelvieira/taipan/internal/domain/http"
	"github.com/mickaelvieira/taipan/internal/domain/syndication"
	"github.com/mickaelvieira/taipan/internal/domain/url"
	"time"
)

// Subscription represents a user'subscription to a syndication source
type Subscription struct {
	ID         string
	UserID     string
	URL        *url.URL
	Domain     *url.URL
	Type       syndication.Type
	Title      string
	Subscribed bool
	Frequency  http.Frequency
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
