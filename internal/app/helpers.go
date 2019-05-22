package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// IsDev is the app running in the development environment
func IsDev() bool {
	env := os.Getenv("TAIPAN_ENV")
	return env == "development" || env == ""
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
			fmt.Printf("Signal received '%s'\n", sig)
			signal.Stop(c)
			onStop()
			os.Exit(1)
		}
	}()
}
