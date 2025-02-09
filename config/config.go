package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port             string
	DatabaseURL      string
	AuthDomain       string
	AuthAudience     string
	AuthClientID     string
	AuthClientSecret string
	AuthCallbackURL  string
	StripeSecretKey  string
	PaymentAdapter   string
	GoogleCredFile   string
	GoogleCalendarID string
	TwilioSID        string
	TwilioAuthToken  string
	TwilioFromPhone  string
}

func Load(configFile string) *Configuration {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No configuration file found, reading from environment: %v", err)
	}

	return &Configuration{
		Port:             viper.GetString("PORT"),
		DatabaseURL:      viper.GetString("DATABASE_URL"),
		AuthDomain:       viper.GetString("AUTH0_DOMAIN"),
		AuthAudience:     viper.GetString("AUTH0_AUDIENCE"),
		AuthClientID:     viper.GetString("AUTH0_CLIENT_ID"),
		AuthClientSecret: viper.GetString("AUTH0_CLIENT_SECRET"),
		AuthCallbackURL:  viper.GetString("AUTH0_CALLBACK_URL"),
		StripeSecretKey:  viper.GetString("STRIPE_SECRET_KEY"),
		PaymentAdapter:   viper.GetString("PAYMENT_ADAPTER"),
		GoogleCredFile:   viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"),
		GoogleCalendarID: viper.GetString("GOOGLE_CALENDAR_ID"),
		TwilioSID:        viper.GetString("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:  viper.GetString("TWILIO_AUTH_TOKEN"),
		TwilioFromPhone:  viper.GetString("TWILIO_FROM_PHONE"),
	}
}
