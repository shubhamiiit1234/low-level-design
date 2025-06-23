package main

type Subject interface {
	Register(observer Observer)
	DeRegister(observer Observer)
	NotifyAll(news string)
}
