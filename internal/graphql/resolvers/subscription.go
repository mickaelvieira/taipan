package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/subscription"
	"log"
)

// UserEventResolver resolves an bookmark event
type UserEventResolver struct {
	event *subscription.Event
}

// Item returns the event's message
func (r *UserEventResolver) Item() *UserResolver {
	u, ok := r.event.Payload.(*user.User)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a bookmark")
	}
	return &UserResolver{User: u}
}

// Emitter returns the event's emitter ID
func (r *UserEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *UserEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *UserEventResolver) Action() string {
	return string(r.event.Action)
}

// BookmarkEventResolver resolves an bookmark event
type BookmarkEventResolver struct {
	event *subscription.Event
}

// Item returns the event's message
func (r *BookmarkEventResolver) Item() *BookmarkResolver {
	b, ok := r.event.Payload.(*bookmark.Bookmark)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a bookmark")
	}
	return &BookmarkResolver{Bookmark: b}
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
	event *subscription.Event
}

// Item returns the event's message
func (r *DocumentEventResolver) Item() *DocumentResolver {
	d, ok := r.event.Payload.(*document.Document)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a document")
	}
	return &DocumentResolver{Document: d}
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

type userSubscriber struct {
	events chan<- *UserEventResolver
}

func (s *userSubscriber) Publish(e *subscription.Event) {
	s.events <- &UserEventResolver{event: e}
}

type bookmarkSubscriber struct {
	events chan<- *BookmarkEventResolver
}

func (s *bookmarkSubscriber) Publish(e *subscription.Event) {
	s.events <- &BookmarkEventResolver{event: e}
}

type documentSubscriber struct {
	events chan<- *DocumentEventResolver
}

func (s *documentSubscriber) Publish(e *subscription.Event) {
	s.events <- &DocumentEventResolver{event: e}
}

// User subscribes to user event
func (r *RootResolver) User(ctx context.Context) <-chan *UserEventResolver {
	c := make(chan *UserEventResolver)
	s := &userSubscriber{events: c}
	r.subscriptions.Subscribe(subscription.User, s, ctx.Done())
	return c
}

// Favorites subscribes to favorites feed bookmarksEvents
func (r *RootResolver) Favorites(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &bookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(subscription.Favorites, s, ctx.Done())
	return c
}

// ReadingList subscribes to reading list feed bookmarksEvents
func (r *RootResolver) ReadingList(ctx context.Context) <-chan *BookmarkEventResolver {
	c := make(chan *BookmarkEventResolver)
	s := &bookmarkSubscriber{events: c}
	r.subscriptions.Subscribe(subscription.ReadingList, s, ctx.Done())
	return c
}

// News subscribes to news feed bookmarksEvents
func (r *RootResolver) News(ctx context.Context) <-chan *DocumentEventResolver {
	c := make(chan *DocumentEventResolver)
	s := &documentSubscriber{events: c}
	r.subscriptions.Subscribe(subscription.News, s, ctx.Done())
	return c
}
