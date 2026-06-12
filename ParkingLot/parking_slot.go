package main

import "github.com/google/uuid"

type ParkingSlot struct {
	Id          string
	SpotType    string
	IsAvailable bool
	Vehicle     *Vehicle
}

func NewParkingSpot(spotType string) *ParkingSlot {
	return &ParkingSlot{
		Id:          uuid.NewString(),
		SpotType:    spotType,
		IsAvailable: true,
		Vehicle:     nil,
	}
}

func (p *ParkingSlot) GetVehicle() *Vehicle {
	if !p.IsAvailable {
		return p.Vehicle
	}
	return nil
}

func (p *ParkingSlot) AssignVehicle(vehicle *Vehicle) {
	p.Vehicle = vehicle
	p.IsAvailable = false
}

func (p *ParkingSlot) ReleaseSpot() {
	p.Vehicle = nil
	p.IsAvailable = true
}
