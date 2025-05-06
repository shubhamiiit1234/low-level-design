package main

type GateType string

const (
	EntryGateType GateType = "Entry"
	ExitGateType  GateType = "Exit"
)

type Gate struct {
	GateID   int      `json:"gate_id"`
	GateType GateType `json:"gate_type"`
	Admin    Admin    `json:"admin"`
}

type EntryGate struct {
	Gate `json:"gate"`
}

type ExitGate struct {
	Gate `json:"gate"`
}
