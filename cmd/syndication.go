package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/rmq"
	"github/mickaelvieira/taipan/internal/usecase"

	"github.com/urfave/cli"
)

// Syndication command
var Syndication = cli.Command{
	Name:        "syndication",
	Usage:       "Start the web syndication worker",
	Description: ``,
	Action:      runSyndicationWorker,
}

func runSyndicationWorker(c *cli.Context) {
	fmt.Println("Starting syndication worker")
	ctx, cancel := context.WithCancel(context.Background())
	repositories := repository.GetRepositories()

	// @TODO Check out how RMQ handle context
	fmt.Println("Creating RabbitMQ client")
	client, err := rmq.New()
	if err != nil {
		log.Fatal(err)
	}

	onStop := func() {
		// Cancel context
		cancel()
		// Close RabbitMQ connection
		client.Close()
	}
	app.Signal(onStop)
	defer onStop()

	ticker := time.NewTicker(5 * time.Minute)

	for t := range ticker.C {
		fmt.Println("Tick at", t)

		sources, err := repositories.Syndication.GetOutdatedSources(ctx, http.Hourly)
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range sources {
			var urls []*url.URL
			urls, err = usecase.ParseSyndicationSource(ctx, repositories, s)
			if err != nil {
				log.Printf("Syndication Parser: URL %s\n", s.URL)
				log.Println(err) // We just log the parsing errors for now
			}
			for _, url := range urls {
				e := client.PublishDocument(url)
				if e != nil {
					log.Println(e)
				}
			}
		}
	}
}
