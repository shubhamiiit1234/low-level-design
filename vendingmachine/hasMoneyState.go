package main

import "fmt"

type HasMoneyState struct {
	State States
}

func (h HasMoneyState) GetState() States {
	return HasMoney
}

func (h *HasMoneyState) SetState(state States) {
	h.State = HasMoney
}

func (h *HasMoneyState) AcceptMoney(vendingMachine *VendingMachine, money int) {
	fmt.Println("Cash is already inserted!!!")
}

func (h *HasMoneyState) SelectProduct(vendingMachine *VendingMachine, product *Product) error {
	fmt.Println("Product already selected!!!")
	return nil
}

func (i *HasMoneyState) DispenseProduct(vendingMachine *VendingMachine) {
	fmt.Println("We have cash, Dispensing is in progress!!!")
	vendingMachine.ChangeState(&DispensedState{State: Dispensed})
	fmt.Println("state: ", vendingMachine.State)
	vendingMachine.State.DispenseProduct(vendingMachine)
}

func (h *HasMoneyState) ReturnChange() int {
	return 0
}

func (i *HasMoneyState) IsEnoughCashInserted(vendingMachine *VendingMachine) bool {
	return vendingMachine.InsertedMoney >= vendingMachine.SelectedProduct.Price
}
