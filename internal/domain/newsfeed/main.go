package newsfeed

import "time"

// Entry represents an entry in the user's newsfeed
type Entry struct {
	UserID     string
	DocumentID string
	CreatedAt  time.Time
}

// Params --
func (e *Entry) Params() []interface{} {
	p := make([]interface{}, 3)
	p[0] = e.UserID
	p[1] = e.DocumentID
	p[2] = e.CreatedAt
	return p
}

// NewEntry creates an new entry
func NewEntry(userID string, documentID string) *Entry {
	return &Entry{
		UserID:     userID,
		DocumentID: documentID,
		CreatedAt:  time.Now(),
	}
}
