package main

type Latte struct {
	Price  int
	Recipe Recipe
}

func CreateLatte() CoffeeType {
	latte := &Latte{}
	return latte.GetCoffee()
}

func (l *Latte) GetPrice() int {
	return 120.0
}

func (l *Latte) GetCoffee() CoffeeType {
	price := l.GetPrice()
	ingredients := []Ingredient{
		{Name: Water, Quantity: 1},
		{Name: Coffee, Quantity: 2},
		{Name: Milk, Quantity: 2},
		{Name: Sugar, Quantity: 1},
	}
	recipe := Recipe{
		Ingredient: ingredients,
	}

	latte := &Latte{
		Price:  price,
		Recipe: recipe,
	}

	return latte
}

func (l *Latte) GetRecipe() Recipe {
	return l.Recipe
}
