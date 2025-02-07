package calendars

import "errors"

func NewCalendarAdapter(adapterFlag string, config CalendarConfig) (CalendarAdapter, error) {
	switch adapterFlag {
		case "google", "default":
			return NewGoogleCalendarAdapter(config)
		default:
			return nil, errors.New("unknown calendar adapter")
	}
}
