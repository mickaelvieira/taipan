package cmd

import (
	"os"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/app/middleware"
	"github/mickaelvieira/taipan/internal/app/paths"
	"github/mickaelvieira/taipan/internal/app/routes"
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/graphql"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/templates"

	"github.com/labstack/echo/v4"
	"github.com/urfave/cli"
)

// Web server command
var Web = cli.Command{
	Name:        "web",
	Usage:       "Start the web server",
	Description: ``,
	Action:      runWeb,
}

func runWeb(c *cli.Context) {
	a := assets.LoadAssetsDefinition(paths.GetScriptsDir(), app.UseFileServer())
	t := templates.NewRenderer(paths.GetTemplatesDir())
	r := repository.GetRepositories()
	s := graphql.LoadAndParseSchema(paths.GetGraphQLSchema(), r)

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.ClientID())
	e.Use(middleware.Session())
	e.Use(middleware.Firewall())

	if app.IsDev() {
		e.Debug = true
		e.Use(middleware.CORS())
	}
	if app.UseFileServer() {
		e.Static(paths.GetBasePath(app.UseFileServer()), paths.GetStaticDir())
	}

	e.POST("/login", routes.Signin(r))
	e.POST("/logout", routes.Signout())
	e.GET("/graphql", routes.GraphQL(s, r))
	e.POST("/graphql", routes.GraphQL(s, r))
	e.GET("/*", routes.Index(a))

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
