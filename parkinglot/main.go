package main

import "fmt"

/*
	Multiple levels - each level has 20 slots

*/

type VehicleType string

const (
	Car   VehicleType = "Car"
	Truck VehicleType = "Truck"
	Bike  VehicleType = "Bike"
)

type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

func main() {
	fmt.Println("Welcome to the Parking Lot System!")
}
