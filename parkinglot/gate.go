package main

type Gate struct {
	GateID   int    `json:"gate_id"`
	GateType string `json:"gate_type"`
}

type EntryGate struct {
	Gate `json:"gate"`
}

type ExitGate struct {
	Gate `json:"gate"`
}
