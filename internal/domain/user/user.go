package user

import (
	"strings"
	"time"
)

// User represents a single user wihthin the application
type User struct {
	ID        string
	Emails    []*Email
	Firstname string
	Lastname  string
	Image     *Image
	Theme     string
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

// Email user's email
type Email struct {
	ID          string
	Value       string
	IsPrimary   bool
	IsConfirmed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// IsEmailValid is the email valid?
func IsEmailValid(e string) bool {
	return len(e) > 0 && strings.Contains(e, "@")
}

// NewEmail creates a new user email address
func NewEmail(value string) *Email {
	return &Email{
		Value:     value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
