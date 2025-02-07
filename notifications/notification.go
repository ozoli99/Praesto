package notifications

import (
	"github.com/ozoli99/Praesto/types"
)

type NotificationChannel int

const (
	ChannelEmail NotificationChannel = iota
	ChannelSMS
	ChannelPush
)

type Notification struct {
	Title     string              `json:"title"`
	Message   string              `json:"message"`
	Channel   NotificationChannel `json:"channel"`
	Recipient string              `json:"recipient"`
}

type NotificationService interface {
	SendNotification(notification Notification) error
	ScheduleReminder(appointment types.AppointmentData, config NotificationConfig)
	CancelReminder(appointment types.AppointmentData)
}
