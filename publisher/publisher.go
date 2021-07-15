package publisher

type PublisherConfig struct {
	TopicName string
	Schema    interface{}
}

func NewPublisher(topicName string, schema interface{}) *PublisherConfig {
	publisherConfig := PublisherConfig{
		TopicName: topicName,
		Schema:    schema,
	}

	return &publisherConfig
}
