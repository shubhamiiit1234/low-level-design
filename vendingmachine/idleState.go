package main

import "fmt"

type IdleState struct {
	state States
}

func (i IdleState) GetState() States {
	return Idle
}

func (i *IdleState) SetState(state States) {
	i.state = Idle
}

func (i *IdleState) AcceptMoney(vendingMachine *VendingMachine, money int) {
	fmt.Println("Please Select Product before inserting cash!!!")
}

func (i *IdleState) SelectProduct(vendingMachine *VendingMachine, product *Product) error {
	fmt.Println("selecting the product")
	inventoryManager := vendingMachine.InventoryManager
	if isAvailable := inventoryManager.IsAvailable(product); !isAvailable {
		vendingMachine.ChangeState(&IdleState{state: Idle})
		// fmt.Println("item out of stock!!!")
		return fmt.Errorf("item out of stock")
	}

	vendingMachine.SetSelectedProduct(product)
	selectedState := &SelectedState{State: Selected}
	vendingMachine.ChangeState(selectedState) // State changed to Selected
	return nil
}

func (i *IdleState) DispenseProduct(vendingMachine *VendingMachine) {
	fmt.Println(("No item selected. Nothing to dispense!!!"))
}

func (i *IdleState) ReturnChange() int {
	fmt.Println("No money to return!!!")
	return 0
}
