package main

import (
	"fmt"
	"time"
)

/*
	Requirements:
		- Multiple levels - each level has 20 slots
		- SpotType - CarSpot, TruckSpot, BikeSpot
		- Assign a spot to a vehicle based on its type
		- Release a spot when the vehicle leaves
		- Check Availability - check if a spot is available for a specific vehicle type
		- Generate a Parking Ticket
		- Payment/Transaction (TBD)

	Entities:
		- ParkingLot
		- ParkingLevel
		- ParkingSlot
		- SpotType (enum): CAR | TRUCK | BIKE
		- Vehicle
		- VehicleType (enum): CAR | TRUCK | BIKE
		- Ticket
		- Transaction (TBD)

	Methods:
		1. ParkingLot
			- AddParkingLevel(ParkingLevel), AddParkingSlot(ParkingLevel, ParkingSlot), GenerateTicket(Vehicle), IsSlotAvailable(Vehicle), Exit(ParkingTicket)
		2. ParkingLevel
			- AddParkingSlot(ParkingSlot), IsSlotAvailable(Vehicle)
		3. ParkingSlot
			- NewParkingSlot(SpotType), GetVehicle() string
		4. Vehicle
			- NewVehicle(VehicleType, vehicleNumber string)
		5. Ticket
			- GenerateTicket(Vehicle) Ticket, GetDueAmount(Ticket)


*/

func main() {
	fmt.Println("Parking Lot LLD!!!")

	parkingLevel := NewParkingLevel(0, 25, 10, 10, 5)
	parkingLot := NewParkingLot()
	parkingLot.AddParkingLevel(parkingLevel)

	// Add Bike Spots
	bikeParkingSpot := NewParkingSpot(BikeSpot)
	parkingLot.AddParkingSlot(0, 10, bikeParkingSpot)

	// Add Car Spots
	carParkingSpot := NewParkingSpot(CarSpot)
	parkingLot.AddParkingSlot(0, 10, carParkingSpot)

	// Add Truck Spots
	truckParkingSpot := NewParkingSpot(TruckSpot)
	parkingLot.AddParkingSlot(0, 5, truckParkingSpot)

	bike := NewVehicle(Bike, "KA51HV3776")
	// car := NewVehicle(Car, "HR26EQ5026")

	ticket := parkingLot.GenerateTicket(bike)
	fmt.Printf("TicketId: %s, Parking Level: %v, Slot: %s, Vehicle", ticket.Id, ticket.ParkingLevel.Id, ticket.Slot.Id)
	fmt.Println("\n--- TICKET DETAILS ---")
	fmt.Printf("ID          : %s\n", ticket.Id)
	fmt.Printf("Level       : %v\n", ticket.ParkingLevel.Id)
	fmt.Printf("Slot        : %s\n", ticket.Slot.Id)
	fmt.Printf("Vehicle Type  : %s\n", ticket.Vehicle.Type)
	fmt.Printf("Vehicle No  : %s\n", ticket.Vehicle.LicenseNumber)
	fmt.Printf("Entered At  : %s\n", ticket.EntryTime.Format(time.Kitchen)) // Prints like 3:04PM
	fmt.Printf("Amount Due  : Rs. %.2f\n", ticket.AmountDue)
	fmt.Printf("Is Settled  : %t\n", ticket.IsPaid) // %t formats booleans true/false
	fmt.Println("----------------------")

	if parkingLot.Exit(ticket) {
		fmt.Printf("Exited At  : %s\n", ticket.ExitTime.Format(time.Kitchen)) // Prints like 3:04PM
		fmt.Printf("Amount Due  : Rs. %.2f\n", ticket.AmountDue)
	}
}
