package cmd

import (
	"github/mickaelvieira/taipan/internal/assets"
	"github/mickaelvieira/taipan/internal/graphql"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/templates"
	"github/mickaelvieira/taipan/internal/web"
	"github/mickaelvieira/taipan/internal/web/middleware"
	"github/mickaelvieira/taipan/internal/web/paths"
	"github/mickaelvieira/taipan/internal/web/routes"
	"os"

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
	a := assets.LoadAssetsDefinition(paths.GetScriptsDir(), web.UseFileServer())
	t := templates.NewRenderer(paths.GetTemplatesDir())
	r := repository.GetRepositories()
	s := graphql.LoadAndParseSchema(paths.GetGraphQLSchema(), r)

	e := echo.New()
	logger.Init(e, os.Getenv("APP_LOG_LEVEL"))

	e.Renderer = t
	e.Use(middleware.ClientID())
	e.Use(middleware.Session())
	e.Use(middleware.Firewall())

	if web.IsDev() {
		e.Debug = true
		e.Use(middleware.CORS())
	}
	if web.UseFileServer() {
		e.Static(paths.GetBasePath(web.UseFileServer()), paths.GetStaticDir())
	}

	e.POST("/signin", routes.Signin(r))
	e.POST("/signout", routes.Signout())
	e.POST("/signup", routes.Signup(r))
	e.GET("/graphql", routes.GraphQL(s, r))
	e.POST("/graphql", routes.GraphQL(s, r))
	e.GET("/*", routes.Index(a))

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
