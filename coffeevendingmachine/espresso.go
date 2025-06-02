package main

type Espresso struct {
	Price  int
	Recipe Recipe
}

func CreateEspresso() CoffeeType {
	espresso := &Espresso{}
	return espresso.GetCoffee()
}

func (e *Espresso) GetPrice() int {
	return 100
}

func (e *Espresso) GetCoffee() CoffeeType {
	price := e.GetPrice()
	ingredients := []Ingredient{
		{Name: Water, Quantity: 2},
		{Name: Coffee, Quantity: 1},
		{Name: Milk, Quantity: 3},
		{Name: Sugar, Quantity: 2},
	}
	recipe := Recipe{
		Ingredient: ingredients,
	}

	espresso := &Espresso{
		Price:  price,
		Recipe: recipe,
	}

	return espresso
}

func (e *Espresso) GetRecipe() Recipe {
	return e.Recipe
}
