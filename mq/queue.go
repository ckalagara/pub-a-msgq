package mq

import (
	"errors"
	"fmt"
	"sync"
)

type Queue interface {
	Subscribe(clientID string, topic string) (chan Message, error)
	Unsubscribe(clientID string, topic string) error
	Publish(topic string, message Message) error
	Shutdown()
}

func NewQueueWithChannelImpl(topics []string) Queue {
	q := &qChannelImpl{
		topics: topics,
	}

	for _, t := range q.topics {
		k := Topic{
			name: t,
			core: make(map[string]chan Message),
			lock: new(sync.Mutex),
		}
		q.core.Store(t, k)
	}

	return q
}

type qChannelImpl struct {
	core   sync.Map
	topics []string
}

type Topic struct {
	name string
	core map[string]chan Message
	lock *sync.Mutex
}

func (t *Topic) Shutdown() {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, ch := range t.core {
		close(ch)
	}
}

func (q *qChannelImpl) Subscribe(clientID string, topic string) (chan Message, error) {
	val, ok := q.core.Load(topic)
	if !ok {
		return nil, errors.New("Topic does not exist.")
	}

	topicCore, ok := val.(Topic)
	if !ok {
		return nil, errors.New("Topic does not implement Topic.")
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()
	ch := make(chan Message)
	topicCore.core[clientID] = ch
	return ch, nil
}

func (q *qChannelImpl) Unsubscribe(clientID string, topic string) error {
	val, ok := q.core.Load(topic)
	if !ok {
		return errors.New("Topic does not exist.")
	}

	topicCore, ok := val.(Topic)
	if !ok {
		return errors.New("Topic does not implement Topic.")
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()

	close(topicCore.core[clientID])

	delete(topicCore.core, clientID)
	return nil
}

func (q *qChannelImpl) Publish(topic string, message Message) error {
	val, ok := q.core.Load(topic)
	if !ok {
		return errors.New("Topic does not exist.")
	}

	topicCore, ok := val.(Topic)
	if !ok {
		return errors.New("Topic does not implement Topic.")
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()
	for client, ch := range topicCore.core {
		fmt.Printf("Publishing message to topic %s for %s client\n", topic, client)
		ch <- message
	}
	return nil
}

func (q *qChannelImpl) Shutdown() {
	q.core.Range(func(key, value interface{}) bool {
		t, ok := value.(Topic)
		if !ok {
			return false
		}
		t.Shutdown()
		return true
	})
}

type Message = map[string]interface{}
