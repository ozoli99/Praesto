package notifications

import "errors"

func NewNotificationServiceFactory(provider string, config NotificationConfig) (NotificationService, error) {
	switch provider {
		case "twilio", "default":
			return NewNotificationService(config), nil
		default:
			return nil, errors.New("Unknown notification provider")
	}
}