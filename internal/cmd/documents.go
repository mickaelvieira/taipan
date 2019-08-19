package cmd

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/messages"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/rmq"
	"github/mickaelvieira/taipan/internal/usecase"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli"
)

// Documents command
var Documents = cli.Command{
	Name:        "documents",
	Usage:       "Start the web document worker",
	Description: ``,
	Action:      runDocumentsWorker,
}

func runDocumentsWorker(c *cli.Context) {
	l := log.New("documents")
	logger.Init(l, os.Getenv("APP_LOG_LEVEL"))

	ctx, cancel := context.WithCancel(context.Background())
	repositories := repository.GetRepositories()

	client, err := rmq.New()
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	onStop := func() {
		// Cancel context
		cancel()
		// Close RabbitMQ connection
		client.Close()
	}
	Signal(onStop)
	defer onStop()

	queue := client.GetDocumentQueue()

	msgs, err := client.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	go func() {
		for d := range msgs {
			dm := &messages.Document{}
			if err := proto.Unmarshal(d.Body, dm); err != nil {
				logger.Error(fmt.Sprintf("Failed to parse document message: %s", err))
				continue
			}
			logger.Warn(fmt.Sprintf("Received a message: %s", dm))

			var u *url.URL
			u, err = url.FromRawURL(dm.Url)
			if err != nil {
				logger.Error(err)
				continue
			}

			d, err := usecase.Document(ctx, repositories, u, false)
			if err != nil {
				logger.Error(err)
				continue
			}

			if usecase.AddDocumentToNewsFeeds(ctx, repositories, dm.SourceId, d.ID); err != nil {
				logger.Error(err)
				continue
			}
		}
	}()

	logger.Warn(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
