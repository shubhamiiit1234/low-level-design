package main

type ParkingLevel struct {
	ParkingLevelID int           `json:"parking_level_id"`
	TotalSlots     int           `json:"total_slots"`
	ParkingSpots   []ParkingSpot `json:"parking_spots"`
	AvailableSlots int           `json:"available_slots"`
}

func NewParkingLevel(parkingLevelID, totalSlots int) *ParkingLevel {
	return &ParkingLevel{
		ParkingLevelID: parkingLevelID,
		TotalSlots:     totalSlots,
		ParkingSpots:   make([]ParkingSpot, totalSlots),
		AvailableSlots: totalSlots,
	}
}

func (pl *ParkingLevel) AddParkingSpot(spot ParkingSpot) {
	pl.ParkingSpots = append(pl.ParkingSpots, spot)
	pl.AvailableSlots++
}

func (pl *ParkingLevel) RemoveParkingSpot(spotID int) {
	// TODO
}
