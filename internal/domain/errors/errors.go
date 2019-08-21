package errors

// DomainError contract
type DomainError interface {
	// Error returns the message of the error that caused the domain error
	Error() string

	// HasReason was this domain error caused by another error
	HasReason() bool

	// Reason returns the error that caused the domain error
	Reason() error

	// DomainError returns the domain error
	Domain() error
}

// Domain error
type domainError struct {
	err    error
	reason error
}

// Error returns the message of the error that caused the domain error
func (d *domainError) Error() string {
	if d.HasReason() {
		return d.reason.Error()
	}
	return d.err.Error()
}

// Reason returns the error that caused the domain error
func (d *domainError) Reason() error {
	return d.reason
}

// HasReason was this domain error caused by another error
func (d *domainError) HasReason() bool {
	return d.reason != nil
}

// DomainError returns the domain error
func (d *domainError) Domain() error {
	return d.err
}

// New creates a new domain error
func New(err error, reason error) DomainError {
	return &domainError{
		err:    err,
		reason: reason,
	}
}
