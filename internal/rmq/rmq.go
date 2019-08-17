package rmq

import (
	"github/mickaelvieira/taipan/internal/domain/messages"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

// getDSN returns the RabbitMQ DSN
func getDSN() string {
	return "amqp://" + os.Getenv("APP_RMQ_USER") + ":" + os.Getenv("APP_RMQ_PWD") + "@" + os.Getenv("APP_RMQ_HOST") + ":" + os.Getenv("APP_RMQ_PORT") + "/"
}

// AMQPClient client to communicate with RabbitMQ
type AMQPClient struct {
	conn    *amqp.Connection
	Channel *amqp.Channel
	queue   *amqp.Queue
}

// Close closes the connections
func (c *AMQPClient) Close() {
	if c.Channel != nil {
		c.Channel.Close()
	}
	c.conn.Close()
}

// GetDocumentQueue creates the document queue
func (c *AMQPClient) GetDocumentQueue() *amqp.Queue {
	queue, err := c.Channel.QueueDeclare(
		os.Getenv("APP_RMQ_DOCUMENTS_QUEUE_NAME"),
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatal("Cant create queue")
	}

	return &queue
}

// PublishDocument publishes a document URL to the channel
func (c *AMQPClient) PublishDocument(m *messages.Document) error {
	if c.queue == nil {
		c.queue = c.GetDocumentQueue()
	}

	return c.Channel.Publish(
		"",           // exchange
		c.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		c.getDocumentMessage(m),
	)
}

func (c *AMQPClient) getDocumentMessage(m *messages.Document) amqp.Publishing {
	body, err := proto.Marshal(m)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}
}

// New creates a AMQPClient
func New() (*AMQPClient, error) {
	conn, err := amqp.Dial(getDSN())
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	client := AMQPClient{
		conn:    conn,
		Channel: channel,
	}

	return &client, nil
}
