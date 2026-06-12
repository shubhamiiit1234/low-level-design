package main

type AvailabilityStrategy interface {
	GetAvailableSlot(parkingSlots []*ParkingSlot) *ParkingSlot
}
