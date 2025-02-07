package types

import "time"

type AppointmentData interface {
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetID() uint
}