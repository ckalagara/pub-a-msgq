package mq

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

const (
	debug = false
)

type Queue interface {
	Subscribe(clientID string, topic string) (chan Message, error)
	Unsubscribe(clientID string, topic string) error
	Publish(topic string, message Message) error
	CreateTopic(topic string) error
	DeleteTopic(topic string) error
	Shutdown()
}

func NewQueueWithChannelImpl(topics []string) Queue {
	fmt.Printf("Creating new queue with channel impl with topics %v \n", topics)
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
	for client, ch := range t.core {
		if debug {
			log.Printf("Shutdown %v subscription %s channel", t.name, client)
		}
		close(ch)
	}
}

func (q *qChannelImpl) Subscribe(clientID string, topic string) (chan Message, error) {
	topicCore, err := q.getTopic(topic)
	if err != nil {
		return nil, err
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()
	ch := make(chan Message)
	topicCore.core[clientID] = ch
	return ch, nil
}

func (q *qChannelImpl) Unsubscribe(clientID string, topic string) error {
	topicCore, err := q.getTopic(topic)
	if err != nil {
		return err
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()

	close(topicCore.core[clientID])

	delete(topicCore.core, clientID)
	return nil
}

func (q *qChannelImpl) Publish(topic string, message Message) error {
	topicCore, err := q.getTopic(topic)
	if err != nil {
		return err
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()
	for client, ch := range topicCore.core {
		if debug {
			fmt.Printf("Publishing message to topic %s for %s client\n", topic, client)
		}
		ch <- message
	}
	return nil
}

func (q *qChannelImpl) CreateTopic(topic string) error {
	k := Topic{
		name: topic,
		core: make(map[string]chan Message),
		lock: new(sync.Mutex),
	}
	q.core.Store(topic, k)
	q.topics = append(q.topics, topic)
	return nil
}

func (q *qChannelImpl) DeleteTopic(topic string) error {
	topicCore, err := q.getTopic(topic)
	if err != nil {
		return err
	}

	topicCore.lock.Lock()
	defer topicCore.lock.Unlock()

	for _, ch := range topicCore.core {
		close(ch)
	}

	delete(topicCore.core, topic)
	return nil

}

func (q *qChannelImpl) getTopic(topic string) (Topic, error) {
	val, ok := q.core.Load(topic)
	if !ok {
		return Topic{}, errors.New("Topic does not exist.")
	}

	topicCore, ok := val.(Topic)
	if !ok {
		return Topic{}, errors.New("Topic does not implement Topic.")
	}
	return topicCore, nil
}

func (q *qChannelImpl) Shutdown() {
	fmt.Println("Shutting down the message queue...")
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
