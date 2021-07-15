package publisher

type PublisherConfig struct {
	TopicName string
	Schema    interface{}
}

type schemaType interface{}

func NewPublisher(topicName string, schema schemaType) *PublisherConfig {
	publisherConfig := PublisherConfig{
		TopicName: topicName,
		Schema:    schema,
	}

	return &publisherConfig
}
