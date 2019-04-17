package main

import (
	"log"
	"net/http"
	"os"

	"github/mickaelvieira/taipan/internal/app"
)

func main() {
	app.LoadEnvironment()
	server := app.Bootstrap()

	port := os.Getenv("APP_PORT")
	env := os.Getenv("TAIPAN_ENV")
	webDir := os.Getenv("APP_WEB_DIR")

	if env == "development" || env == "" {
		fs := http.FileServer(http.Dir(webDir))
		http.Handle("/static/", fs)
	}

	// Routing
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/graphql", server.QueryHandler)
	// http.HandleFunc("/ws", websocketHandler)

	// Start the server
	log.Println("Listening: http://localhost:" + port)
	log.Println("Environment", env)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
