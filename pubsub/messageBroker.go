package main

import (
	"fmt"
	"sync"
)

type MessageBroker struct {
	Topics map[string]*Topic // topicName -> topic
	mu     sync.RWMutex
}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		Topics: map[string]*Topic{},
	}
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

func (m *MessageBroker) AddTopic(id, topicName string) {
	m.mu.Lock()
	if _, exist := m.Topics[topicName]; exist {
		fmt.Println("Topic ", topicName, " already exists!!!")
		return
	}

	newTopic := NewTopic(id, topicName)
	m.Topics[topicName] = newTopic
	m.mu.Unlock()

	go newTopic.Start()
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
