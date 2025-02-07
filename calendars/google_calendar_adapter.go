package calendars

import (
	"context"
	"log"
	"time"
	"errors"

	"github.com/ozoli99/Praesto/types"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarAdapter struct {
	CalendarService *calendar.Service
	Configuration   CalendarConfig
}

func NewGoogleCalendarAdapter(config CalendarConfig) (CalendarAdapter, error) {
	if config.CredentialsFile == "" {
		return nil, errors.New("missing Google Calendar credentials")
	}
	context := context.Background()
	service, err := calendar.NewService(context, option.WithCredentialsFile(config.CredentialsFile))
	if err != nil {
		return nil, err
	}
	return &GoogleCalendarAdapter{
		CalendarService: service,
		Configuration: config,
	}, nil
}

func (adapter *GoogleCalendarAdapter) SyncAppointment(appointment types.AppointmentData) error {
	if adapter.CalendarService == nil {
		log.Println("Calendar service not initialized")
		return nil
	}
	if adapter.Configuration.CalendarID == "" {
		log.Println("CalendarID not provided in configuration")
		return nil
	}

	event := &calendar.Event{
		Summary: "Appointment",
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
	createdEvent, err := adapter.CalendarService.Events.Insert(adapter.Configuration.CalendarID, event).Do()
	if err != nil {
		log.Printf("Unable to create event: %v", err)
		return err
	}
	log.Printf("Google Calendar event created: %s", createdEvent.HtmlLink)
	return nil
}

func (adapter *GoogleCalendarAdapter) RemoveAppointment(appointment types.AppointmentData) error {
	if adapter.CalendarService == nil {
		log.Println("Calendar service not initialized")
		return nil
	}
	log.Printf("Google Calendar: RemoveAppointment not implemented for appointment %d", appointment.GetID())
	return nil
}