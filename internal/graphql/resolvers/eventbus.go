package resolvers

// import (
// 	"math/rand"
// 	"time"
// )

// func randomID() string {
// 	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 	b := make([]rune, 16)
// 	for i := range b {
// 		b[i] = letter[rand.Intn(len(letter))]
// 	}
// 	return string(b)
// }

// // feedAction - an action identifies a type of operation that we want to process on a feed
// type feedAction string

// // List of actions
// const (
// 	add    feedAction = "Add"
// 	remove feedAction = "Remove"
// )

// // feedTopic - a topic identifies a feed
// // events are dispatched for a given topic
// type feedTopic string

// // List of topics
// const (
// 	// user        feedTopic = "User"
// 	news        feedTopic = "News"
// 	favorites   feedTopic = "Favorites"
// 	readingList feedTopic = "ReadingList"
// )

// // feedEvent represents an operation on a feed:
// // - an action
// // - a topic identifying a feed
// // - a payload, either a document or a bookmark
// type feedEvent struct {
// 	ID      string
// 	Emitter string
// 	Action  feedAction
// 	Topic   feedTopic
// 	Payload interface{}
// }

// func newFeedEvent(e string, t feedTopic, a feedAction, p interface{}) *feedEvent {
// 	return &feedEvent{
// 		ID:      randomID(),
// 		Emitter: e,
// 		Topic:   t,
// 		Action:  a,
// 		Payload: p,
// 	}
// }

// type subscriber interface {
// 	publish(e *feedEvent)
// }

// type bookmarkSubscriber struct {
// 	events chan<- *BookmarkEventResolver
// }

// func (s *bookmarkSubscriber) publish(e *feedEvent) {
// 	s.events <- &BookmarkEventResolver{event: e}
// }

// type documentSubscriber struct {
// 	events chan<- *DocumentEventResolver
// }

// func (s *documentSubscriber) publish(e *feedEvent) {
// 	s.events <- &DocumentEventResolver{event: e}
// }

// type subscribers map[string]subscriber

// type subscription struct {
// 	subscribers map[feedTopic]subscribers
// }

// func (bus *subscription) subscribe(t feedTopic, s subscriber, stop <-chan struct{}) {
// 	if bus.subscribers[t] == nil {
// 		bus.subscribers[t] = make(subscribers)
// 	}

// 	id := randomID()
// 	bus.subscribers[t][id] = s

// 	go func(id string, s <-chan struct{}) {
// 		for {
// 			select {
// 			case <-s:
// 				bus.unsubscribe(id)
// 				return
// 			case <-time.After(time.Second):
// 			}
// 		}
// 	}(id, stop)
// }

// func (bus *subscription) unsubscribe(id string) {
// 	for _, v := range bus.subscribers {
// 		for i := range v {
// 			if i == id {
// 				delete(v, id)
// 			}
// 		}
// 	}
// }

// func (bus *subscription) publish(e *feedEvent) {
// 	t := e.Topic
// 	if bus.subscribers[t] != nil {
// 		for _, s := range bus.subscribers[t] {
// 			go s.publish(e)
// 		}
// 	}
// }
