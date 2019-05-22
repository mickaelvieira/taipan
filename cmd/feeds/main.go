package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/rmq"
	"github/mickaelvieira/taipan/internal/usecase"
)

func main() {
	fmt.Println("Starting feeds worker")
	app.LoadEnvironment()
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

	for {
		fmt.Println("loop")
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)

			feeds, err := repositories.Feeds.GetOutdatedFeeds(ctx)
			if err != nil {
				log.Fatal(err)
			}

			for _, feed := range feeds {
				var entries []string
				entries, err = usecase.ParseFeed(ctx, feed, repositories)
				if err != nil {
					// We just log the parsing errors for now
					log.Println(err)
				}

				for _, entry := range entries {
					e := client.PublishDocument(entry)
					if e != nil {
						log.Println(e)
					}
				}
			}
		}
	}
}
