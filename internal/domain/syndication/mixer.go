package syndication

import (
	"github/mickaelvieira/taipan/internal/domain/url"
)

// Stack a stack of string
type Stack []string

// Stacks a stack of stack
type Stacks []Stack

// MakeMixerStack creates a stack of URLs
func MakeMixerStack(in []*url.URL) Stack {
	var s = make(Stack, len(in))
	for i, u := range in {
		s[i] = u.String()
	}
	return s
}

// count returns the number of elements in all stacks
func (s Stacks) count() (t int) {
	for _, v := range s {
		t = t + len(v)
	}
	return
}

// Mixup distributes evenly stacks' elements into a new stack
func Mixup(s Stacks) Stack {
	if s.count() == 0 {
		return make(Stack, 0)
	}

	t := s.count()
	o := make(Stack, t)
	c, i := 0, 0
	for i < t {
		var e string
		if len(s[c]) > 0 {
			e, s[c] = s[c][0], s[c][1:]
			o[i] = e
			i = i + 1
		}
		c = c + 1
		if c >= len(s) {
			c = 0
		}
	}
	return o
}
