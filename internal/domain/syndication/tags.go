package syndication

import "time"

// Tag syndication's tag
type Tag struct {
	ID        string
	Label     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
