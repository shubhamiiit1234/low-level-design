package main

import "fmt"

type DispensedState struct {
	State States
}

func (d DispensedState) GetState() States {
	return Idle
}

func (d *DispensedState) SetState(state States) {
	d.State = Dispensed
}

func (d *DispensedState) AcceptMoney(vendingMachine *VendingMachine, money int) {
	fmt.Println("Cash is already inserted!!!")
}

func (d *DispensedState) SelectProduct(vendingMachine *VendingMachine, product *Product) error {
	fmt.Println("Already selected the product")
	return nil
}

func (d *DispensedState) DispenseProduct(vendingMachine *VendingMachine) {
	dispensedProduct := vendingMachine.SelectedProduct
	vendingMachine.InventoryManager.ConsumeProduct(dispensedProduct)
	vendingMachine.ChangeState(&IdleState{state: Idle})
}

func (d *DispensedState) ReturnChange() int {
	return 0
}
