package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/labstack/echo/v4"
)

func forgotPassword(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusInternalServerError, &apiError{Error: auth.ErrorServerIssue.Error()})
	}

	if err := usecase.ForgotPassword(ctx, r, creds.Email); err != nil {
		if err == usecase.ErrInvalidEmail {
			return c.JSON(http.StatusBadRequest, &apiError{Error: err.Error()})
		}
		logger.Warn(err)
	}
	return c.JSON(http.StatusOK, struct{}{})
}

// ForgoPassword route
func ForgoPassword(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return forgotPassword(c, r)
	}
}
