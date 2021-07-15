package main

import (
	"fmt"

	"github.com/Gate2Up/rabbitmq-go/subscriber"
)

// create handler function
func handler(data []byte) error {
	fmt.Println(data)
	return nil
}

// create instance of the subscriber
var Subscriber = subscriber.NewSubscriber("SUBSCRIBER_NAME", nil, handler)
