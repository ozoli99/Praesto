package calendars

import "errors"

func NewCalendarAdapter(adapterFlag string, config interface{}) (CalendarAdapter, error) {
	switch adapterFlag {
		case "google", "default":
			googleConfig, ok := config.(CalendarConfig)
			if !ok {
				return nil, errors.New("invalid configuration for Google Calendar adapter")
			}
			return NewGoogleCalendarAdapter(googleConfig)
		case "outlook":
			outlookConfig, ok := config.(OutlookCalendarConfig)
			if !ok {
				return nil, errors.New("invalid configuration for Outlook Calendar adapter")
			}
			return NewOutlookCalendarAdapter(outlookConfig)
		case "apple":
			appleConfig, ok := config.(AppleCalendarConfig)
			if !ok {
				return nil, errors.New("invalid configuration for Apple Calendar adapter")
			}
			return NewAppleCalendarAdapter(appleConfig)
		default:
			return nil, errors.New("unknown calendar adapter")
	}
}
