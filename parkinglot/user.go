package main

// type UserActionable interface {
// 	GenerateTicket(vehicle Vehicle) ParkingTicket
// 	ResolveTicket(ticket ParkingTicket) error
// }

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

type Admin struct {
	User `json:"user"`
}

func (a *Admin) GenerateTicket(vehicle Vehicle, parkingLevel []ParkingLevel) *ParkingTicket {

	// get all available parking spots and assign the vehicle to the first one
	var spotID int
	parkingLevel0 := parkingLevel[0]
	for i := range parkingLevel0.ParkingSpots {
		if !parkingLevel0.ParkingSpots[i].IsOccupied {
			spotID = parkingLevel0.ParkingSpots[i].SpotID
			parkingLevel0.ParkingSpots[i].MarkAsOccupied()
			break
		}
	}

	parkingTicket := &ParkingTicket{
		TicketID:      1,
		VehicleID:     vehicle.VehicleID,
		SpotID:        spotID,
		EntryTime:     1,
		PaymentStatus: Pending,
	}

	return parkingTicket
}

func (a *Admin) ResolveTicket(ticket ParkingTicket) float64 {

	// Calculate the total amount based on the time spent in the parking lot
	// Update the payment status to completed
	// Release the parking spot
	// return the total amount
	totalTime := float64(ticket.ExitTime) - float64(ticket.EntryTime)
	perHourCharge := 10
	return totalTime * float64(perHourCharge)
}

type Customer struct {
	User          `json:"user"`
	Vehicle       Vehicle       `json:"vehicle"`
	ParkingTicket ParkingTicket `json:"parking_ticket"`
}
