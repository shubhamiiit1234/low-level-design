package main

import (
	"fmt"
	"time"
)

/*
	Requirements:
		1. Publishers should be able to publish messages to a topic
		2. Subscribers should be able to subscribe to topics
		3. Subscribers should be able to receive messages published to those topics
		4. Multiple Publishers and Subscribers
		5. Messages should be delivered to all subscribers of a topic in real time
		6. System should handle concurrent access and ensure thread safety

	Core Entities:
		1. Subject interface
			- Register(), DeRegister(), NotifyAll(msg)
		2. Topic (implements Subject)
			- Id, Name, []Observer, chan, mutex
			- CreateTopic(name), AddMessage(), Start()
		3. Observer interface
			- Update(msg)
		4. Subscriber (implements Observer)
			- Id, chan
		5. MessageBroker
			- map[string]*Topic, mutex
			- AddSubscriberToTopic(observer, topicID), Publish(topicID, msg), AddTopic(topicID, topicName)

	Flow:
		Producer produces/adds a message to a topic (into its channel) -> A goroutine Notifies all the subcribers of that topic (call Update()) ->
																		-> Update() sends the message to the channels of each subscribers ->
																		-> A goroutine Listen() will read the message from the channels

*/

func main() {
	fmt.Println("PubSub LLD!!!")

	messageBroker := NewMessageBroker()

	subscriber1 := NewSubscriber("S1")
	subscriber2 := NewSubscriber("S2")

	go messageBroker.Listen(subscriber1)
	go messageBroker.Listen(subscriber2)

	messageBroker.AddTopic("T1", "Sports")

	messageBroker.AddSubscriberToTopic(subscriber1, "Sports")
	messageBroker.AddSubscriberToTopic(subscriber2, "Sports")

	messageBroker.Publish("Sports", "Runs 120!!!")
	messageBroker.Publish("Sports", "Runs 121!!!")
	messageBroker.Publish("Sports", "Runs 122!!!")
	messageBroker.Publish("Sports", "Runs 124!!!")
	messageBroker.Publish("Sports", "Runs 125!!!")
	messageBroker.Publish("Sports", "Runs 128!!!")

	time.Sleep(1 * time.Second)

}
