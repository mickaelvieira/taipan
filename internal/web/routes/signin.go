package routes

import (
	"net/http"

	"github.com/mickaelvieira/taipan/internal/domain/errors"
	"github.com/mickaelvieira/taipan/internal/logger"
	"github.com/mickaelvieira/taipan/internal/repository"
	"github.com/mickaelvieira/taipan/internal/usecase"
	"github.com/mickaelvieira/taipan/internal/web"
	"github.com/mickaelvieira/taipan/internal/web/auth"

	"github.com/labstack/echo/v4"
)

func signin(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := web.GetSession(c)

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}

	u, err := usecase.Signin(ctx, r, creds.Email, creds.Password)
	if err != nil {
		if err, ok := err.(errors.DomainError); ok {
			if err.HasReason() {
				logger.Debug(err.Reason())
			}
			return c.JSON(http.StatusUnauthorized, jsonError(err.Domain()))
		}
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
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
