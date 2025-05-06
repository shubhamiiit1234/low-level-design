package main

type PaymentMethod string

const (
	Cash PaymentMethod = "Cash"
	UPI  PaymentMethod = "UPI"
)

type PaymentStatus string

const (
	Pending   PaymentStatus = "Pending"
	Completed PaymentStatus = "Completed"
	Failed    PaymentStatus = "Failed"
)

type ParkingTicket struct {
	TicketID      int           `json:"ticket_id"`
	VehicleID     int           `json:"vehicle_id"`
	SpotID        int           `json:"spot_id"`
	EntryTime     int           `json:"entry_time"`
	ExitTime      int           `json:"exit_time"`
	TotalAmount   float64       `json:"total_amount"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}
