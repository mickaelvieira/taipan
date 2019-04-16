package app

import (
	"context"
	"github/mickaelvieira/taipan/internal/assets"
	"html/template"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Server is the main application
type Server struct {
	templates *template.Template
	schema    *graphql.Schema
	assets    *assets.Assets
}

// IndexHandler is the method to handle / route
func (s *Server) IndexHandler(w http.ResponseWriter, req *http.Request) {
	err := s.templates.ExecuteTemplate(w, "index.html", s.assets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// QueryHandler handles GraphQL requests
func (s *Server) QueryHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	handler := &relay.Handler{Schema: s.schema}
	handler.ServeHTTP(w, req.WithContext(ctx))
}
