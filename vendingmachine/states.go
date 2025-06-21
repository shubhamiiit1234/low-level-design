package main

type States string

const (
	Idle       States = "IDLE"
	HasMoney   States = "HASMONEY"
	Selected   States = "SELECTED"
	Dispensed  States = "DISPENSED"
	OutOfStock States = "OUTOFSTOCK"
)
