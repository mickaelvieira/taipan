package syndication

import (
	"github.com/mickaelvieira/taipan/internal/domain/messages"
	"github.com/mickaelvieira/taipan/internal/domain/url"
)

// Queue can be used to store a list of messages
type Queue []*messages.Document

// Shift returns the entity at the front of the queue
func (q *Queue) shift() (m *messages.Document) {
	s := *q
	if len(s) == 0 {
		return m
	}
	m, *q = s[0], s[1:]
	return
}

// MakeQueue creates a queue of URLs
func MakeQueue(in []*url.URL, id string) Queue {
	var s = make(Queue, len(in))
	for i, u := range in {
		s[i] = &messages.Document{
			Url:      u.String(),
			SourceId: id,
		}
	}
	return s
}

// Mixer mix up a list of web syndication source while
// keeping their position in their respective queue
type Mixer struct {
	queues []Queue
	idx    int
}

// Count returns the number of elements in all queues
func (s Mixer) Count() (t int) {
	for _, v := range s.queues {
		t = t + len(v)
	}
	return
}

// Push new entities into the mixer
func (s *Mixer) Push(q Queue) {
	s.queues[s.idx] = q
	s.idx = s.idx + 1
}

// IsFull returns the number of created queues
func (s *Mixer) IsFull() bool {
	return s.idx == len(s.queues)
}

// Mixup distributes evenly queues' entities into a new slice
func (s *Mixer) Mixup() Queue {
	if s.Count() == 0 {
		return make(Queue, 0)
	}

	t := s.Count()
	o := make(Queue, t)
	i, c := 0, 0
	for c < t {
		if len(s.queues[i]) > 0 {
			o[c] = s.queues[i].shift()
			c = c + 1
		}
		i = i + 1
		if i >= len(s.queues) {
			i = 0
		}
	}

	if c != t {
		panic("Logic error: We must have the exact same number of entities")
	}

	return o
}

// MakeMixer creates a mixer
func MakeMixer(size int) *Mixer {
	return &Mixer{
		queues: make([]Queue, size),
	}
}
