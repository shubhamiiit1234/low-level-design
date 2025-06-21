package main

import "fmt"

type SelectedState struct {
	State States
}

func (s SelectedState) GetState() States {
	return Selected
}

func (s *SelectedState) SetState(state States) {
	s.State = Selected
}

func (i *SelectedState) AcceptMoney(vendingMachine *VendingMachine, money int) {
	if money < ProductCatalog()[vendingMachine.SelectedProduct.Name] {
		fmt.Println("insufficient amount inserted")
		return
	}
	vendingMachine.SetInsertedMoney(money)
	hasMoneyState := HasMoneyState{State: HasMoney}
	vendingMachine.ChangeState(&hasMoneyState)
}

func (i *SelectedState) SelectProduct(vendingMachine *VendingMachine, product *Product) error {
	fmt.Println("already selected the product!!!")
	return nil
}

func (i *SelectedState) DispenseProduct(vendingMachine *VendingMachine) {
	fmt.Println("Insert cash before dispensing!!!")
}

func (i *SelectedState) ReturnChange() int {
	fmt.Println("No Change!!!")
	return 0
}
