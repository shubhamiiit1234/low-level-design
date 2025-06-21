package main

// Strategy Pattern
type PaymentProcessor interface {
	ProcessPayment(CoffeeType, int) bool
	ReturnMoney(CoffeeType, int) int
}
