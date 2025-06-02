package main

import "fmt"

type IngredientStore struct {
	IngredientMap map[string]int
}

func CreateStore(quantity int) *IngredientStore {
	storeMap := map[string]int{
		Coffee: quantity,
		Milk:   quantity,
		Water:  quantity,
		Sugar:  quantity,
	}

	store := &IngredientStore{
		IngredientMap: storeMap,
	}

	return store
}

func (i *IngredientStore) ConsumeIngredient(coffee CoffeeType) bool {

	// fmt.Println("inside ConsumeIngredient: ", i.IngredientMap)

	recipe := coffee.GetRecipe()
	ingredients := recipe.Ingredient

	// fmt.Println("ingredients: ", ingredients)

	for _, ingredient := range ingredients {
		// fmt.Println(ingredient.Name)
		// fmt.Println(i.IngredientMap[ingredient.Name])
		if ingredient.Quantity > i.IngredientMap[ingredient.Name] {
			fmt.Println("Can't process, Insufficient ingredients")
			return false
		}
	}

	for _, ingredient := range ingredients {
		i.IngredientMap[ingredient.Name] -= ingredient.Quantity
	}

	return true
}

func (i *IngredientStore) RefilIngredient(ingredient Ingredient, quantity int) bool {
	i.IngredientMap[ingredient.Name] = quantity
	return true
}
