package bookmark

import (
	"time"
)

// MarkAsRead marks the user's bookmark as read
func MarkAsRead(b *UserBookmark) *UserBookmark {
	b.IsRead = READ
	b.UpdatedAt = time.Now()

	return b
}

// MarkAsUnead marks the user's bookmark as read
func MarkAsUnead(b *UserBookmark) *UserBookmark {
	b.IsRead = UNREAD
	b.UpdatedAt = time.Now()

	return b
}
