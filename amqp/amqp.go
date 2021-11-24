package amqp

import (
	"fmt"
	"log"

	"github.com/Gate2Up/rabbitmq-go/subscriber"
	"github.com/streadway/amqp"
)

type Config struct {
	ServiceName string
	Host        string
	Port        int
	User        string
	Password    string
}

type AmqpClient struct {
	Connection  *amqp.Connection
	ServiceName string
}

func SetUri(config Config) string {
	return fmt.Sprintf(`amqp://%s:%s@%s:%d`, config.User, config.Password, config.Host, config.Port)
}

// Open Connection to AMQP Server
func NewClient(config Config) (*AmqpClient, error) {
	uri := SetUri(config)
	amqpConn, err := amqp.Dial(uri)

	if err != nil {
		return nil, err
	}

	return &AmqpClient{Connection: amqpConn, ServiceName: config.ServiceName}, nil
}

type Publisher interface {
	Build(client *AmqpClient)
}

func (a *AmqpClient) AddPublisher(publisher Publisher) {
	publisher.Build(a)
}

func (a *AmqpClient) AddSubscriber(subscriber *subscriber.SubscriberConfig) {
	channel, err := a.Connection.Channel()
	if err != nil {
		log.Println(err.Error())
		return
	}

	subscriberName := fmt.Sprintf(`%s:%s`, a.ServiceName, subscriber.TopicName)

	_, err = channel.QueueDeclare(
		subscriberName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(`Failed create queue:`, err.Error())
		return
	}

	err = channel.QueueBind(subscriberName, "*", subscriber.TopicName, false, nil)

	if err != nil {
		log.Println(`Failed to binding: `, err.Error())
		return
	}

	msgs, err := channel.Consume(
		subscriberName,
		"*",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(`Failed to consume: `, err.Error())
		return
	}

	forever := make(chan bool)
	log.Println(`Listening Subcriber:`, subscriberName)
	go func() {
		for d := range msgs {
			subscriber.Handler(d.Body)
		}
	}()

	<-forever
}
