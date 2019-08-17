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

	"github.com/labstack/gommon/log"

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
	l, ok := e.Logger.(*log.Logger)
	if !ok {
		panic("Cannot init logger")
	}
	logger.Init(l, os.Getenv("APP_LOG_LEVEL"))

	e.Renderer = t
	e.Use(middleware.ClientID())
	e.Use(middleware.Session())
	e.Use(middleware.Firewall(r))

	if web.IsDev() {
		e.Debug = true
		e.Use(middleware.CORS())
	}
	if web.UseFileServer() {
		e.Static(paths.GetBasePath(web.UseFileServer()), paths.GetStaticDir())
	}

	index := routes.Index(a)
	api := routes.GraphQL(s, r)

	e.POST("/signout", routes.Signout())

	e.POST("/signin", routes.Signin(r))
	e.GET("/signin", index)

	e.POST("/join", routes.Signup(r))
	e.GET("/join", index)

	e.POST("/forgot-password", routes.ForgoPassword(r))
	e.GET("/forgot-password", index)

	e.POST("/reset-password", routes.ResetPassword(r))
	e.GET("/reset-password", index)

	e.GET("/graphql", api)
	e.POST("/graphql", api)

	e.GET("/*", index)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
