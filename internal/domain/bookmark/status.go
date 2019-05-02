package bookmark

import (
	"time"
)

// MarkAsRead marks the user's bookmark as read
func MarkAsRead(b *UserBookmark) *UserBookmark {
	b.IsRead = true
	b.UpdatedAt = time.Now()

	return b
}

// MarkAsUnead marks the user's bookmark as read
func MarkAsUnead(b *UserBookmark) *UserBookmark {
	b.IsRead = false
	b.UpdatedAt = time.Now()

	return b
}
