package user

import "time"

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
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}
