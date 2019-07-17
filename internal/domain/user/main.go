package user

import (
	"time"
)

// Status the user's status, whether or not the account has been approved
type Status int

// Status values
const (
	PENDING Status = iota
	APPROVED
)

// Email user's email
type Email struct {
	ID        string
	Email     string
	Primary   bool
	CreatedAt time.Time
	IpdatedAt time.Time
}

// User represents a single user wihthin the application
type User struct {
	ID        string
	Emails    []*Email
	Username  string
	Firstname string
	Lastname  string
	Password  string
	Image     *Image
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HasImage determine whether the user has an image associated to it
func (u *User) HasImage() bool {
	return u.Image != nil
}

// Image represents a user's avatar
type Image struct {
	Name   string
	Width  int32
	Height int32
	Format string
}

// SetDimensions image's information
func (i *Image) SetDimensions(w int, h int) {
	i.Width = int32(w)
	i.Height = int32(h)
}

// NewImage returns a document's image
func NewImage(name string, width int32, height int32, format string) *Image {
	return &Image{
		Name:   name,
		Width:  width,
		Height: height,
		Format: format,
	}
}
