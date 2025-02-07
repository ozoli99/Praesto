package notifications

type NotificationConfig struct {
	TwilioAccountSID   string
	TwilioAuthToken    string
	TwilioFromPhone    string
	DummyCustomerPhone string
}
