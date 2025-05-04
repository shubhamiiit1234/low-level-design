package main

type ParkingLevel struct {
	ParkingLevelID int           `json:"parking_level_id"`
	TotalSlots     int           `json:"total_slots"`
	ParkingSpots   []ParkingSpot `json:"parking_spots"`
	AvailableSlots int           `json:"available_slots"`
}
