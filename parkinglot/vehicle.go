package main

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
