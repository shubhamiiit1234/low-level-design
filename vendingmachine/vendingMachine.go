package main

import "fmt"

type VendingMachine struct {
	State            MachineState
	InventoryManager *InventoryManager
	InsertedMoney    int
	SelectedProduct  *Product
}

func NewVendingMachine() *VendingMachine {

	inventoryManager := NewInventoryManager()
	return &VendingMachine{
		State:            &IdleState{state: Idle},
		InventoryManager: inventoryManager,
	}
}

func (v *VendingMachine) ChangeState(machineState MachineState) {
	v.State = machineState
}

func (v *VendingMachine) SetInsertedMoney(money int) {
	v.InsertedMoney = money
}

func (v *VendingMachine) SetSelectedProduct(product *Product) {
	v.SelectedProduct = product
}

func (v *VendingMachine) SetInvetoryManager(inventoryManager *InventoryManager) {
	v.InventoryManager = inventoryManager
}

func (v *VendingMachine) SelectProduct(product *Product) {
	v.State.SelectProduct(v, product)
}

func (v *VendingMachine) InsertMoney(money int) {
	v.State.AcceptMoney(v, money)
}

func (v *VendingMachine) DispenseProduct() {

	v.State.DispenseProduct(v)
	fmt.Println("Product dispensed: ", v.SelectedProduct.Name)
	fmt.Println("remaining: ", v.InventoryManager.Inventory.Products[v.SelectedProduct.Name])
}
