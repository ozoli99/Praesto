package integration

import (
	"log"

	"github.com/ozoli99/Praesto/pkg/appointment"
)

func SyncAppointmentToCalendar(appointment *appointment.Appointment) {
	// TODO: Integrate with Google Calendar, Outlook, etc.
	log.Printf("Syncing appointment %d to calendar: %v - %v", appointment.ID, appointment.StartTime, appointment.EndTime)
}

func RemoveAppointmentFromCalendar(appointment *appointment.Appointment) {
	// TODO: Remove event using the calendar API
	log.Printf("Removing appointment %d from calendar", appointment.ID)
}