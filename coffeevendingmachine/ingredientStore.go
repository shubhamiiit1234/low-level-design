package main

type IngredientStore struct {
	CoffeeAmount int
	SugarAmount  int
	MilkAmount   float64
}

func CreateStore(amount, sugar int, milk float64) *IngredientStore {
	store := &IngredientStore{
		CoffeeAmount: amount,
		SugarAmount:  sugar,
		MilkAmount:   milk,
	}

	return store
}
