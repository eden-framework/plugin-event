package event

import (
	"fmt"
	"github.com/derekparker/trie"
	"github.com/profzone/data_structures/list"
	"reflect"
	"testing"
)

var m = MessageDriverMemory{
	tree: trie.New(),
}

var testCase = map[string][]MessageRunner{
	"foo": {
		func(message Message) (Message, error) {
			return Message{
				Data: fmt.Sprintf("%s, this is foo handler1 result", message.Data),
			}, nil
		},
		func(message Message) (Message, error) {
			return Message{
				Data: fmt.Sprintf("%s, this is foo handler2 result", message.Data),
			}, nil
		},
	},
}

func TestSubscribe(t *testing.T) {

	for topic, handlers := range testCase {
		for _, handler := range handlers {
			m.Subscribe(topic, handler)
		}
	}

	for topic := range testCase {
		if node, ok := m.tree.Find(topic); ok {
			meta := node.Meta()
			if handlers, ok := meta.(*list.SyncedList); ok {
				for e := range handlers.Iter() {
					handler := e.Value.(MessageRunner)
					result, _ := handler(Message{
						Data: "TestSubscribe",
					})
					t.Log(result.Data)
				}
			} else {
				t.Errorf("Subscribe faild, expected *list.SyncedList, got %v", reflect.TypeOf(meta).String())
			}
		} else {
			t.Errorf("Subscribe faild, expected %v, got nil", topic)
		}
	}
}

func TestPublish(t *testing.T) {
	result, err := m.Publish(Message{
		Topic: "foo",
		Data:  "TestPublish",
	})

	if err != nil {
		t.Fatal(err)
		return
	}

	if len(result) != 2 {
		t.Fatalf("len(result) expected 2, got %d", len(result))
	}

	for _, message := range result {
		t.Log(message.Data)
	}
}

func TestAsyncPublish(t *testing.T) {
	msg := Message{
		Topic: "foo",
		Data:  "TestAsyncPublish",
	}

	result, err := m.AsyncPublish(msg)
	if err != nil {
		t.Fatal(err)
		return
	}

	for reply := range result {
		t.Log(reply.Data)
	}

}

func TestUnsubscribe(t *testing.T) {
	m.Unsubscribe("foo", testCase["foo"][0])

	if node, ok := m.tree.Find("foo"); ok {
		meta := node.Meta()
		if handlers, ok := meta.(*list.SyncedList); ok {
			length := handlers.Len()
			if length != 1 {
				t.Errorf("handlers.Len() expected 1, got %d", length)
			}
			for e := range handlers.Iter() {
				handler := e.Value.(MessageRunner)
				result, _ := handler(Message{
					Data: "TestUnsubscribe",
				})
				t.Log(result.Data)
			}
		} else {
			t.Errorf("Unsubscribe faild, expected *list.SyncedList, got %v", reflect.TypeOf(meta).String())
		}
	}
}
