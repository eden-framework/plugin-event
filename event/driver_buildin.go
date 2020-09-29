package event

import (
	"errors"
	"fmt"
	"github.com/derekparker/trie"
	"github.com/profzone/data_structures/list"
	"reflect"
)

var _ interface {
	messageDriver
} = (*messageDriverMemory)(nil)

type messageDriverMemory struct {
	tree *trie.Trie
}

func (m *messageDriverMemory) Subscribe(topic string, handler MessageRunner) {
	var handlers *list.SyncedList
	if node, ok := m.tree.Find(topic); ok {
		handlers, _ = node.Meta().(*list.SyncedList)
		handlers.PushBack(handler)
	} else {
		handlers = list.NewSyncedList()
		handlers.PushBack(handler)
		m.tree.Add(topic, handlers)
	}
}

func (m *messageDriverMemory) Unsubscribe(topic string, handler MessageRunner) {
	var handlers *list.SyncedList
	if node, ok := m.tree.Find(topic); ok {
		handlers, _ = node.Meta().(*list.SyncedList)
	} else {
		return
	}

	if handlers.Len() == 0 {
		return
	}

	for e := range handlers.Iter() {
		if reflect.ValueOf(e.Value).Pointer() == reflect.ValueOf(handler).Pointer() {
			handlers.Remove(e)
		}
	}

	if handlers.Len() == 0 {
		m.tree.Remove(topic)
	} else {
		m.tree.Add(topic, handlers)
	}
}

func (m *messageDriverMemory) Publish(message Message) (result []Message, err error) {
	if message.Topic == "" {
		return nil, errors.New("topic must NOT be empty")
	}

	topics := m.tree.PrefixSearch(message.Topic)
	if len(topics) == 0 {
		return nil, errors.New(fmt.Sprintf("can NOT find any topics by topic(prefix) %s", message.Topic))
	}

	for _, topic := range topics {

		if node, ok := m.tree.Find(topic); ok {

			handlers, _ := node.Meta().(*list.SyncedList)
			for e := range handlers.Iter() {
				handler := e.Value.(MessageRunner)
				handlerResult, err := handler(message)
				if err != nil {
					return nil, err
				}

				result = append(result, handlerResult)
			}
		} else {
			return nil, errors.New(fmt.Sprintf("can NOT find any handler by topic %s", topic))
		}
	}

	return
}

func (m *messageDriverMemory) AsyncPublish(message Message) (<-chan Message, error) {
	if message.Topic == "" {
		return nil, errors.New("topic must NOT be empty")
	}

	topics := m.tree.PrefixSearch(message.Topic)
	if len(topics) == 0 {
		return nil, errors.New(fmt.Sprintf("can NOT find any topics by topic(prefix) %s", message.Topic))
	}

	replyChannel := make(chan Message)

	go func() {
		for _, topic := range topics {

			if node, ok := m.tree.Find(topic); ok {

				handlers, _ := node.Meta().(*list.SyncedList)
				for e := range handlers.Iter() {
					handler := e.Value.(MessageRunner)
					handlerResult, _ := handler(message)

					replyChannel <- handlerResult
				}
			}
		}
		close(replyChannel)
	}()

	return replyChannel, nil
}

func newMemoryMessageBus() *messageDriverMemory {
	return &messageDriverMemory{
		tree: trie.New(),
	}
}
