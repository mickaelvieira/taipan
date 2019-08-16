package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web"

	"github.com/labstack/echo/v4"
)

type apiError struct {
	Error string `json:"error"`
}

func signin(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := web.GetSession(c)

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusInternalServerError, &apiError{Error: auth.ErrorServerIssue.Error()})
	}

	u, err := usecase.Signin(ctx, r, creds.Email, creds.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &apiError{Error: err.Error()})
	}

	// open a new session for this user
	sess.Values["user_id"] = u.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, &u)
}

// Signin route
func Signin(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return signin(c, r)
	}
}