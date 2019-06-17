package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/repository"
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

// RootResolver resolvers
type RootResolver struct {
	repositories *repository.Repositories
	subscription chan *Subscriber
	events       chan *BookmarkEvent
}

func (r *RootResolver) broadcast() {
	subscribers := map[string]*Subscriber{} // map of subscribers
	unsubscribe := make(chan string)        // unsubscribe channel

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.subscription:
			subscribers[randomID()] = s
		case e := <-r.events:
			for id, s := range subscribers {
				// log.Printf("Event topic %s", e.topic)
				// log.Printf("Subsciber topic %s", s.topic)
				if e.topic != s.topic {
					continue
				}

				go func(id string, s *Subscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

// BookmarkEvent is a subscription event
type BookmarkEvent struct {
	id       string
	bookmark *bookmark.Bookmark
	topic    Topic
	action   Action
}

// Bookmark returns the event's message
func (r *BookmarkEvent) Bookmark() *BookmarkResolver {
	return &BookmarkResolver{Bookmark: r.bookmark}
}

// ID returns the event's ID
func (r *BookmarkEvent) ID() string {
	return r.id
}

// Topic returns the event's topic
func (r *BookmarkEvent) Topic() string {
	return string(r.topic)
}

// Action returns the event's action
func (r *BookmarkEvent) Action() string {
	return string(r.action)
}

// Subscriber handles the pool of events
type Subscriber struct {
	stop   <-chan struct{}
	events chan<- *BookmarkEvent
	topic  Topic
}

// Topic Subscription topics
type Topic string

// List of topics
const (
	News        Topic = "News"
	Favorites   Topic = "Favorites"
	ReadingList Topic = "ReadingList"
)

// Action what type of action we want to perform the feed
type Action string

// List of actions
const (
	Add    Action = "Add"
	Remove Action = "Remove"
)

// GetRootResolver returns the root resolver. Queries and mutations are methods of this resolver
func GetRootResolver(repositories *repository.Repositories) (r *RootResolver) {
	r = &RootResolver{
		repositories: repositories,
		events:       make(chan *BookmarkEvent),
		subscription: make(chan *Subscriber),
	}

	// initialiaze subscriptions
	go r.broadcast()

	return
}
