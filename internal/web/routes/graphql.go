package routes

import (
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/graphql/generated"
	"github/mickaelvieira/taipan/internal/web/graphql/resolvers"

	"github.com/labstack/echo/v4"

	// gql "github.com/graph-gophers/graphql-go"
	//

	"github.com/99designs/gqlgen/handler"
	// "github.com/graph-gophers/graphql-go/relay"
)

// // QueryHandler handles GraphQL requests
func graphql(c echo.Context, r *repository.Repositories) error {
	// handler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	// handler := &relay.Handler{Schema: s}

	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.RootResolver{}}))
	h(c.Response(), c.Request())

	return nil
}

// GraphQL routes
func GraphQL(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return graphql(r)
	}
}
