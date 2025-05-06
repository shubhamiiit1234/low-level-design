package main

import "fmt"

/*
	Requirements:
		Multiple levels - each level has 20 slots
		SpotType - CarSpot, TruckSpot, BikeSpot
		Assign a spot to a vehicle based on its type
		Release a spot when the vehicle leaves
		Check Availability - check if a spot is available for a specific vehicle type
		Customer
		Gate(Entrance/Exit) - to enter/exit the parking lot

*/

/*
	Entities:
		- ParkingLot (Central Class)
		- User (Customer, Admin)
		- Gate
		- Vehicle
		- ParkingSpot
		- ParkingLevel
		- ParkingTicket
*/

func main() {
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 level

	parkingLevels := parkingLot.Level
	fmt.Println("length of parkingLevels: ", len(parkingLevels))

	parkingTotalSpots := parkingLot.Level[0].TotalSlots
	fmt.Println("length of parkingSpots: ", parkingTotalSpots)

	vehicle := Vehicle{
		VehicleID:    1,
		Type:         Bike,
		LicensePlate: "KA51HV3776",
	}

	// suppose a vehicle comes from gate 1
	entryGate := parkingLot.Gate[0]
	entryGateAdmin := entryGate.Admin

	parkingTicket := entryGateAdmin.GenerateTicket(vehicle, parkingLot.Level)
	fmt.Println("Parking Ticket: ", *parkingTicket)

	parkingTicket.ExitTime = 3

	exitGate := parkingLot.Gate[1]
	exitGateAdmin := exitGate.Admin

	totalAmount := exitGateAdmin.ResolveTicket(*parkingTicket)
	fmt.Println("Total Amount: ", totalAmount)

	for _, spot := range parkingLot.Level[0].ParkingSpots {
		fmt.Println("Occupied: ", spot.IsOccupied)
	}
}
