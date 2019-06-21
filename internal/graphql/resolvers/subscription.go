package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"log"
	"math/rand"
	"time"
)

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// FeedAction - an action identifies a type of operation that we want to process on a feed
type FeedAction string

// List of actions
const (
	Add    FeedAction = "Add"
	Remove FeedAction = "Remove"
)

// FeedTopic - a topic identifies a feed
// events are dispatched for a given topic
type FeedTopic string

// List of topics
const (
	News        FeedTopic = "News"
	Favorites   FeedTopic = "Favorites"
	ReadingList FeedTopic = "ReadingList"
)

// FeedEvent represents an operation on a feed:
// - an action
// - a topic identifying a feed
// - a payload, either a document or a bookmark
type FeedEvent struct {
	ID      string
	Action  FeedAction
	Topic   FeedTopic
	Payload interface{}
}

// NewFeedEvent creates a new event
func NewFeedEvent(t FeedTopic, a FeedAction, p interface{}) *FeedEvent {
	return &FeedEvent{
		ID:      randomID(),
		Topic:   t,
		Action:  a,
		Payload: p,
	}
}

// BookmarkEventResolver resolves an bookmark event
type BookmarkEventResolver struct {
	event *FeedEvent
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
	event *FeedEvent
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

// Topic returns the event's topic
func (r *DocumentEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *DocumentEventResolver) Action() string {
	return string(r.event.Action)
}

// Subscriber defines what is a subscriber
type Subscriber interface {
	Publish(e *FeedEvent)
}

// BookmarkSubscriber publishes bookmark events to the GraphQL subscription channel
type BookmarkSubscriber struct {
	events chan<- *BookmarkEventResolver
}

// Publish publishes a bookmark event
func (s *BookmarkSubscriber) Publish(e *FeedEvent) {
	log.Printf("Bookmark subscriber: received events %s", e.Topic)
	s.events <- &BookmarkEventResolver{event: e}
}

// DocumentSubscriber publishes document events to the GraphQL subscription channel
type DocumentSubscriber struct {
	events chan<- *DocumentEventResolver
}

// Publish publishes a document event
func (s *DocumentSubscriber) Publish(e *FeedEvent) {
	log.Printf("Document subscriber: received events %s", e.Topic)
	s.events <- &DocumentEventResolver{event: e}
}

// Subscribers containers the list of subscribers for a given topic
type Subscribers map[string]Subscriber

// Subscription bus
type Subscription struct {
	subscribers map[FeedTopic]Subscribers
}

// Subscribe subscribes a subscriber to a topic
func (bus *Subscription) Subscribe(t FeedTopic, s Subscriber, stop <-chan struct{}) {
	if bus.subscribers[t] == nil {
		bus.subscribers[t] = make(Subscribers)
	}

	id := randomID()

	log.Printf("Subscribe with id [%s] to topic [%s]", id, t)
	bus.subscribers[t][id] = s

	go func(id string, s <-chan struct{}) {
		for {
			select {
			case <-s:
				bus.Unsubscribe(id)
				return
			case <-time.After(time.Second):
			}
		}
	}(id, stop)
}

// Unsubscribe unsubscribes a subscriber from all topic
func (bus *Subscription) Unsubscribe(id string) {
	for _, v := range bus.subscribers {
		for i := range v {
			if i == id {
				log.Printf("Unsubscribe [%s]", id)
				delete(v, id)
			}
		}
	}
}

// Publish notifies subscribers of an event of a specific topic
func (bus *Subscription) Publish(e *FeedEvent) {
	t := e.Topic
	log.Printf("Publish event: topic [%s], action [%s]", t, e.Action)
	log.Printf("Number of subscriber: [%d]", len(bus.subscribers[t]))
	if bus.subscribers[t] != nil {
		for _, s := range bus.subscribers[t] {
			go s.Publish(e)
		}
	}
}
