package main

type Product struct {
	ID    int
	Name  string
	Price int
}

func ProductCatalog() map[string]int {
	products := map[string]int{
		Pepsi: 20,
		Coke:  25,
	}

	return products
}
