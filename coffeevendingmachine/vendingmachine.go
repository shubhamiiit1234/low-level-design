package main

type VendingMachine struct {
	ID               string       `json:"id"`
	MachineState     MachineState `json:"machine_state"`
	IngredientStore  IngredientStore
	PaymentProcessor PaymentProcessor
	Dispenser        Dispenser
	SelectedCoffee   Coffee
	InsertedMoney    float64
	PaymentMethod    PaymentMethod
	// Coffee           []Coffee
}
