package app

import (
	"context"
	"encoding/json"
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/repository"
	"html/template"
	"net/http"
	"os"

	"github.com/go-session/session"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
)

// Server is the main application
type Server struct {
	templates    *template.Template
	schema       *graphql.Schema
	assets       *assets.Assets
	Repositories *repository.Repositories
}

type tmplData struct {
	Assets   *assets.Assets
	BasePath string
	CDN      string
}

// IndexHandler is the method to handle / route
func (s *Server) IndexHandler(w http.ResponseWriter, req *http.Request) {
	err := s.templates.ExecuteTemplate(w, "index.html", tmplData{
		Assets:   s.assets,
		BasePath: assets.GetBasePath(UseFileServer()),
		CDN:      assets.MakeCDNBaseURL(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SigninHandler sign into the application
func (s *Server) SigninHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	store, err := session.Start(ctx, w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := s.Repositories.Users.GetByID(ctx, "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} else {
		store.Set("user_id", user.ID)
		err = store.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
			js, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}
}

// SignoutHandler sign out of the application
func (s *Server) SignoutHandler(w http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), w, req)
	if err != nil {
		return
	}

	store.Delete("user_id")
	err = store.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// QueryHandler handles GraphQL requests
func (s *Server) QueryHandler(w http.ResponseWriter, req *http.Request) {
	if isDev() {
		// In order to have GraphiQL on a different domain
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"+", "+os.Getenv("APP_CLIENT_ID_HEADER"))
		if req.Method == "OPTIONS" {
			return
		}
	}

	handler := graphqlws.NewHandlerFunc(s.schema, &relay.Handler{Schema: s.schema})
	handler.ServeHTTP(w, req)
}
