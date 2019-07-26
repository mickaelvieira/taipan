package logger

import (
	"fmt"
	"os"
	"log"
)

func isDev() bool {
	env := os.Getenv("TAIPAN_ENV")
	return env == "development" || env == ""
}

// Info prints stuff in dev mode only
func Info(message string) {
	if isDev() {
		fmt.Println(message)
	}
}

// Warn prints warning messages
func Warn(message string) {
	fmt.Println(message)
}

// Error prints error messages
func Error(err interface{}) {
	log.Println(err)
}
