# RabbitMQ x Golang

RabbitMQ Library to using Event Driven Application based on Publish/Subscribe built with Go

## How to use:

1. First before use this library you should create new amqp client

  ```go
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
  ```

2. Setup Publisher
   
  ```go
    import "github.com/Gate2Up/rabbitmq-go/publisher"

    // define topic name
    var topicName = "EXAMPLE"

    // create instance of publisher
    var Publisher = publisher.NewPublisher(topicName, nil)

    // register publisher to amqp
	  client.AddPublisher(Publisher)
  
  ```
3. Setup Subscriber

  ```go
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

	// register subscriber to amqp
	client.AddSubscriber(Subscriber)

  ```



### References

- [Go client for AMQP 0.9.1]('https://github.com/streadway/amqp')
