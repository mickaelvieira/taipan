package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"log"
)

// BookmarkEvent is a bookmarksSubscription event
type BookmarkEvent struct {
	event *Event
}

// Item returns the event's message
func (r *BookmarkEvent) Item() *BookmarkResolver {
	b, ok := r.event.Payload.(*bookmark.Bookmark)
	if !ok {
		log.Fatal("not a bookmark")
	}
	return &BookmarkResolver{Bookmark: b}
}

// ID returns the event's ID
func (r *BookmarkEvent) ID() string {
	return r.event.ID
}

// Topic returns the event's topic
func (r *BookmarkEvent) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *BookmarkEvent) Action() string {
	return string(r.event.Action)
}

// DocumentEvent is a bookmarksSubscription event
type DocumentEvent struct {
	event *Event
}

// Item returns the event's message
func (r *DocumentEvent) Item() *DocumentResolver {
	d, ok := r.event.Payload.(*document.Document)
	if !ok {
		log.Fatal("not a document")
	}
	return &DocumentResolver{Document: d}
}

// ID returns the event's ID
func (r *DocumentEvent) ID() string {
	return r.event.ID
}

// Topic returns the event's topic
func (r *DocumentEvent) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *DocumentEvent) Action() string {
	return string(r.event.Action)
}
