package main

type Gate struct {
	GateID   int    `json:"gate_id"`
	GateType string `json:"gate_type"`
	AdminID  int    `json:"admin_id"`
}

type EntryGate struct {
	Gate `json:"gate"`
}

type ExitGate struct {
	Gate `json:"gate"`
}
