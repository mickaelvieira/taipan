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

func resetPassword(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()

	args := new(auth.ResetPassword)
	if err := c.Bind(args); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}

	if err := usecase.ResetPassword(ctx, r, args.Token, args.Password); err != nil {
		if err, ok := err.(errors.DomainError); ok {
			if err.HasReason() {
				logger.Debug(err.Reason())
			}
			return c.JSON(http.StatusBadRequest, jsonError(err.Domain()))
		}
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}
	return c.JSON(http.StatusOK, struct{}{})
}

// ResetPassword route
func ResetPassword(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return resetPassword(c, r)
	}
}
