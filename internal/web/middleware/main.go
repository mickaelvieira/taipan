package middleware

import (
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/clientid"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web"

	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func isSecuredArea(path string) bool {
	return path == "/graphql"
}

func isWebSocket(req *http.Request) bool {
	return req.Header.Get("Upgrade") == "websocket"
}

// Firewall middleware
func Firewall(r *repository.Repositories) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			sess := web.GetSession(c)
			authorized := false

			// @TODO I need to investigate a better way of handling websocket authentication
			if !isSecuredArea(c.Request().RequestURI) {
				authorized = true
			} else {
				ctx := req.Context()
				userID, ok := sess.Values["user_id"].(string)
				if ok {
					user, err := r.Users.GetByID(ctx, userID)
					if err == nil && user != nil {
						req = req.WithContext(auth.NewContext(ctx, user))
						c.SetRequest(req)
						authorized = true
					}
				}
			}

			if !authorized && !isWebSocket(c.Request()) {
				return c.JSON(http.StatusUnauthorized, struct{}{})
			}
			return next(c)
		}
	}
}

// Session get the session middleware
func Session() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte("secret")))
}

// CORS get the CORS middleware
func CORS() echo.MiddlewareFunc {
	return mw.CORSWithConfig(mw.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			os.Getenv("APP_CLIENT_ID_HEADER"),
		},
	})
}

// ClientID garbs the client ID from the HTTP Header and stick it in the context
func ClientID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			clientID := req.Header.Get(os.Getenv("APP_CLIENT_ID_HEADER"))

			ctx = clientid.NewContext(ctx, clientID)

			c.SetRequest(req.WithContext(ctx))

			return next(c)
		}
	}
}
