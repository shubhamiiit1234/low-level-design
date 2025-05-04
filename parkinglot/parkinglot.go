package main

type ParkingLot struct {
	TotalLevels int            `json:"total_levels"`
	Level       []ParkingLevel `json:"level"`
	User        []User         `json:"user"`
	Gate        []Gate         `json:"gate"`
	Vehicle     []Vehicle      `json:"vehicle"`
}

func NewParkingLot(totalLevels int) *ParkingLot {
	return &ParkingLot{
		TotalLevels: totalLevels,
		Level:       make([]ParkingLevel, totalLevels),
		User:        []User{},
		Gate:        []Gate{},
		Vehicle:     []Vehicle{},
	}
}

func (pl *ParkingLot) AddParkingLevel(level ParkingLevel) {
	pl.Level = append(pl.Level, level)
	pl.TotalLevels++
}

func (pl *ParkingLot) RemoveParkingLevel(levelID int) {
	// TODO
}

func (pl *ParkingLot) AddUser(user User) {
	pl.User = append(pl.User, user)
}

func (pl *ParkingLot) RemoveUser(userID int) {
	// TODO
}

func (pl *ParkingLot) AddGate(gate Gate) {
	pl.Gate = append(pl.Gate, gate)
}

func (pl *ParkingLot) RemoveGate(gateID int) {
	// TODO
}

func (pl *ParkingLot) GetAllVehicles() []Vehicle {
	return pl.Vehicle
}
