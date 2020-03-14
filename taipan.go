package main

import (
	"os"

	"github/mickaelvieira/taipan/internal/cmd"
	"github/mickaelvieira/taipan/internal/config"

	"github.com/urfave/cli"
)

func main() {
	config.LoadEnvironment("./")
	app := cli.NewApp()
	app.Name = "Taipan"
	app.Usage = "Self-hosted News Aggregator"
	app.Version = os.Getenv("APP_VERSION")
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Syndication,
		cmd.Documents,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
