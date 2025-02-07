package appointment

import (
	"time"

	"github.com/ozoli99/Praesto/models"
	"github.com/ozoli99/Praesto/types"
)

type AppointmentStatus string

const (
	StatusBooked      AppointmentStatus = "booked"
	StatusRescheduled AppointmentStatus = "rescheduled"
	StatusCanceled    AppointmentStatus = "canceled"
)

type Appointment struct {
	models.Base
	ProviderID uint              `json:"provider_id"`
	CustomerID uint              `json:"customer_id"`
	StartTime  time.Time         `json:"start_time"`
	EndTime    time.Time         `json:"end_time"`
	Status     AppointmentStatus `json:"status"`
}

func (appointment *Appointment) GetStartTime() time.Time {
	return appointment.StartTime
}

func (appointment *Appointment) GetEndTime() time.Time {
	return appointment.EndTime
}

func (appointment *Appointment) GetID() uint {
	return appointment.ID
}

var _ types.AppointmentData = (*Appointment)(nil)
