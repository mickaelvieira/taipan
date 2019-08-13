package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/repository"

	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

type apiError struct {
	Error string `json:"error"`
}

func signin(c echo.Context, r *repository.Repositories) error {
	req := c.Request()
	ctx := req.Context()
	sess := app.GetSession(c)

	creds := new(auth.Credentials)
	if err := c.Bind(creds); err != nil {
		return c.JSON(http.StatusInternalServerError, &apiError{Error: auth.ErrorServerIssue.Error()})
	}

	// can we find the user?
	user, err := r.Users.GetByUsername(ctx, creds.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &apiError{Error: auth.ErrorInvalidCreds.Error()})
	}

	// can we find the user's password?
	password, err := r.Users.GetPassword(ctx, user.ID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &apiError{Error: auth.ErrorInvalidCreds.Error()})
	}

	// do the password match?
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(creds.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, &apiError{Error: auth.ErrorInvalidCreds.Error()})
	}

	// open a new session for this user
	sess.Values["user_id"] = user.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusInternalServerError, &user)
}

// Signin route
func Signin(r *repository.Repositories) func(c echo.Context) error {
	return func(c echo.Context) error {
		return signin(c, r)
	}
}
