package main

type MachineState interface {
	GetState() States
	SetState(state States)

	AcceptMoney(vendingMachine *VendingMachine, money int)
	SelectProduct(vendingMachine *VendingMachine, product *Product) error
	DispenseProduct(vendingMachine *VendingMachine)
	ReturnChange() int
}
