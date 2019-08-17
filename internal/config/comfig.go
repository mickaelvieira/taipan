package config

import (
	"os"

	"github.com/joho/godotenv"
)

type key int

const (
	// UserContextKey logged in user's context identifier
	UserContextKey key = iota

	// ClientIDContextKey client's context identifier
	ClientIDContextKey
)

// LoadEnvironment load environment variables
// See for details: https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func LoadEnvironment(path string) {
	env := os.Getenv("TAIPAN_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(path + ".env." + env + ".local")
	if "test" != env {
		godotenv.Load(path + ".env.local")
	}
	godotenv.Load(path + ".env." + env)
	godotenv.Load(path + ".env")
}
