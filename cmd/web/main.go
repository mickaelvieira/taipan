package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"text/template"

	"github/mickaelvieira/taipan/internal/graphql"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
)

var schema = graphql.LoadAndParseSchema()
var templates = template.Must(template.New("html-tmpl").ParseGlob("./web/templates/*.html"))

type Data struct {
}

func queryHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	handler := &relay.Handler{Schema: schema}
	handler.ServeHTTP(w, req.WithContext(ctx))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", Data{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func loadEnvironment() {
	env := os.Getenv("TAIPAN_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}

func main() {
	loadEnvironment()
	port := os.Getenv("APP_PORT")

	// Serve static files
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", fs)

	// Routing
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/graphql", queryHandler)
	// http.HandleFunc("/ws", websocketHandler)

	// Start the server
	log.Println("Listening: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
