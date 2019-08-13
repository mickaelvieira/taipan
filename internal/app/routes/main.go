package routes

import (
	"net/http"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/app/paths"
	"github/mickaelvieira/taipan/internal/assets"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context, a assets.Assets) error {
	data := struct {
		Assets   assets.Assets
		BasePath string
		CDN      string
	}{
		Assets:   a,
		BasePath: paths.GetBasePath(app.UseFileServer()),
		CDN:      paths.GetCDNBaseURL(),
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// Index route
// we want to load assets files when we start the application
// not at every request. So we pass the list of assets files
// to the index route
func Index(a assets.Assets) func(echo.Context) error {
	return func(c echo.Context) error {
		return index(c, a)
	}
}
