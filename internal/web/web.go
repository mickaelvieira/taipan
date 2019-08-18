package web

import (
	"errors"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// ErrServer application level errors
var (
	ErrServer = errors.New("Something went wrong. Please try again later")
)

// IsDev is the app running with development mode
func IsDev() bool {
	env := os.Getenv("TAIPAN_ENV")
	return env == "development" || env == ""
}

// UseFileServer should the application serve assets
func UseFileServer() bool {
	return os.Getenv("APP_FILE_SERVER") != ""
}

// GetSessionOptions returns session's options
func GetSessionOptions() *sessions.Options {
	return &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   !IsDev(),
	}
}

// GetSession get the current session
func GetSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = GetSessionOptions()
	return sess
}
