package routes

import (
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web"
	"net/http"

	"github.com/labstack/echo/v4"

	gql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
)

// // QueryHandler handles GraphQL requests
func graphql(c echo.Context, s *gql.Schema, r *repository.Repositories) error {
	sess := web.GetSession(c)

	userID, ok := sess.Values["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, struct{}{})
	}

	req := c.Request()
	ctx := req.Context()
	user, err := r.Users.GetByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, struct{}{})
	}

	handler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	handler.ServeHTTP(c.Response(), req.WithContext(auth.NewContext(ctx, user)))

	return nil
}

// GraphQL routes
func GraphQL(s *gql.Schema, r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return graphql(c, s, r)
	}
}
