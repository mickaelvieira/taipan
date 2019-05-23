package bookmark

import (
	"errors"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"time"
)

// ReadStatus are the values defineing whether or not a bookmark has been read
type ReadStatus bool

// ReadStatus values
const (
	UNREAD ReadStatus = false
	READ   ReadStatus = true
)

// Value converts the value going into the DB
// func (s ReadStatus) Value() (driver.Value, error) {
// 	var v int64
// 	if s == READ {
// 		v = 1
// 	}
// 	return v, nil
// }

// Scan converts the value coming from the DB
func (s *ReadStatus) Scan(value interface{}) error {
	if value == nil {
		*s = UNREAD
		return nil
	}
	if v, ok := value.(int64); ok {
		if v == 1 {
			*s = READ
		} else {
			*s = UNREAD
		}
		return nil
	}
	return errors.New("failed to scan read status")
}

// ReadStatusFromBoolean returns the read status based on a boolean
func ReadStatusFromBoolean(status bool) ReadStatus {
	if status {
		return READ
	}
	return UNREAD
}

// Bookmark struct represents what is a bookmark from a user's perspective
type Bookmark struct {
	ID          string
	URL         *uri.URI
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *image.Image
	AddedAt     time.Time
	UpdatedAt   time.Time
	IsRead      ReadStatus
	IsLinked    bool
}
