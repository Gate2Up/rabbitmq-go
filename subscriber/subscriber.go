package subscriber

type handlerFunc = func([]byte) error

type SubscriberConfig struct {
	TopicName string
	Schema    interface{}
	Handler   handlerFunc
}

func NewSubscriber(topicName string, schema interface{}, handler handlerFunc) *SubscriberConfig {
	subscriberConfig := SubscriberConfig{
		TopicName: topicName,
		Schema:    schema,
		Handler:   handler,
	}

	return &subscriberConfig
}
