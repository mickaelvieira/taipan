package app

import (
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/graphql"
	"github/mickaelvieira/taipan/internal/repository"
	"html/template"
	"os"
)

// Bootstrap the application
func Bootstrap() *Server {
	var webDir = os.Getenv("APP_WEB_DIR")
	var templates = template.Must(template.New("html-tmpl").ParseGlob(webDir + "/templates/*.html"))
	var repositories = repository.GetRepositories()
	var schema = graphql.LoadAndParseSchema(webDir+"/graphql/schema.graphql", repositories)
	var assetsDef = assets.LoadAssetsDefinition(webDir + "/static/hashes.json")

	var server = Server{
		templates:    templates,
		schema:       schema,
		assets:       assetsDef,
		Repositories: repositories,
	}

	return &server
}
