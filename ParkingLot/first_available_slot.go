package main

type FirstAvailableSlot struct{}

func (f *FirstAvailableSlot) GetAvailableSlot(parkingSlots []*ParkingSlot) *ParkingSlot {
	for _, slot := range parkingSlots {
		if slot.IsAvailable {
			return slot
		}
	}
	return nil
}
