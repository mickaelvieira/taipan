package main

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/rmq"
	"github/mickaelvieira/taipan/internal/usecase"
	"log"

	"github/mickaelvieira/taipan/internal/app"

	"github.com/gogo/protobuf/proto"
)

func main() {
	app.LoadEnvironment()
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
	app.Signal(onStop)
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
			dm := &document.DocumentMessage{}
			if err := proto.Unmarshal(d.Body, dm); err != nil {
				log.Fatalln("Failed to parse document message:", err)
			}
			fmt.Printf("Received a message: %s\n", dm)

			var u *url.URL
			u, err = url.FromRawURL(dm.Url)
			if err == nil {
				_, err = usecase.Document(ctx, u, false, repositories)
			}

			if err != nil {
				log.Printf("Document Parser: URL %s\n", u)
				log.Println(err)
			}
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
