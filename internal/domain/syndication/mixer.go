package syndication

import "github/mickaelvieira/taipan/internal/domain/url"

type fifo []string

// Shift returns the entity at the front of the queue
func (q *fifo) shift() (e string) {
	s := *q
	if len(s) == 0 {
		return e
	}
	e, *q = s[0], s[1:]
	return
}

// makeQueue creates a queue of URLs
func makeQueue(in []*url.URL) fifo {
	var s = make(fifo, len(in))
	for i, u := range in {
		s[i] = u.String()
	}
	return s
}

// Mixer mix up a list of web syndication source while
// keeping their position in their respective queue
type Mixer struct {
	q []fifo
	i int
}

// count returns the number of elements in all queues
func (s Mixer) count() (t int) {
	for _, v := range s.q {
		t = t + len(v)
	}
	return
}

// Push new entities into the mixer
func (s *Mixer) Push(in []*url.URL) {
	q := makeQueue(in)
	s.q[s.i] = q
	s.i = s.i + 1
}

// Mixup distributes evenly queues' entities into a new slice
func (s *Mixer) Mixup() []string {
	if s.count() == 0 {
		return make([]string, 0)
	}

	t := s.count()
	o := make([]string, t)
	i, c := 0, 0
	for c < t {
		if len(s.q[i]) > 0 {
			o[c] = s.q[i].shift()
			c = c + 1
		}
		i = i + 1
		if i >= len(s.q) {
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
		q: make([]fifo, size),
	}
}
