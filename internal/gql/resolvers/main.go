package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
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
	repositories          *repository.Repositories
	bookmarksSubscription chan *BookmarkSubscriber
	documentsSubscription chan *DocumentSubscriber
	bookmarksEvents       chan *BookmarkEvent
	documentsEvents       chan *DocumentEvent
}

func (r *RootResolver) broadcast() {
	bookmarks := map[string]*BookmarkSubscriber{} // map of subscribers
	documents := map[string]*DocumentSubscriber{} // map of subscribers
	unsubscribeBookmarks := make(chan string)     // unsubscribe channel
	unsubscribeDocuments := make(chan string)     // unsubscribe channel

	for {
		select {
		case id := <-unsubscribeBookmarks:
			delete(bookmarks, id)
		case id := <-unsubscribeDocuments:
			delete(documents, id)
		case s := <-r.bookmarksSubscription:
			bookmarks[randomID()] = s
		case s := <-r.documentsSubscription:
			documents[randomID()] = s
		case e := <-r.bookmarksEvents:
			// log.Printf("====>>> Bookmark Event received %s with topic %s\n", e.id, e.topic)
			for id, s := range bookmarks {
				// log.Printf("Subscriber topic %s\n", s.topic)
				if e.topic != s.topic {
					// log.Println("Skipped!")
					continue
				}

				go func(id string, s *BookmarkSubscriber) {
					select {
					case <-s.stop:
						unsubscribeBookmarks <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		case e := <-r.documentsEvents:
			// log.Printf("====>>> Document Event received %s with topic %s\n", e.id, e.topic)
			for id, s := range documents {
				// log.Printf("Subscriber topic %s\n", s.topic)
				if e.topic != s.topic {
					// log.Println("Skipped!")
					continue
				}

				go func(id string, s *DocumentSubscriber) {
					select {
					case <-s.stop:
						unsubscribeDocuments <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

// BookmarkEvent is a bookmarksSubscription event
type BookmarkEvent struct {
	id       string
	bookmark *bookmark.Bookmark
	topic    Topic
	action   Action
}

// Item returns the event's message
func (r *BookmarkEvent) Item() *BookmarkResolver {
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

// DocumentEvent is a bookmarksSubscription event
type DocumentEvent struct {
	id       string
	document *document.Document
	topic    Topic
	action   Action
}

// Item returns the event's message
func (r *DocumentEvent) Item() *DocumentResolver {
	return &DocumentResolver{Document: r.document}
}

// ID returns the event's ID
func (r *DocumentEvent) ID() string {
	return r.id
}

// Topic returns the event's topic
func (r *DocumentEvent) Topic() string {
	return string(r.topic)
}

// Action returns the event's action
func (r *DocumentEvent) Action() string {
	return string(r.action)
}

// BookmarkSubscriber handles the pool of bookmarksEvents
type BookmarkSubscriber struct {
	stop   <-chan struct{}
	events chan<- *BookmarkEvent
	topic  Topic
}

// DocumentSubscriber handles the pool of documentsEvents
type DocumentSubscriber struct {
	stop   <-chan struct{}
	events chan<- *DocumentEvent
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
		repositories:          repositories,
		bookmarksEvents:       make(chan *BookmarkEvent),
		documentsEvents:       make(chan *DocumentEvent),
		bookmarksSubscription: make(chan *BookmarkSubscriber),
		documentsSubscription: make(chan *DocumentSubscriber),
	}

	// initialiaze subscriptions
	go r.broadcast()

	return
}
