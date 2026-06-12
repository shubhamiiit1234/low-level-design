package main

import "github.com/google/uuid"

type Vehicle struct {
	Id            string
	Type          string
	LicenseNumber string
}

func NewVehicle(vehicleType, licenseNo string) *Vehicle {
	return &Vehicle{
		Id:            uuid.NewString(),
		Type:          vehicleType,
		LicenseNumber: licenseNo,
	}
}
