package main

// Strategy Pattern
type CoffeeType interface {
	GetPrice() int
	GetRecipe() Recipe
	GetCoffee() CoffeeType
}
