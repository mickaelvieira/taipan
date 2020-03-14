package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type key int

const (
	// UserContextKey logged in user's context identifier
	UserContextKey key = iota

	// ClientIDContextKey client's context identifier
	ClientIDContextKey

	// LoadersContextKey Dataloaders identifier
	LoadersContextKey
)

// LoadEnvironment load environment variables
// See for details: https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func LoadEnvironment(path string) {
	env := os.Getenv("TAIPAN_ENV")
	if "" == env {
		env = "development"
	}

	if err := godotenv.Load(path + ".env." + env + ".local"); err != nil {
		log.Println(err)
	}
	if "test" != env {
		if err := godotenv.Load(path + ".env.local"); err != nil {
			log.Println(err)
		}
	}
	if err := godotenv.Load(path + ".env." + env); err != nil {
		log.Println(err)
	}
	if err := godotenv.Load(path + ".env"); err != nil {
		log.Println(err)
	}
}
