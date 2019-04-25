package app

import (
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/gql"
	"html/template"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment load environment variables
// See for details: https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func LoadEnvironment() {
	root := "../../"
	env := os.Getenv("TAIPAN_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(root + ".env." + env + ".local")
	if "test" != env {
		godotenv.Load(root + ".env.local")
	}
	godotenv.Load(root + ".env." + env)
	godotenv.Load(root + ".env")
}

// Bootstrap the application
func Bootstrap() *Server {
	var webDir = os.Getenv("APP_WEB_DIR")
	var templates = template.Must(template.New("html-tmpl").ParseGlob(webDir + "/templates/*.html"))
	var schema = gql.LoadAndParseSchema(webDir + "/graphql/schema.graphql")
	var assetsDef = assets.LoadAssetsDefinition(webDir + "/static/hashes.json")

	var server = Server{templates: templates, schema: schema, assets: assetsDef}

	return &server
}
