package calendars

import "github.com/ozoli99/Praesto/types"

type CalendarAdapter interface {
	SyncAppointment(appointment types.AppointmentData) error
	RemoveAppointment(appointment types.AppointmentData) error
}