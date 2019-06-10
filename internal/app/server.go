package app

import (
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/repository"
	"html/template"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
)

// Server is the main application
type Server struct {
	templates    *template.Template
	schema       *graphql.Schema
	assets       *assets.Assets
	Repositories *repository.Repositories
}

type tmplData struct {
	Assets *assets.Assets
	CDN    string
}

// IndexHandler is the method to handle / route
func (s *Server) IndexHandler(w http.ResponseWriter, req *http.Request) {
	err := s.templates.ExecuteTemplate(w, "index.html", tmplData{
		Assets: s.assets,
		CDN:    assets.MakeCDNBaseURL(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// QueryHandler handles GraphQL requests
func (s *Server) QueryHandler(w http.ResponseWriter, req *http.Request) {
	if IsDev() {
		// In order to have GraphiQL on a different domain
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if req.Method == "OPTIONS" {
			return
		}
	}

	handler := &relay.Handler{Schema: s.schema}
	handler.ServeHTTP(w, req)
}
