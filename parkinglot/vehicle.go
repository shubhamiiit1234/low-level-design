package main

type VehicleType string

const (
	Car   VehicleType = "Car"
	Truck VehicleType = "Truck"
	Bike  VehicleType = "Bike"
)

type Vehicle struct {
	VehicleID    int         `json:"vehicle_id"`
	LicensePlate string      `json:"license_plate"`
	Type         VehicleType `json:"type"`
}
