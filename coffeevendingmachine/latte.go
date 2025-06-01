package main

type Latte struct {
	Price  float64
	Recipe Recipe
}

func (l *Latte) GetPrice() float64 {
	return 120.0
}
