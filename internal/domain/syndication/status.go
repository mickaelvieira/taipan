package syndication

import (
	"net/http"
)

var unrecoverable = []int{
	http.StatusUnauthorized,
	http.StatusPaymentRequired,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusGone,
	http.StatusRequestURITooLong,
	http.StatusUnsupportedMediaType,
	http.StatusMisdirectedRequest,
	http.StatusUnprocessableEntity,
	http.StatusLocked,
	http.StatusFailedDependency,
	http.StatusUpgradeRequired,
	http.StatusPreconditionRequired,
	http.StatusUnavailableForLegalReasons,
}

// IsHTTPErrorPermanent is the document still available
func IsHTTPErrorPermanent(s int) bool {
	for _, v := range unrecoverable {
		if s == v {
			return true
		}
	}
	return false
}

// IsHTTPErrorTemporary does the server return an error
func IsHTTPErrorTemporary(s int) bool {
	return s >= http.StatusInternalServerError || s >= http.StatusBadRequest && !IsHTTPErrorPermanent(s)
}

// IsHTTPError should the status code be consider an error
func IsHTTPError(s int) bool {
	return IsHTTPErrorPermanent(s) || IsHTTPErrorTemporary(s)
}
