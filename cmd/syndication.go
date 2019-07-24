package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/mickaelvieira/taipan/internal/app"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/logger"
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

type fetchResult struct {
	source *syndication.Source
	result *http.Result
}

func runSyndicationWorker(c *cli.Context) {
	logger.Warn("Starting syndication worker")
	ctx, cancel := context.WithCancel(context.Background())
	repos := repository.GetRepositories()

	// @TODO Check out how RMQ handle context
	logger.Warn("Creating RabbitMQ client")
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
		logger.Warn(fmt.Sprintf("Get outdated sources with frequency [%s]", http.Hourly))
		sources, err := repos.Syndication.GetOutdatedSources(ctx, http.Hourly)
		if err != nil {
			log.Fatal(err)
		}

		if len(sources) == 0 {
			d := 60 * time.Second
			logger.Warn(fmt.Sprintf("Waiting for sources to be outdated, sleep for [%ds]", 60))
			time.Sleep(d)
			continue
		}

		cr := make(chan *fetchResult)
		cu := make(chan syndication.Queue)
		cf := make(chan bool)

		// We store web syndication sources in a queue
		// which will be empty when all sources will be parsed
		queue := syndication.QueueSources(sources)
		// The mixer contains a queue for each sources
		// containing a list of messages
		mixer := syndication.MakeMixer(len(queue))

		logger.Info(fmt.Sprintf("Process queue of [%d] entities", len(queue)))

		go func(cr chan<- *fetchResult, q syndication.QueueSources) {
			for len(q) > 0 {
				go func(s *syndication.Source) {
					r, err := usecase.FetchResource(ctx, repos, s.URL)
					// we might have a SQL error
					// but we always get an HTTP result
					if err != nil {
						log.Fatalln(err)
					}
					cr <- &fetchResult{source: s, result: r}
				}(q.Shift())
			}
		}(cr, queue)

		go func(cr <-chan *fetchResult, cu chan<- syndication.Queue) {
			// @TODO it might be better to use for r := range cr {} here
			for {
				select {
				case r := <-cr:
					urls, err := usecase.ParseSyndicationSource(ctx, repos, r.result, r.source)
					if err != nil {
						logger.Error(fmt.Sprintf("Syndication Parser: URL %s", r.source.URL))
						logger.Error(err)
					}
					cu <- syndication.MakeQueue(urls, r.source.ID)
				}
			}
		}(cr, cu)

		go func(c <-chan syndication.Queue, f chan<- bool, m *syndication.Mixer) {
			for {
				select {
				case u := <-c:
					m.Push(u)
					if m.IsFull() {
						f <- true
						return
					}
				}
			}
		}(cu, cf, mixer)

		<-cf

		go func(s *syndication.Mixer) {
			for _, u := range s.Mixup() {
				logger.Info(fmt.Sprintf("Publishing '%s'", u))
				e := client.PublishDocument(u)
				if e != nil {
					logger.Error(e)
				}
			}
		}(mixer)
	}
}
