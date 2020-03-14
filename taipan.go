package main

import (
	"os"

	"github.com/mickaelvieira/taipan/internal/cmd"
	"github.com/mickaelvieira/taipan/internal/config"

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
	app.Run(os.Args)
}
