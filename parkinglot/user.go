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

func (a *Admin) GenerateTicket(vehicle Vehicle) ParkingTicket {
	// get all available parking spots and assign the vehicle to the first one

	return ParkingTicket{}
}

func (a *Admin) ResolveTicket(ticket ParkingTicket) float64 {

	// Calculate the total amount based on the time spent in the parking lot
	// Update the payment status to completed
	// Release the parking spot
	// return the total amount
	return 0.0
}

type Customer struct {
	User          `json:"user"`
	Vehicle       Vehicle       `json:"vehicle"`
	ParkingTicket ParkingTicket `json:"parking_ticket"`
}
