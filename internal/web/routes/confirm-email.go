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

func confirmEmail(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := web.GetSession(c)

	// The user could be logged in
	// In that case, we want to make sure the logged in user matches the token
	userID, ok := sess.Values["user_id"].(string)
	if !ok {
		userID = ""
	}

	args := new(auth.ConfirmEmail)
	if err := c.Bind(args); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, jsonError(web.ErrServer))
	}

	if err := usecase.ConfirmEmail(ctx, r, args.Token, userID); err != nil {
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

// ConfirmEmail route
func ConfirmEmail(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return confirmEmail(c, r)
	}
}
