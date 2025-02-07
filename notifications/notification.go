package notifications

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
