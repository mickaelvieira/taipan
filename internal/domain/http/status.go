package http

import (
	nethttp "net/http"
)

// IsNoLongerAvailable is the document still available
func IsNoLongerAvailable(s int) bool {
	return s == nethttp.StatusNotFound || s == nethttp.StatusGone
}

// IsServerError does the server return an error
func IsServerError(s int) bool {
	return s == nethttp.StatusNotAcceptable ||
		s == nethttp.StatusTooManyRequests ||
		s == nethttp.StatusInternalServerError
}

// IsError should the status code be consider an error
func IsError(s int) bool {
	return IsNoLongerAvailable(s) || IsServerError(s)
}
