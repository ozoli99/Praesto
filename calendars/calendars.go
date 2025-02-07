package calendars

import (
	"context"
	"log"
	"time"

	"github.com/ozoli99/Praesto/types"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var calendarService *calendar.Service

func InitCalendarService(calendarConfig CalendarConfig) {
	if calendarConfig.CredentialsFile == "" {
		log.Println("Credentials file not provided; calendar integration disabled")
		return
	}
	context := context.Background()
	service, err := calendar.NewService(context, option.WithCredentialsFile(calendarConfig.CredentialsFile))
	if err != nil {
		log.Printf("Unable to create Google Calendar service: %v", err)
		return
	}
	calendarService = service
}

func SyncAppointmentToCalendar(appointment types.AppointmentData, calendarConfig CalendarConfig) {
	if calendarService == nil {
		log.Println("Calendar service not initialized")
		return
	}
	if calendarConfig.CalendarID == "" {
		log.Println("CalendarID not provided in config")
		return
	}

	event := &calendar.Event{
		Summary:     "Appointment",
		Description: "Service appointment",
		Start: &calendar.EventDateTime{
			DateTime: appointment.GetStartTime().Format(time.RFC3339),
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: appointment.GetEndTime().Format(time.RFC3339),
			TimeZone: "UTC",
		},
	}
	createdEvent, err := calendarService.Events.Insert(calendarConfig.CalendarID, event).Do()
	if err != nil {
		log.Printf("Unable to create event: %v", err)
		return
	}
	log.Printf("Event created: %s", createdEvent.HtmlLink)
}

func RemoveAppointmentFromCalendar(appointment types.AppointmentData) {
	if calendarService == nil {
		log.Println("Calendar service not initialized")
		return
	}
	log.Printf("RemoveAppointmentFromCalendar: Not implemented because event ID is not stored for appointment %d", appointment.GetID())
}
