package main

type Cappuccino struct {
	Price  float64
	Recipe Recipe
}

func (c *Cappuccino) GetPrice() float64 {
	return 140.0
}
