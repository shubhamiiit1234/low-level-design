package main

import (
	"fmt"
	"sync"
)

type Topic struct {
	ID          string
	Name        string
	ch          chan string
	Subscribers []Observer
	mu          sync.RWMutex
}

func NewTopic(id, name string) *Topic {
	return &Topic{
		ID:          id,
		Name:        name,
		ch:          make(chan string, 5),
		Subscribers: []Observer{},
		mu:          sync.RWMutex{},
	}
}

func (t *Topic) Register(observer Observer) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Subscribers = append(t.Subscribers, observer)
	fmt.Println("subscribed!!!")
}

func (t *Topic) DeRegister(observer Observer) {
	for i, obs := range t.Subscribers {
		if obs == observer {
			t.Subscribers = append(t.Subscribers[:i], t.Subscribers[i+1:]...)
			fmt.Println("UnSubscribed!!!")
			break
		}
	}
}

func (t *Topic) AddMessage(msg string) {
	t.mu.Lock()
	t.ch <- msg
	t.mu.Unlock()
}

func (t *Topic) Start() {
	for msg := range t.ch {
		t.NotifyAll(msg)
	}
}

func (t *Topic) NotifyAll(msg string) {
	t.mu.RLock()

	for _, subscriber := range t.Subscribers {
		subscriber.Update(msg)
	}
	t.mu.RUnlock()
}
