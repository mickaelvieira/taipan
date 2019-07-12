package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"log"
)

// BookmarkEventResolver resolves an bookmark event
type BookmarkEventResolver struct {
	event *feedEvent
}

// Item returns the event's message
func (r *BookmarkEventResolver) Item() *BookmarkResolver {
	b, ok := r.event.Payload.(*bookmark.Bookmark)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a bookmark")
	}
	return &BookmarkResolver{Bookmark: b}
}

// ID returns the event's ID
func (r *BookmarkEventResolver) ID() string {
	return r.event.ID
}

// Emitter returns the event's emitter ID
func (r *BookmarkEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *BookmarkEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *BookmarkEventResolver) Action() string {
	return string(r.event.Action)
}

// DocumentEventResolver is a document event
type DocumentEventResolver struct {
	event *feedEvent
}

// Item returns the event's message
func (r *DocumentEventResolver) Item() *DocumentResolver {
	d, ok := r.event.Payload.(*document.Document)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a document")
	}
	return &DocumentResolver{Document: d}
}

// ID returns the event's ID
func (r *DocumentEventResolver) ID() string {
	return r.event.ID
}

// Emitter returns the event's emitter ID
func (r *DocumentEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *DocumentEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *DocumentEventResolver) Action() string {
	return string(r.event.Action)
}

// Favorites subscribes to favorites feed bookmarksEvents
func (r *RootResolver) Favorites(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &bookmarkSubscriber{events: c}
	r.subscriptions.subscribe(favorites, s, ctx.Done())
	return c
}

// ReadingList subscribes to reading list feed bookmarksEvents
func (r *RootResolver) ReadingList(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &bookmarkSubscriber{events: c}
	r.subscriptions.subscribe(readingList, s, ctx.Done())
	return c
}

// News subscribes to news feed bookmarksEvents
func (r *RootResolver) News(ctx context.Context) <-chan *DocumentEventResolver {
	c := make(chan *DocumentEventResolver)
	s := &documentSubscriber{events: c}
	r.subscriptions.subscribe(news, s, ctx.Done())
	return c
}
