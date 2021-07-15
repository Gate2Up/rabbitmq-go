package amqp_test

import (
	"testing"

	"github.com/Gate2Up/rabbitmq-go/amqp"
	"github.com/Gate2Up/rabbitmq-go/publisher"
	"github.com/Gate2Up/rabbitmq-go/subscriber"
	amqpDriver "github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

var config = amqp.Config{
	Host:        "localhost",
	Port:        5672,
	User:        "guest",
	Password:    "guest",
	ServiceName: "TEST",
}

// func amqpConn() *amqpDriver.Connection {
// 	uri := amqp.SetUri(config)
// 	connection, _ := amqpDriver.Dial(uri)

// 	return connection
// }

func TestNewClient(t *testing.T) {

	conn, err := amqp.NewClient(config)

	assert.Equal(t, err, nil)
	assert.NotNil(t, conn)
}

func TestAddPublisher(t *testing.T) {

	publisherConfig := publisher.NewPublisher("TEST_TOPIC", nil)
	conn, _ := amqp.NewClient(config)

	// if no error this case is passed - void
	conn.AddPublisher(publisherConfig)
}

func TestAddSubscriber(t *testing.T) {
	dataDemo := []byte("Hello World")

	forever := make(chan bool)
	subscriberConfig := subscriber.SubscriberConfig{
		TopicName: "TEST_TOPIC",
		Schema:    nil,
		Handler: func(data []byte) error {
			assert.Equal(t, data, dataDemo)
			forever <- true
			return nil
		},
	}

	conn, _ := amqp.NewClient(config)

	// if no error this case is passed
	go conn.AddSubscriber(subscriberConfig)

	// send data
	channel, err := conn.Connection.Channel()
	if err != nil {
		t.Fatal("Failed to open channel")
	}

	content := amqpDriver.Publishing{
		ContentType: "text/plain",
		Body:        dataDemo,
	}
	channel.Publish(subscriberConfig.TopicName, "*", false, false, content)

	<-forever
}
