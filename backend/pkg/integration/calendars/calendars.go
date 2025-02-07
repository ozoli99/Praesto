package calendars

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ozoli99/Praesto/pkg/appointment"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var calendarService *calendar.Service

func init() {
	context := context.Background()
	credFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credFile == "" {
		log.Println("GOOGLE_APPLICATION_CREDENTIALS not set; calendar integration will be disabled")
		return
	}
	service, err := calendar.NewService(context, option.WithCredentialsFile(credFile))
	if err != nil {
		log.Printf("Unable to create Google Calendar service: %v", err)
		return
	}
	calendarService = service
}

func SyncAppointmentToCalendar(appointment *appointment.Appointment) {
	if calendarService == nil {
		log.Println("Calendar service not initialized")
		return
	}
	calendarID := os.Getenv("GOOGLE_CALENDAR_ID")
	if calendarID == "" {
		log.Println("GOOGLE_CALENDAR_ID not set")
		return
	}

	event := &calendar.Event{
		Summary:     "Appointment",
		Description: "Service appointment",
		Start: &calendar.EventDateTime{
			DateTime: appointment.StartTime.Format(time.RFC3339),
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: appointment.EndTime.Format(time.RFC3339),
			TimeZone: "UTC",
		},
	}
	createdEvent, err := calendarService.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Printf("Unable to create event: %v", err)
		return
	}
	log.Printf("Event created: %s", createdEvent.HtmlLink)
}

func RemoveAppointmentFromCalendar(appointment *appointment.Appointment) {
	if calendarService == nil {
		log.Println("Calendar service not initialized")
		return
	}
	log.Printf("RemoveAppointmentFromCalendar: Not implemented because event ID is not stored for appointment %d", appointment.ID)
}
