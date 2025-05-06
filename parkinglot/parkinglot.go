package main

import "fmt"

type ParkingLot struct {
	TotalLevels int            `json:"total_levels"`
	Level       []ParkingLevel `json:"level"`
	User        []User         `json:"user"`
	Gate        []Gate         `json:"gate"`
	Vehicle     []Vehicle      `json:"vehicle"`
}

func NewParkingLot(totalLevels int) *ParkingLot {
	entryGateAdmin := Admin{
		User: User{
			UserID:   1,
			Username: "admin1",
			Mobile:   "1234567890",
		},
	}

	exitGateAdmin := Admin{
		User: User{
			UserID:   2,
			Username: "admin2",
			Mobile:   "0987654321",
		},
	}

	entryGate := EntryGate{
		Gate: Gate{
			GateID:   1,
			GateType: EntryGateType,
			Admin:    entryGateAdmin,
		},
	}

	exitGate := ExitGate{
		Gate: Gate{
			GateID:   2,
			GateType: ExitGateType,
			Admin:    exitGateAdmin,
		},
	}

	parkingLevel := NewParkingLevel(1, 4) // Create a parking level with 4 slots
	fmt.Println("in lot: ", parkingLevel.TotalSlots)
	parkingLevel.ParkingSpots[0] = ParkingSpot{SpotID: 1, SpotType: Car}
	parkingLevel.ParkingSpots[1] = ParkingSpot{SpotID: 2, SpotType: Bike}
	parkingLevel.ParkingSpots[2] = ParkingSpot{SpotID: 3, SpotType: Truck}
	parkingLevel.ParkingSpots[3] = ParkingSpot{SpotID: 4, SpotType: Truck}

	fmt.Println("total spots: ", len(parkingLevel.ParkingSpots))

	parkingLevelList := make([]ParkingLevel, 0)
	parkingLevelList = append(parkingLevelList, *parkingLevel)

	return &ParkingLot{
		TotalLevels: totalLevels,
		Level:       parkingLevelList,
		User:        []User{entryGateAdmin.User, exitGateAdmin.User},
		Gate:        []Gate{entryGate.Gate, exitGate.Gate},
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

func (pl *ParkingLot) GetParkingLevels() []ParkingLevel {
	return pl.Level
}
