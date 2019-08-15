package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/labstack/echo/v4"
)

func resetPassword(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()

	args := new(auth.ResetPassword)
	if err := c.Bind(args); err != nil {
		return c.JSON(http.StatusInternalServerError, &apiError{Error: auth.ErrorServerIssue.Error()})
	}

	if err := usecase.ResetPassword(ctx, r, args.Token, args.Password); err != nil {
		if err == usecase.ErrNoPassword || err == usecase.ErrWeakPassword || err == usecase.ErrInvalidResetToken {
			return c.JSON(http.StatusBadRequest, &apiError{Error: err.Error()})
		}
		logger.Warn(err)
	}
	return c.JSON(http.StatusOK, struct{}{})
}

// ResetPassword route
func ResetPassword(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return resetPassword(c, r)
	}
}
