package resolvers

import (
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

// Action what type of action we want to perform the feed
type Action string

// List of actions
const (
	Add    Action = "Add"
	Remove Action = "Remove"
)

// Topic Subscription topics
type Topic string

// List of topics
const (
	News        Topic = "News"
	Favorites   Topic = "Favorites"
	ReadingList Topic = "ReadingList"
)

// Subscriber defines what is a subscriber
type Subscriber interface {
	Publish(e *Event)
}

// BookmarkSubscriber Bookmark subscriber
type BookmarkSubscriber struct {
	Events chan<- *BookmarkEvent
}

// Publish publishes an bookmark event
func (s *BookmarkSubscriber) Publish(e *Event) {
	log.Printf("bookmark subscriber: received events %s", e.Topic)
	be := &BookmarkEvent{
		event: e,
	}

	select {
	case s.Events <- be:
	case <-time.After(time.Second):
	}
}

// DocumentSubscriber handles the pool of documentsEvents
type DocumentSubscriber struct {
	Events chan<- *DocumentEvent
}

// Publish publishes a document event
func (s *DocumentSubscriber) Publish(e *Event) {
	log.Printf("document subscriber: received events %s", e.Topic)
	de := &DocumentEvent{
		event: e,
	}

	select {
	case s.Events <- de:
	case <-time.After(time.Second):
	}
}

// Event defines what is a event
type Event struct {
	ID      string
	Action  Action
	Topic   Topic
	Payload interface{}
}

// NewEvent creates a new event
func NewEvent(t Topic, a Action, p interface{}) *Event {
	return &Event{
		ID:      randomID(),
		Topic:   t,
		Action:  a,
		Payload: p,
	}
}

// PubSub interface
type PubSub interface {
	Subscribe(t Topic, s Subscriber)
	Unsubscribe(id string)
	Publish(e Event)
}

// Subscription bus
type Subscription struct {
	subscribers map[Topic]map[string]Subscriber
}

// Subscribe subscribes a subscriber to a topic
func (ps *Subscription) Subscribe(t Topic, s Subscriber, stop <-chan struct{}) {
	if ps.subscribers[t] == nil {
		ps.subscribers[t] = make(map[string]Subscriber)
	}

	id := randomID()

	log.Printf("Subscribe %s %s", id, t)
	ps.subscribers[t][id] = s

	go func(id string, s <-chan struct{}) {
		select {
		case <-s:
			ps.Unsubscribe(id)
		case <-time.After(time.Second):
		}
	}(id, stop)
}

// Unsubscribe unsubscribes a subscriber from all topic
func (ps *Subscription) Unsubscribe(id string) {
	for _, v := range ps.subscribers {
		for i := range v {
			if i == id {
				delete(v, id)
			}
		}
	}
}

// Publish notifies subscribers of an event of a specific topic
func (ps *Subscription) Publish(e *Event) {
	log.Printf("publish event %s", e.Topic)
	for _, s := range ps.subscribers[e.Topic] {
		s.Publish(e)
	}
}
