package main

type CoffeeFactory struct {
	CoffeeName string
}

func ReturnCoffee(name string) CoffeeType {
	if name == "espresso" {
		return CreateEspresso()
	} else if name == "latte" {
		return CreateLatte()
	} else if name == "cappucino" {
		return CreateCappuccino()
	}
	return nil
}
