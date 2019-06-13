package resolvers

import (
	"github/mickaelvieira/taipan/internal/repository"
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

// Resolvers resolvers
type Resolvers struct {
	repositories *repository.Repositories
	subscriber   chan *Subscriber
	events       chan *SubEvent
}

func (r *Resolvers) broadcast() {
	subscribers := map[string]*Subscriber{} // map of subscribers
	unsubscribe := make(chan string)        // unsubscribe channel

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.subscriber:
			id := randomID()
			log.Printf("new subscriber %s\n", id)
			subscribers[id] = s
			log.Printf("subscribers %d\n", len(subscribers))
		case e := <-r.events:
			for id, s := range subscribers {
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

// SubEvent is a subscription event
type SubEvent struct {
	id  string
	msg string
}

// Msg returns the event's message
func (r *SubEvent) Msg() string {
	return r.msg
}

// ID returns the event's ID
func (r *SubEvent) ID() string {
	return r.id
}

// Subscriber handles the pool of events
type Subscriber struct {
	stop   <-chan struct{}
	events chan<- *SubEvent
}

// GetRootResolver returns the root resolver. Queries and mutations are methods of this resolver
func GetRootResolver(repositories *repository.Repositories) (r *Resolvers) {
	r = &Resolvers{
		repositories: repositories,
		events:       make(chan *SubEvent),
		subscriber:   make(chan *Subscriber),
	}

	// initialiaze subscriptions
	// go r.broadcast()

	return
}
