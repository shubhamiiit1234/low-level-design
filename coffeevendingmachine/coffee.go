package main

type Coffee interface {
	GetPrice() float64
	// SetPrice(price float64)
	SetRecipe(recipe Recipe)
}
