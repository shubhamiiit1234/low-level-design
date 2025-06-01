package main

type Espresso struct {
	Price  float64
	Recipe Recipe
}

func (e *Espresso) GetPrice() float64 {
	return 100.0
}

func (e *Espresso) GetEspresso(recipe Recipe) Espresso {
	price := e.GetPrice()
	espresso := Espresso{
		Price:  price,
		Recipe: recipe,
	}

	return espresso
}
