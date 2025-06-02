package main

import (
	"fmt"
)

type VendingMachine struct {
	ID               string       `json:"id"`
	MachineState     MachineState `json:"machine_state"`
	IngredientStore  *IngredientStore
	PaymentProcessor PaymentProcessor
	Dispenser        Dispenser
	SelectedCoffee   CoffeeType
	InsertedMoney    int
	PaymentMethod    PaymentMethod
}

func CreateVendingMachine() *VendingMachine {

	ingredientStore := CreateStore(100)
	payementProcessor := &CashPayment{}

	vendingMachine := VendingMachine{
		ID:               "v1",
		MachineState:     Idle,
		IngredientStore:  ingredientStore,
		PaymentProcessor: payementProcessor,
		PaymentMethod:    Cash,
	}
	return &vendingMachine
}

func (v *VendingMachine) ChangeState(currentState MachineState, failure bool) {
	if currentState == Idle {
		v.MachineState = CashAcceptingState
	} else if currentState == CashAcceptingState {
		v.MachineState = SelectionState
	} else if currentState == SelectionState {
		v.MachineState = ProcessingState
	} else if currentState == ProcessingState {
		v.MachineState = DispenseState
	} else if currentState == ProcessingState && failure {
		v.MachineState = Idle
	} else if currentState == DispenseState {
		v.MachineState = Idle
	}
}

func (v *VendingMachine) ValidatePayment(insertedAmount int, coffee CoffeeType) bool {

	if insertedAmount < int(coffee.GetPrice()) {
		fmt.Println("Insufficient money inserted!!")
		return false
	}

	fmt.Println("Payment is Valid!!")
	return true
}

func (v *VendingMachine) StartMakingCoffee(money int) bool {
	failure := false
	v.ChangeState(v.MachineState, failure) // Idle -> CashAcceptingState
	v.InsertedMoney = money
	v.ChangeState(v.MachineState, failure) // CashAcceptingState -> SelectionState

	fmt.Println("what coffee do you want? espresso | latte | cappucino")

	var coffeeName string
	fmt.Scanln(&coffeeName)

	coffee := ReturnCoffee(coffeeName)
	if coffee == nil {
		fmt.Println("Some Error!!")
		return false
	}

	fmt.Println("coffee:", coffee)

	if !v.ValidatePayment(money, coffee) {
		return false
	}

	v.ChangeState(v.MachineState, failure) // SelectionState -> ProcessingState
	// coffee := coffeeFromFactory.GetCoffee()

	fmt.Println("Before ingredient Store: ", v.IngredientStore.IngredientMap)
	if !v.IngredientStore.ConsumeIngredient(coffee) {
		failure = true
		v.ChangeState(v.MachineState, failure) // insufficient ingredients => ProcessingState -> IdleState
		v.PaymentProcessor.ReturnMoney(coffee, money)
		return false
	}
	fmt.Println("After ingredient Store: ", v.IngredientStore.IngredientMap)

	v.ChangeState(v.MachineState, failure) // ProcessingState -> DispensingState
	dispenser := Dispenser{
		CoffeeObj: coffee,
	}

	v.Dispenser = dispenser

	v.Dispenser.DispenseCoffee(coffee)

	return true
}
