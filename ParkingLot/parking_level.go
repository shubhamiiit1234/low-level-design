package main

import (
	"fmt"
)

type ParkingLevel struct {
	Id           int // Denotes the floor number
	IsFull       bool
	ParkingSlots map[string][]*ParkingSlot
	Capacity     int
}

func NewParkingLevel(floor, capacity, bikeCapacity, carCapacity, truckCapacity int) *ParkingLevel {
	pl := &ParkingLevel{
		Id:           floor,
		IsFull:       false,
		ParkingSlots: make(map[string][]*ParkingSlot),
		Capacity:     capacity,
	}

	pl.ParkingSlots[BikeSpot] = make([]*ParkingSlot, 0, bikeCapacity)
	pl.ParkingSlots[CarSpot] = make([]*ParkingSlot, 0, carCapacity)
	pl.ParkingSlots[TruckSpot] = make([]*ParkingSlot, 0, truckCapacity)

	return pl
}

func (p *ParkingLevel) AddParkingSlot(parkingSlot *ParkingSlot) {
	if !p.IsFull {
		if len(p.ParkingSlots[parkingSlot.SpotType]) < cap(p.ParkingSlots[parkingSlot.SpotType]) {
			p.ParkingSlots[parkingSlot.SpotType] = append(p.ParkingSlots[parkingSlot.SpotType], parkingSlot)
		} else {
			fmt.Printf("Can't add %s slot in this parking level!", parkingSlot.SpotType)
		}
	} else {
		fmt.Println("This Parking Level is Full!")
	}
}

func (p *ParkingLevel) GetAvailableSpot(vehicle *Vehicle, availabilityStrategy AvailabilityStrategy) (*ParkingSlot, bool) {
	if p.IsFull {
		fmt.Println("This Parking Level is Full!")
		return nil, false
	}
	return availabilityStrategy.GetAvailableSlot(p.ParkingSlots[vehicle.Type]), true
}

func (p *ParkingLevel) AssignSpot(vehicle *Vehicle, parkingSlot *ParkingSlot) {
	parkingSlot.AssignVehicle(vehicle)
	p.ParkingSlots[vehicle.Type] = append(p.ParkingSlots[vehicle.Type], parkingSlot)
}
