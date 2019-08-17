package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/web"

	"github.com/labstack/echo/v4"
)

func signout(c echo.Context) error {
	sess := web.GetSession(c)

	// delete user session
	sess.Values["user_id"] = nil
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, struct{}{})
}

// Signout route
func Signout() func(c echo.Context) error {
	return func(c echo.Context) error {
		return signout(c)
	}
}
