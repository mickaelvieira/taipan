package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/domain/errors"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web"
	"github/mickaelvieira/taipan/internal/web/auth"

	"github.com/labstack/echo/v4"
)

func signup(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := web.GetSession(c)

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}

	// Let's the user's account
	u, err := usecase.Signup(ctx, r, creds.Email, creds.Password)
	if err != nil {
		if err, ok := err.(errors.DomainError); ok {
			if err.HasReason() {
				logger.Debug(err.Reason())
			}
			return c.JSON(http.StatusBadRequest, jsonError(err.Domain()))
		}
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}

	// all good, open a new session for this user
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
