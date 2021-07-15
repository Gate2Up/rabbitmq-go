package main

import "github.com/Gate2Up/rabbitmq-go/publisher"

// define topic name
var topicName = "EXAMPLE"

// create instance of publisher
var Publisher = publisher.NewPublisher(topicName, nil)
