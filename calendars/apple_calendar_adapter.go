package calendars

import (
	"errors"
	"fmt"
	"log"

	"github.com/emersion/go-caldav"
	"github.com/ozoli99/Praesto/types"
)

type AppleCalendarConfig struct {
	ServerURL  string // CalDAV server URL
	Username   string
	Password   string
	CalendarID string
}

type AppleCalendarAdapter struct {
	Configuration AppleCalendarConfig
	Client        *caldav.Client
}

func NewAppleCalendarAdapter(config AppleCalendarConfig) (CalendarAdapter, error) {
	if config.ServerURL == "" || config.Username == "" || config.Password == "" {
		return nil, errors.New("missing Apple Calendar credentials")
	}
	if config.CalendarID == "" {
		return nil, errors.New("missing Apple Calendar ID")
	}
	client, err := caldav.Dial(config.ServerURL, &caldav.Options{
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return nil, err
	}
	return &AppleCalendarAdapter{
		Configuration: config,
		Client: client,
	}, nil
}

func (adapter *AppleCalendarAdapter) SyncAppointment(appointment types.AppointmentData) error {
	eventStr := fmt.Sprintf("BEGIN:VCALENDAR\r\nVERSION:2.0\r\nBEGIN:VEVENT\r\nUID:%d@apple\r\nDTSTART:%s\r\nDTEND:%s\r\nSUMMARY:Appointment\r\nDESCRIPTION:Service appointment\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		appointment.GetID(),
		appointment.GetStartTime().Format("20060102T150405Z"),
		appointment.GetEndTime().Format("20060102T150405Z"),
	)

	calendarURL := fmt.Sprintf("%s/calendars/%s/%s/", adapter.Configuration.ServerURL, adapter.Configuration.Username, adapter.Configuration.CalendarID)
	createdEvent, err := adapter.Client.CreateEvent(calendarURL, eventStr)
	if err != nil {
		log.Printf("Apple Calendar: Unable to create event: %v", err)
		return err
	}
	log.Printf("Apple Calendar event created: %v", createdEvent)
	return nil
}

func (adapter *AppleCalendarAdapter) RemoveAppointment(appointment types.AppointmentData) error {
	// Stub implementation.
	log.Printf("Apple Calendar: RemoveAppointment not fully implemented for appointment %d", appointment.GetID())
	return errors.New("RemoveAppointment not implemented")
}