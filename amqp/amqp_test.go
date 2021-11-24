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

	publisher := publisher.NewPublisher("TEST_TOPIC", nil)
	conn, _ := amqp.NewClient(config)

	// if no error this case is passed - void
	conn.AddPublisher(publisher)
	status, err := publisher.Publish([]byte(`Hello World`))

	assert.Equal(t, status, true)
	assert.Equal(t, err, nil)
}

func TestAddSubscriber(t *testing.T) {
	dataDemo := []byte("Hello World")

	forever := make(chan bool)

	subscriberConfig := subscriber.NewSubscriber("TEST_TOPIC", nil, func(data []byte) error {
		assert.Equal(t, data, dataDemo)
		forever <- true
		return nil
	})

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
