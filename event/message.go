package event

type Message struct {
	TransID   int64
	ServiceID int64
	Topic     string
	Data      interface{}
}
