package routes

import (
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/labstack/echo/v4"

	gql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
)

// // QueryHandler handles GraphQL requests
func graphql(c echo.Context, s *gql.Schema, r *repository.Repositories) error {
	// handler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	handler := &relay.Handler{Schema: s}
	handler.ServeHTTP(c.Response(), c.Request())

	return nil
}

// GraphQL routes
func GraphQL(s *gql.Schema, r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return graphql(c, s, r)
	}
}
