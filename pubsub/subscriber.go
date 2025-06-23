package main

import "fmt"

type Subscriber struct {
	ID string
	ch chan string
}

func NewSubscriber(id string) *Subscriber {
	return &Subscriber{
		ID: id,
		ch: make(chan string, 5),
	}
}

func (s *Subscriber) Update(msg string) {
	s.ch <- msg
}

func (s *Subscriber) Listen() {
	for msg := range s.ch {
		fmt.Printf("Subscriber %s received: %s\n", s.ID, msg)
	}
}
