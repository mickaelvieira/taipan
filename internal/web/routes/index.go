package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/web"
	"github/mickaelvieira/taipan/internal/web/assets"
	"github/mickaelvieira/taipan/internal/web/paths"

	"github.com/labstack/echo/v4"
)

type apiError struct {
	Error string `json:"error"`
}

func jsonError(e error) *apiError {
	return &apiError{
		Error: e.Error(),
	}
}

func index(c echo.Context, a *assets.Assets) error {
	data := struct {
		Assets   *assets.Assets
		BasePath string
		CDN      string
	}{
		Assets:   a,
		BasePath: paths.GetBasePath(web.UseFileServer()),
		CDN:      paths.GetCDNBaseURL(),
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// Index route
// we want to load assets files when we start the application
// not at every request. So we pass the list of assets files
// to the index route
func Index(a *assets.Assets) func(echo.Context) error {
	return func(c echo.Context) error {
		return index(c, a)
	}
}
