package bookmark

import "time"

// Bookmark is a model
type Bookmark struct {
	ID          string
	URL         string
	Hash        string
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AddedAt     time.Time
	IsRead      bool
}

// Status defines the status of a bookmark
type Status string

// Status values
const (
	FETCHED  Status = "fetched"
	PENDING  Status = "pending"
	FETCHING Status = "fetching"
	FAILED   Status = "failed"
)
