package internal

import (
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/graphql"
	"html/template"
	"net/http"
	"os"

	gogql "github.com/graph-gophers/graphql-go"

	"github.com/joho/godotenv"
)

// https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func LoadEnvironment() {
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

// Bootstrap the application
func Bootstrap() Server {
	var webDir = os.Getenv("APP_WEB_DIR")
	var templates = template.Must(template.New("html-tmpl").ParseGlob(webDir + "/templates/*.html"))
	var schema = graphql.LoadAndParseSchema(webDir + "/graphql/schema.graphql")
	var asset = assets.LoadAssetsDefinition(webDir + "/static/hashes.json")

	var server = Server{templates: templates, schema: schema, assets: asset}

	return server
}

// Server is the main application
type Server struct {
	templates *template.Template
	schema    *gogql.Schema
	assets    *assets.Assets
}

// IndexHandler is the method to handle / route
func (s *Server) IndexHandler(w http.ResponseWriter, req *http.Request) {
	err := s.templates.ExecuteTemplate(w, "index.html", s.assets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
