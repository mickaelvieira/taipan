package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github/mickaelvieira/taipan/internal/logger"
)

func isDev() bool {
	env := os.Getenv("TAIPAN_ENV")
	return env == "development" || env == ""
}

// UseFileServer should the application serve assets
func UseFileServer() bool {
	return os.Getenv("APP_FILE_SERVER") != ""
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
