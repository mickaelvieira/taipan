package app

import (
	"os"
)

// IsDev is the app running in the development environment
func IsDev() bool {
	env := os.Getenv("TAIPAN_ENV")
	return env == "development" || env == ""
}
