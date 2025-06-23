package main

type Observer interface {
	Update(msg string)
	Listen()
}
