package main

type InventoryManager struct {
	Inventory Inventory
}

func NewInventoryManager() *InventoryManager {

	// pepsi := Product{
	// 	ID:    1,
	// 	Name:  "Pepsi",
	// 	Price: 20,
	// }

	// coke := Product{
	// 	ID:    2,
	// 	Name:  "Coke",
	// 	Price: 40,
	// }

	inventory := Inventory{
		Products: map[string]int{
			Pepsi: 5,
			Coke:  10,
		},
	}

	return &InventoryManager{
		Inventory: inventory,
	}
}

func (i *InventoryManager) AddProduct(product Product, quantity int) {

}

func (i *InventoryManager) ConsumeProduct(product *Product) {
	i.Inventory.Products[product.Name] -= 1
}

func (i *InventoryManager) IsAvailable(product *Product) bool {
	return true
}
