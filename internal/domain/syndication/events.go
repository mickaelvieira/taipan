package syndication

// Event represents an event in the syndication lifecycle
type Event string

const (
	// TypeWasUndefined The parse could not determine the feed type, .ie RSS or Atom
	TypeWasUndefined Event = "TypeWasUndefined"

	// SourceWasUnreachable The HTTP request returned a 404 or 410 error
	SourceWasUnreachable Event = "SourceWasUnreachable"

	// ServerWasUnreachable The server returned an error
	ServerWasUnreachable Event = "ServerWasUnreachable"
)
