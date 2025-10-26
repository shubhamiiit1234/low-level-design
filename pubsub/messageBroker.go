package main

import (
	"fmt"
	"sync"
)

/*
	We can create another PublisherService which will have the MessageBroker
	and it will call its Publish() method inside anothe method like AnnounceResult()

	In this way it can be converted into a Notification System!!
*/

type MessageBroker struct {
	Topics map[string]*Topic // topicName -> topic
	mu     sync.RWMutex
}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		Topics: map[string]*Topic{},
	}
}

func (m *MessageBroker) AddTopic(id, topicName string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exist := m.Topics[topicName]; exist {
		fmt.Println("Topic ", topicName, " already exists!!!")
		return
	}

	newTopic := NewTopic(id, topicName)
	m.Topics[topicName] = newTopic

	go newTopic.Start()
}

func (m *MessageBroker) AddSubscriberToTopic(observer Observer, topicName string) {
	m.mu.RLock()
	topic, exist := m.Topics[topicName]
	m.mu.RUnlock()

	if !exist {
		fmt.Println("Topic doesn't exist!!!")
		return
	}
	topic.Register(observer)
}

func (m *MessageBroker) Publish(topicName, msg string) {
	m.mu.RLock()
	topic, exist := m.Topics[topicName]
	m.mu.RUnlock()

	if !exist {
		fmt.Println("Topic ", topicName, " doesn't exist!!!")
		return
	}
	go topic.AddMessage(msg)
}

func (m *MessageBroker) Listen(subscriber Observer) {
	subscriber.Listen()
}
