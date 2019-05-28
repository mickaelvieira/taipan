package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/auth"
)

func main() {
	app.LoadEnvironment()
	server := app.Bootstrap()

	port := os.Getenv("APP_PORT")
	env := os.Getenv("TAIPAN_ENV")
	webDir := os.Getenv("APP_WEB_DIR")

	if app.IsDev() {
		fs := http.FileServer(http.Dir(webDir))
		http.Handle("/static/", fs)
	}

	// Routing
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/graphql", auth.WithUser(server.QueryHandler, server.Repositories.Users))

	// Start the server
	fmt.Println("Listening: http://localhost:" + port)
	fmt.Println("Environment", env)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
