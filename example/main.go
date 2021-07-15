package main

import (
	"fmt"

	"github.com/Gate2Up/rabbitmq-go/amqp"
)

func main() {

	// set configuration
	config := amqp.Config{
		ServiceName: "EXAMPLE",
		Host:        "localhost",
		Port:        5672,
		User:        "guest",
		Password:    "guest",
	}

	// open connection to amqp server
	client, err := amqp.NewClient(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// register publisher to amqp
	client.AddPublisher(Publisher)

	// register subscriber to amqp
	client.AddSubscriber(Subscriber)

}
