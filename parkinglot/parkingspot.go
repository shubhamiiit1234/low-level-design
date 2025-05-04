package main

type ParkingSpot struct {
	SpotID       int         `json:"spot_id"`
	SpotType     VehicleType `json:"spot_type"`
	IsOccupied   bool        `json:"is_occupied"`
	Vehicle      Vehicle     `json:"vehicle"`
	AssignedTime string      `json:"assigned_time"`
	ReleasedTime string      `json:"released_time"`
}

func NewParkingSpot(spotID int, spotType VehicleType) *ParkingSpot {
	return &ParkingSpot{
		SpotID:     spotID,
		SpotType:   spotType,
		IsOccupied: false,
	}
}

func (ps *ParkingSpot) AssignVehicle(vehicle Vehicle, time string) {
	ps.Vehicle = vehicle
	ps.IsOccupied = true
	ps.AssignedTime = time
}

func (ps *ParkingSpot) ReleaseVehicle(time string) {
	ps.Vehicle = Vehicle{}
	ps.IsOccupied = false
	ps.ReleasedTime = time
}
