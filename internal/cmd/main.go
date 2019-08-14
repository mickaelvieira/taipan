package cmd

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

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
