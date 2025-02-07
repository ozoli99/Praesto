package calendars

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ozoli99/Praesto/types"
)

type OutlookCalendarConfig struct {
	ClientID     string
	ClientSecret string
	TenantID     string
	CalendarID   string
	AccessToken  string
}

type OutlookCalendarAdapter struct {
	Configuration OutlookCalendarConfig
	Client        *http.Client
}

func NewOutlookCalendarAdapter(config OutlookCalendarConfig) (CalendarAdapter, error) {
	if config.ClientID == "" || config.ClientSecret == "" || config.TenantID == "" || config.CalendarID == "" || config.AccessToken == "" {
		return nil, errors.New("missing required Outlook configuration")
	}
	client := &http.Client{Timeout: 30 * time.Second}
	return &OutlookCalendarAdapter{
		Configuration: config,
		Client: client,
	}, nil
}

func (adapter *OutlookCalendarAdapter) SyncAppointment(appointment types.AppointmentData) error {
	payload := map[string]interface{}{
		"subject": "Appointment",
		"body": map[string]string{
			"contentType": "HTML",
			"content":     "Service appointment",
		},
		"start": map[string]string{
			"dateTime": appointment.GetStartTime().Format(time.RFC3339),
			"timeZone": "UTC",
		},
		"end": map[string]string{
			"dateTime": appointment.GetEndTime().Format(time.RFC3339),
			"timeZone": "UTC",
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Build the Graph API URL. For the default calendar, the endpoint is typically:
	// https://graph.microsoft.com/v1.0/me/calendars/{calendarID}/events
	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/calendars/%s/events", adapter.Configuration.CalendarID)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+adapter.Configuration.AccessToken)

	response, err := adapter.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("outlook API error: %s", string(body))
	}
	body, _ := ioutil.ReadAll(response.Body)
	log.Printf("Outlook event created: %s", body)
	return nil
}

func (adapter *OutlookCalendarAdapter) RemoveAppointment(appointment types.AppointmentData) error {
	// Stub implementation.
	log.Printf("Outlook: RemoveAppointment not fully implemented for appointment %d", appointment.GetID())
	return errors.New("RemoveAppointment not implemented")
}