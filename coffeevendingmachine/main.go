package main

import (
	"fmt"
)

/*
	Requirements:

		Multiple coffee types - espresso, cappuccino, latte
		Ingredient management - track ingredient levels, prevent dispensing if insufficient
		Refill ingredients
		Process payment - process payment before dispensing
		Extensibility - easy to add new coffee type and payment method (strategy pattern)

	Core Entities:
		- Vending machine
		- Coffee (espresso, cappuccino, latte) (Strategy pattern)
		- IngredientStore (Milk, Coffee, Sugar, etc) (Extensible => strategy pattern)
		- PaymentProcessor - Cash, (TBD - UPI) (Strategy pattern)
		- Payment - Payment object (Transaction details)
		- Dispenser

*/

/*
	Flow -> Machine will be in idle state (Showing different coffee options) -> Press button to insert cash -> InsertPayment State -> Insert Cash -> SelectCoffee State -> Select Coffee -> Process Payment if
			enough cash inserted -> Processing State -> Check ingredients if enough to make coffee -> Dispense State -> Dispense the coffee -> Return change if any -> Idle State
															|
													Return Money -> Idle State
*/

func main() {
	fmt.Println("Welcome to the Coffee Vending Machine!")
	vendingMachine := CreateVendingMachine()

	fmt.Println("Espresso: 100Rs. | Latte: 120Rs. | Cappucino: 140Rs.")
	// Insert Cash
	fmt.Println("insert cash")
	var money int
	fmt.Scanln(&money)

	if !vendingMachine.StartMakingCoffee(money) {
		fmt.Println("Sorry!!")
	} else {
		fmt.Println("Enjoy!!")
	}

}
