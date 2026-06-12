package main

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	ParkingLevels map[int]*ParkingLevel
	mu            sync.RWMutex
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		ParkingLevels: map[int]*ParkingLevel{},
		mu:            sync.RWMutex{},
	}
}

func (p *ParkingLot) AddParkingLevel(parkingLevel *ParkingLevel) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exist := p.ParkingLevels[parkingLevel.Id]; exist {
		fmt.Println("Parking level already exists!!")
		return
	}

	p.ParkingLevels[parkingLevel.Id] = parkingLevel
}

func (p *ParkingLot) AddParkingSlot(parkingLevelId, numberOfSlots int, parkingSlot *ParkingSlot) {
	p.mu.Lock()
	defer p.mu.Unlock()

	parkingLevel, exist := p.ParkingLevels[parkingLevelId]
	if !exist {
		fmt.Println("Parking level does not exist!!")
		return
	}

	for range numberOfSlots {
		parkingLevel.AddParkingSlot(parkingSlot)
	}

}

func (p *ParkingLot) GenerateTicket(vehicle *Vehicle) *Ticket {
	// 1. Get Available Slot
	// 2. Assign slot to this vehicle
	// 3. Generate Ticket

	var (
		availableParkingLevel *ParkingLevel
		availableParkingSlot  *ParkingSlot
		available             bool
	)

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, parkingLevel := range p.ParkingLevels {
		availableParkingSlot, available = parkingLevel.GetAvailableSpot(vehicle, &FirstAvailableSlot{})
		if available {
			availableParkingLevel = parkingLevel
			break
		}
	}

	if !available {
		fmt.Println("No Parking Slot is available!")
		return nil
	}

	availableParkingLevel.AssignSpot(vehicle, availableParkingSlot)

	parkingTicket := NewTicket(availableParkingLevel, availableParkingSlot, vehicle)
	return parkingTicket
}

func (p *ParkingLot) Exit(ticket *Ticket) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Lock in case of multiple gates
	if !ticket.ExitTime.IsZero() {
		fmt.Println("This ticket has already been processed for exit")
		return false
	}

	ticket.ExitTime = time.Now().Add(2 * time.Hour)
	dueAmount := ticket.GetDueAmount()
	// Route to Payment Service
	// Release Spot only when payment is successful
	fmt.Printf("Payment of Rs. %f is Done!\n", dueAmount)
	ticket.Slot.ReleaseSpot()
	return true
}
