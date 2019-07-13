package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
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

	for {
		fmt.Printf("Get outdated sources with frequency [%s]", http.Hourly)
		sources, err := repositories.Syndication.GetOutdatedSources(ctx, http.Hourly)
		if err != nil {
			log.Fatal(err)
		}

		if len(sources) == 0 {
			d := 60 * time.Second
			fmt.Printf("No outdated sources, sleep for [%ds]\n", 60)
			time.Sleep(d)
			continue
		}

		// We store web syndication sources in a queue
		// which will be empty when all sources will be parsed
		queue := syndication.QueueSources(sources)
		// The mixer contains a queue for each sources
		// containing a list of URLs
		mixer := syndication.MakeMixer(len(queue))

		fmt.Printf("Process queue of [%d] entities\n", len(queue))

		for len(queue) > 0 {
			// We get the source from the front of the queue
			source := queue.Shift()
			// We then parse the source
			urls, err := usecase.ParseSyndicationSource(ctx, repositories, source)
			if err != nil {
				// We just log the parsing errors for now
				// urls slice will be empty anyway so we can keep going
				log.Printf("Syndication Parser: URL %s\n", source.URL)
				log.Println(err)
			}
			// push the list of URLs into their respective queue
			mixer.Push(urls)
		}

		fmt.Printf("Queue is empty: [%d] entities\n", len(queue))

		go func(s *syndication.Mixer) {
			// mix up the URLs and send publish them
			for _, u := range s.Mixup() {
				fmt.Printf("Publishing '%s'\n", u)
				e := client.PublishDocument(u)
				if e != nil {
					log.Println(e)
				}
			}
		}(mixer)
	}
}
