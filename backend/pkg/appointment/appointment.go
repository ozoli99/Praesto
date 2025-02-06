package appointment

import (
	"time"
	"github.com/ozoli99/Praesto/pkg/models"
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

func TimeNow() time.Time {
	return time.Now()
}