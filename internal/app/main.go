package app

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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
	}
}

// GetSession get the current session
func GetSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = GetSessionOptions()
	return sess
}

// Signal enables os signal catching
func Signal(onStop func()) {
	// Create a channel to handle os signals
	c := make(chan os.Signal, 1)
	// Send the following signals to the channel
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGTERM)

	// Clean up when we receive a signal
	go func() {
		select {
		case sig := <-c:
			logger.Warn(fmt.Sprintf("Signal received '%s'", sig))
			signal.Stop(c)
			onStop()
			os.Exit(1)
		}
	}()
}
