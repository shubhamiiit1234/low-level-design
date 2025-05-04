package main

type ParkingSpot struct {
	SpotID       int         `json:"spot_id"`
	SpotType     VehicleType `json:"spot_type"`
	IsOccupied   bool        `json:"is_occupied"`
	Vehicle      Vehicle     `json:"vehicle"`
	AssignedTime string      `json:"assigned_time"`
	ReleasedTime string      `json:"released_time"`
}
