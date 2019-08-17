package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web"

	"github.com/labstack/echo/v4"
)

func signup(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := web.GetSession(c)

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusInternalServerError, &apiError{Error: auth.ErrorServerIssue.Error()})
	}

	// Let's the user's account
	u, err := usecase.Signup(ctx, r, creds.Email, creds.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &apiError{Error: err.Error()})
	}

	// all goo, open a new session for this user
	sess.Values["user_id"] = u.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, struct{}{})
}

// Signup route
func Signup(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return signup(c, r)
	}
}
