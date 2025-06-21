package main

type Cappuccino struct {
	Price  int
	Recipe Recipe
}

func CreateCappuccino() CoffeeType {
	cappucino := &Cappuccino{}
	return cappucino.GetCoffee()
}

func (c *Cappuccino) GetPrice() int {
	return 140.0
}

func (e *Cappuccino) GetCoffee() CoffeeType {
	price := e.GetPrice()
	cappuccino := &Cappuccino{
		Price:  price,
		Recipe: Recipe{},
	}

	return cappuccino
}

func (c *Cappuccino) GetRecipe() Recipe {
	return c.Recipe
}
