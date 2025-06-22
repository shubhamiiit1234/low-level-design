package main

import (
	"fmt"
)

/*
	Requirements:
		1. Multiple products with different prices and quantities
		2. Machine should accept coins and notes of different denominations
		3. Should dispense the product and return the change
		4. Inventory management (available products and their quantities)
		5. Should handle cases like - out-of-stock products / insufficient money

	Let's identify the core entities:
		1. VendingMachine (Central Class)
		2. Product
		3. Money - enum (const)
		4. Inventory
		5. MachineState (Idle, HasMoney, Selected, Dispense, OutOfStock)

	Happy Flow:
											|-> if out of stock -> show out of stock -> Idle State
											|
		Idle State/Selection State -> Select the product -> Selected State -> Insert Money (Accepting cash) -> Has Money State -> Press Button -> if enough cash inserted -> true -> Dispensed State -> Product will be dispensed -> Return change if any -> Idle State
																																			|
																																			|-> false -> return money -> Idle State
*/

func main() {
	fmt.Println("vending machine lld")

	vendingMachine := NewVendingMachine()

	// vendingMachine.InsertMoney(20) // Invalid in IdleState
	// vendingMachine.SelectProduct(&Product{ID: 1, Name: "Pepsi"})
	// vendingMachine.InsertMoney(20)
	// vendingMachine.DispenseProduct()

	fmt.Println("\n--- Second Transaction ---")
	vendingMachine.SelectProduct(&Product{ID: 1, Name: Coke})
	vendingMachine.InsertMoney(20)
	vendingMachine.DispenseProduct()

}
