package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port              string
	DatabaseURL       string
	AuthProvider      string
	AuthDomain        string
	AuthAudience      string
	AuthClientID      string
	AuthClientSecret  string
	AuthCallbackURL   string
	ClerkSecretKey    string
	ClerkPublicKeyURL string
	StripeSecretKey   string
	PaymentAdapter    string
	GoogleCredFile    string
	GoogleCalendarID  string
	TwilioSID         string
	TwilioAuthToken   string
	TwilioFromPhone   string
}

func Load(configFile string) *Configuration {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No configuration file found, reading from environment: %v", err)
	}

	configuration := &Configuration{
		Port:              viper.GetString("Port"),
		DatabaseURL:       viper.GetString("DatabaseURL"),
		AuthProvider:      viper.GetString("AuthProvider"),
		AuthDomain:        viper.GetString("AuthDomain"),
		AuthAudience:      viper.GetString("AuthAudience"),
		AuthClientID:      viper.GetString("AuthClientID"),
		AuthClientSecret:  viper.GetString("AuthClientSecret"),
		AuthCallbackURL:   viper.GetString("AuthCallbackURL"),
		ClerkSecretKey:    viper.GetString("ClerkSecretKey"),
		ClerkPublicKeyURL: viper.GetString("ClerkPublicKeyURL"),
		StripeSecretKey:   viper.GetString("StripeSecretKey"),
		PaymentAdapter:    viper.GetString("PaymentAdapter"),
		GoogleCredFile:    viper.GetString("GoogleCredFile"),
		GoogleCalendarID:  viper.GetString("GoogleCalendarID"),
		TwilioSID:         viper.GetString("TwilioSID"),
		TwilioAuthToken:   viper.GetString("TwilioAuthToken"),
		TwilioFromPhone:   viper.GetString("TwilioFromPhone"),
	}

	log.Printf("Loaded DatabaseURL: %s", configuration.DatabaseURL)
	return configuration
}
