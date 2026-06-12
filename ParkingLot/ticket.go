package main

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id           string
	ParkingLevel *ParkingLevel
	Slot         *ParkingSlot
	Vehicle      *Vehicle
	EntryTime    time.Time
	ExitTime     time.Time
	AmountDue    float64
	IsPaid       bool
}

func NewTicket(parkingLevel *ParkingLevel, slot *ParkingSlot, vehicle *Vehicle) *Ticket {
	return &Ticket{
		Id:           uuid.NewString(),
		ParkingLevel: parkingLevel,
		Slot:         slot,
		Vehicle:      vehicle,
		EntryTime:    time.Now(),
		AmountDue:    0,
		IsPaid:       false,
	}
}

func (t *Ticket) GetDueAmount() float64 {
	// Calculate the amount and after successful payment, close the ticket
	// Let's say Fee is Rs. 10 per hour

	noOfHours := t.ExitTime.Sub(t.EntryTime).Hours()
	t.AmountDue = 10 * math.Ceil(noOfHours)
	return t.AmountDue
}
