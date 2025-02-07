package notifications

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ozoli99/Praesto/appointment"
)

type NotificationService interface {
	SendNotification(notification Notification) error
}

type NotificationService struct{}

func NewNotificationService() NotificationService {
	return &NotificationService{}
}

func (service *NotificationService) SendNotification(notification Notification) error {
	switch notification.Channel {
	case ChannelEmail:
		// TODO: Integrate with an email provider (SendGrid)
		log.Printf("Sending Email to %s: %s - %s", notification.Recipient, notification.Title, notification.Message)
	case ChannelSMS:
		if err := sendSMS(notification.Recipient, notification.Message); err != nil {
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

func sendSMS(to string, message string) error {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_FROM_PHONE")
	if accountSID == "" || authToken == "" || fromPhone == "" {
		log.Println("Twilio configuration not set; cannot send SMS")
		return nil
	}

	apiURL := "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"
	data := url.Values{}
	data.Set("From", fromPhone)
	data.Set("To", to)
	data.Set("Body", message)

	request, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	request.SetBasicAuth(accountSID, authToken)
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

func ScheduleReminder(appointment *appointment.Appointment) {
	reminderTime := appointment.StartTime.Add(-1 * time.Hour)
	if time.Now().After(reminderTime) {
		toPhone := os.Getenv("DUMMY_CUSTOMER_PHONE") // Retrieve a dummy phone number for testing or pull it from appointment data
		if toPhone == "" {
			toPhone = "+1234567890"
		}
		message := "Reminder: You have an appointment scheduled at " + appointment.StartTime.Format(time.RFC1123)
		if err := sendSMS(toPhone, message); err != nil {
			log.Printf("Error sending SMS reminder: %v", err)
		}
	} else {
		log.Printf("Scheduled reminder for appointment %d at %v", appointment.ID, reminderTime)
	}
}

func CancelReminder(appointment *appointment.Appointment) {
	log.Printf("Canceling reminder for appointment %d", appointment.ID)
	// TODO: Cancel the scheduled job if applicable
}
