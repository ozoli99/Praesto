package appointment

import "time"

type Appointment struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	ProviderID uint `json:"provider_id"`
	Category string `json:"category"`
	TimeSlot time.Time `json:"time_slot"`
	Status string `json:"status"`
}