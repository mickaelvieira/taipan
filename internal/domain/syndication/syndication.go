package syndication

import (
	"errors"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
	"strings"
	"time"
)

// Syndication domain errors
var (
	ErrXMLTypeIsNotValid  = errors.New("The XML type is not valid")
	ErrFeedTypeISNotValid = errors.New("The feed type is not valid")
)

// DefaultWPFeedTitle a default title for WP feeds
// @TODO I need to remove this shit, it brings more problems than solutions
const DefaultWPFeedTitle = "wordpress feed"

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

// GetSourceType returns the feed type based on a provided string
func GetSourceType(t string) (Type, error) {
	t = strings.ToLower(t)
	if isRSS(t) {
		return RSS, nil
	}
	if isAtom(t) {
		return ATOM, nil
	}
	return INVALID, ErrXMLTypeIsNotValid
}

// FromGoFeedType returns the feed type based on the gofeed type
func FromGoFeedType(t string) (Type, error) {
	t = strings.ToLower(t)
	if t == "rss" {
		return RSS, nil
	}
	if t == "atom" {
		return ATOM, nil
	}
	return INVALID, ErrFeedTypeISNotValid
}

var blacklist = []string{"github.com"}

// IsBlacklisted checks whether the feed's URL matches a pattern that is black listed
func IsBlacklisted(url string) bool {
	for _, v := range blacklist {
		if strings.Index(url, v) != -1 {
			return true
		}
	}
	return false
}

// Source represents what is the feed within the application
type Source struct {
	ID        string
	URL       *url.URL
	Domain    *url.URL
	Type      Type
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	ParsedAt  time.Time
	IsDeleted bool
	IsPaused  bool
	Frequency http.Frequency
}

// NewSource creates a new syndication source
func NewSource(url *url.URL, title string, feedType Type) *Source {
	return &Source{
		URL:       url,
		Title:     title,
		Type:      feedType,
		IsPaused:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Frequency: http.Hourly,
	}
}

// UserSource represents a feed from a user's prespective
type UserSource struct {
	Source
	AddedAt time.Time
}
