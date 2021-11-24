package publisher

import (
	"fmt"
	"log"

	"github.com/Gate2Up/rabbitmq-go/amqp"
	amqpLegacy "github.com/streadway/amqp"
)

type PublisherConfig struct {
	TopicName string
	Schema    interface{}
	Client    *amqp.AmqpClient
}

type schemaType interface{}

func NewPublisher(topicName string, schema schemaType) *PublisherConfig {
	publisherConfig := PublisherConfig{
		TopicName: topicName,
		Schema:    schema,
		Client:    nil,
	}

	return &publisherConfig
}

func (p *PublisherConfig) Build(client *amqp.AmqpClient) {

	if client == nil {
		log.Fatalln(`amqp client is nil`)
	}

	p.Client = client

	channel, err := client.Connection.Channel()
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = channel.ExchangeDeclare(
		p.TopicName,
		amqpLegacy.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(`Create exchange failed: `, err.Error())
	}

	log.Println(fmt.Sprintf(`Exchange: %s created`, p.TopicName))
}

func (p *PublisherConfig) Publish(data []byte) (bool, error) {
	channel, err := p.Client.Connection.Channel()
	if err != nil {
		return false, err
	}

	content := amqpLegacy.Publishing{
		ContentType: "text/plain",
		Body:        data,
	}

	err = channel.Publish(p.TopicName, "*", true, true, content)

	if err != nil {
		return false, err
	}

	return true, nil
}
