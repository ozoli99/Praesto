package notifications

import (
	"log"
	"time"

	"github.com/ozoli99/Praesto/pkg/appointment"
)

func ScheduleReminder(appointment *appointment.Appointment) {
	reminderTime := appointment.StartTime.Add(-1 * time.Hour)
	log.Printf("Scheduling reminder for appointment %d at %v", appointment.ID, reminderTime)

	// TODO: Use a job scheduler or message queue to send an email/SMS via SendGrid/Twilio
}

func CancelReminder(appointment *appointment.Appointment) {
	log.Printf("Canceling reminder for appointment %d", appointment.ID)
	// TODO: Cancel the scheduled job if applicable
}