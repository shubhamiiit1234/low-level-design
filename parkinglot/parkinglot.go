package main

type ParkingLot struct {
	TotalLevels   int            `json:"total_levels"`
	Level         []ParkingLevel `json:"level"`
	User          []User         `json:"user"`
	Gate          []Gate         `json:"gate"`
	Vehicle       []Vehicle      `json:"vehicle"`
	ParkingTicket ParkingTicket  `json:"parking_ticket"`
}
