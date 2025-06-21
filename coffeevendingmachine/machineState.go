package main

type MachineState string

const (
	Idle               MachineState = "IDLE"
	CashAcceptingState MachineState = "CASH_ACCEPTING"
	SelectionState     MachineState = "SELECTION"
	ProcessingState    MachineState = "PROCESSING"
	DispenseState      MachineState = "DISPENSE"
	OutOfServiceState  MachineState = "OUTOFSERVICE"
)
