package main

import "fmt"

/*
	Requirements:
		1. Publishers should be able to publish messages to a topic
		2. Subscribers should be able to suscribe to topics
		3. Subscribers should be able to receive messages published to those topics
		4. Multiple Publishers and Subscribers
		5. Messages should be delivered to all subscribers of a topic in real time
		6. System should handle concurrent access and ensure thread safety

	Core Entities:
		1. Subject interface
			- Register(), DeRegister(), Notify(msg)
		2. Topic (implements Subject)
			- Name, []Observer, chan, mutex
			- CreateTopic(name),
		3. Observer interface
			- Update(msg)
		4. Subscriber (implements Observer)
			- Id, chan
		5. MessageBroker
			- map[string]*Topic, mutex

	Flow:
		Producer produces/adds a message to a topic (into its channel) -> A go routine Notifies all the subcribers of that topic (call Update()) ->
																		-> Update() sends the message to the channels of each subscribers ->
																		-> A go routine Listen() will read the message from the channels

*/

func main() {
	fmt.Println("PubSub LLD")
}
