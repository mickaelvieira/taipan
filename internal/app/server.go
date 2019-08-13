package app

import (
	"context"
	"encoding/json"
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/repository"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-session/session"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"golang.org/x/crypto/bcrypt"
)

// Server is the main application
type Server struct {
	templates    *template.Template
	schema       *graphql.Schema
	assets       assets.Assets
	Repositories *repository.Repositories
}

type apiError struct {
	Error string `json:"error"`
}

type tmplData struct {
	Assets   assets.Assets
	BasePath string
	CDN      string
}

type emptyJSON struct{}

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
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	var c *auth.Credentials
	if err = json.Unmarshal(body, &c); err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	// can we find the user?
	user, err := s.Repositories.Users.GetByUsername(ctx, c.Username)
	if err != nil {
		writeJSONError(w, http.StatusUnauthorized, auth.ErrorInvalidCreds)
		return
	}

	// can we find the user's password?
	password, err := s.Repositories.Users.GetPassword(ctx, user.ID)
	if err != nil {
		writeJSONError(w, http.StatusUnauthorized, auth.ErrorInvalidCreds)
		return
	}

	// do the password match?
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(c.Password)); err != nil {
		writeJSONError(w, http.StatusUnauthorized, auth.ErrorInvalidCreds)
		return
	}

	// open a new session for this user
	store.Set("user_id", user.ID)
	if err = store.Save(); err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	// send the response
	j, err := json.Marshal(user)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

// SignoutHandler sign out of the application
func (s *Server) SignoutHandler(w http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), w, req)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	// delete user session
	store.Delete("user_id")
	if err = store.Save(); err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	// send the response
	var r emptyJSON
	j, err := json.Marshal(r)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, auth.ErrorServerIssue)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
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
