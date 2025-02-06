package notifications

import "log"

type Notifier interface {
	SendEmail(to, subject, body string) error
	SendSMS(to, message string) error
}

// DummyNotifier is a placeholder implementation that logs notifications.
type DummyNotifier struct{}

func NewDummyNotifier() *DummyNotifier {
	return &DummyNotifier{}
}

func (notifier *DummyNotifier) SendEmail(to, subject, body string) error {
	log.Printf("Sending Email -> To: %s, Subject: %s, Body: %s", to, subject, body)
	return nil
}

func (notifier *DummyNotifier) SendSMS(to, message string) error {
	log.Printf("Sending SMS -> To: %s, Message: %s", to, message)
	return nil
}