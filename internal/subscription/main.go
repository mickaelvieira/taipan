package subscription

import (
	"math/rand"
	"time"
)

// Topic - a topic identifies a subscription topic
// events are dispatched for a given topic
type Topic string

// List of topics
const (
	User        Topic = "User"
	News        Topic = "News"
	Favorites   Topic = "Favorites"
	ReadingList Topic = "ReadingList"
)

// Action - an action identifies a type of operation
type Action string

// List of actions
const (
	Add    Action = "Add"
	Update Action = "Update"
	Remove Action = "Remove"
)

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Event represents an operation:
// - an action
// - an emitter (.ie client ID)
// - a topic
// - a payload
type Event struct {
	Emitter string
	Action  Action
	Topic   Topic
	Payload interface{}
}

// NewEvent creates a subscription event
func NewEvent(e string, t Topic, a Action, p interface{}) *Event {
	return &Event{
		Emitter: e,
		Topic:   t,
		Action:  a,
		Payload: p,
	}
}

// NewEventBus creates a new event bus
func NewEventBus() *Subscription {
	return &Subscription{
		subscribers: make(map[Topic]Subscribers),
	}
}

// Subscriber defined what a subscriber looks like
type Subscriber interface {
	Publish(e *Event)
}

// Subscribers a list of subscribers identified by their ID
type Subscribers map[string]Subscriber

// Subscription is an event bus
type Subscription struct {
	subscribers map[Topic]Subscribers
}

// Subscribe adds a subscriber to the bus
func (bus *Subscription) Subscribe(t Topic, s Subscriber, stop <-chan struct{}) {
	if bus.subscribers[t] == nil {
		bus.subscribers[t] = make(Subscribers)
	}

	id := randomID()
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

// Unsubscribe removes a subscriber from the bus
func (bus *Subscription) Unsubscribe(id string) {
	for _, v := range bus.subscribers {
		for i := range v {
			if i == id {
				delete(v, id)
			}
		}
	}
}

// Publish publishes an event to subscribers
func (bus *Subscription) Publish(e *Event) {
	t := e.Topic
	if bus.subscribers[t] != nil {
		for _, s := range bus.subscribers[t] {
			go s.Publish(e)
		}
	}
}
