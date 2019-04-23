package uuid

import "github.com/google/uuid"

// New create a new UUID
func New() string {
	return uuid.New().String()
}
