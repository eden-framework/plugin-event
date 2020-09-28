package event

type MessageRunner func(message Message) (Message, error)

type MessageHandler struct {
	Topic  string
	Runner MessageRunner
}

type MessageDriver interface {
	Subscribe(topic string, handler MessageRunner)
	Unsubscribe(topic string, handler MessageRunner)
	Publish(message Message) ([]Message, error)
	AsyncPublish(message Message) (<-chan Message, error)
}
