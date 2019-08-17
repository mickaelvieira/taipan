package newsfeed

import "time"

// Entry represents an entry in the user's newsfeed
type Entry struct {
	UserID     string
	DocumentID string
	CreatedAt  time.Time
}

// NewEntry creates an new entry
func NewEntry(userID string, documentID string) *Entry {
	return &Entry{
		UserID:     userID,
		DocumentID: documentID,
		CreatedAt:  time.Now(),
	}
}
