package syndication

import "github/mickaelvieira/taipan/internal/domain/url"

type queue []string

// Shift returns the entity at the front of the queue
func (q *queue) shift() (e string) {
	s := *q
	if len(s) == 0 {
		return e
	}
	e, *q = s[0], s[1:]
	return
}

// makeQueue creates a queue of URLs
func makeQueue(in []*url.URL) queue {
	var s = make(queue, len(in))
	for i, u := range in {
		s[i] = u.String()
	}
	return s
}

// Mixer mix up a list of web syndication source while
// keeping their position in their respective queue
type Mixer struct {
	queues []queue
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
func (s *Mixer) Push(in []*url.URL) {
	s.queues[s.idx] = makeQueue(in)
	s.idx = s.idx + 1
}

// IsFull returns the number of created queues
func (s *Mixer) IsFull() bool {
	return s.idx == len(s.queues)
}

// Mixup distributes evenly queues' entities into a new slice
func (s *Mixer) Mixup() []string {
	if s.Count() == 0 {
		return make([]string, 0)
	}

	t := s.Count()
	o := make([]string, t)
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
		queues: make([]queue, size),
	}
}
