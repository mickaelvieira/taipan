package feed

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"strings"
	"time"
)

// Type represents whether the feed is an atom or rss feed
type Type string

// Type values
const (
	ATOM    Type = "application/atom+xml"
	RSS     Type = "application/rss+xml"
	INVALID Type = ""
)

func isRSS(t string) bool {
	return t == string(RSS)
}

func isAtom(t string) bool {
	return t == string(ATOM)
}

// GetFeedType returns the feed type based on a provided string
func GetFeedType(t string) (Type, error) {
	t = strings.ToLower(t)
	if isRSS(t) {
		return RSS, nil
	}
	if isAtom(t) {
		return ATOM, nil
	}
	return INVALID, fmt.Errorf("Invalid feed type %s", t)
}

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
	URL       *uri.URI
	Type      Type
	Title     string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// New creates a new Feed with a UUID
func New(url *uri.URI, title string, feedType Type) Feed {
	return Feed{
		URL:       url,
		Title:     title,
		Type:      feedType,
		Status:    NEW,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UserFeed represents a feed from a user's prespective
type UserFeed struct {
	Feed
	AddedAt time.Time
}
