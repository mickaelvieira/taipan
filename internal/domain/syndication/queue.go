package syndication

// QueueSources a queue of sources
type QueueSources []*Source

// Shift returns the entity at the front of the queue
func (q *QueueSources) Shift() (e *Source) {
	s := *q
	if len(s) == 0 {
		return e
	}
	e, *q = s[0], s[1:]
	return
}
