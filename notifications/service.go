package notifications

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ozoli99/Praesto/appointment"
)

type notificationService struct {
	Config NotificationConfig
}

func NewNotificationService(config NotificationConfig) NotificationService {
	return &notificationService{Config: config}
}

func (service *notificationService) SendNotification(notification Notification) error {
	switch notification.Channel {
	case ChannelEmail:
		// TODO: Integrate with an email provider (SendGrid)
		log.Printf("Sending Email to %s: %s - %s", notification.Recipient, notification.Title, notification.Message)
	case ChannelSMS:
		if err := sendSMS(notification.Recipient, notification.Message, service.Config); err != nil {
			log.Printf("Error sending SMS to %s: %v", notification.Recipient, err)
			return err
		}
	case ChannelPush:
		// TODO: Integrate with a push notification provider (Firebase Cloud Messaging)
		log.Printf("Sending Push Notification to %s: %s - %s", notification.Recipient, notification.Title, notification.Message)
	default:
		log.Printf("Unknown notification channel for recipient %s", notification.Recipient)
	}
	return nil
}

func sendSMS(to string, message string, config NotificationConfig) error {
	if config.TwilioAccountSID == "" || config.TwilioAuthToken == "" || config.TwilioFromPhone == "" {
		log.Println("Twilio configuration not set; cannot send SMS")
		return nil
	}

	apiURL := "https://api.twilio.com/2010-04-01/Accounts/" + config.TwilioAccountSID + "/Messages.json"
	data := url.Values{}
	data.Set("From", config.TwilioFromPhone)
	data.Set("To", to)
	data.Set("Body", message)

	request, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	request.SetBasicAuth(config.TwilioAccountSID, config.TwilioAuthToken)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		log.Println("SMS sent successfully")
	} else {
		log.Printf("Failed to send SMSl status: %s", response.Status)
	}
	return nil
}

func (service *notificationService) ScheduleReminder(appointment *appointment.Appointment, config NotificationConfig) {
	reminderTime := appointment.StartTime.Add(-1 * time.Hour)
	if time.Now().After(reminderTime) {
		toPhone := config.DummyCustomerPhone
		if toPhone == "" {
			toPhone = "+1234567890"
		}
		message := "Reminder: You have an appointment scheduled at " + appointment.StartTime.Format(time.RFC1123)
		notification := Notification{
			Title:     "Appointment Reminder",
			Message:   message,
			Channel:   ChannelSMS,
			Recipient: toPhone,
		}
		if err := service.SendNotification(notification); err != nil {
			log.Printf("Error sending SMS reminder: %v", err)
		}
	} else {
		log.Printf("Scheduled reminder for appointment %d at %v", appointment.ID, reminderTime)
	}
}

func (service *notificationService) CancelReminder(appointment *appointment.Appointment) {
	log.Printf("Canceling reminder for appointment %d", appointment.ID)
	// TODO: Cancel the scheduled job if applicable
}
