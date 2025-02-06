package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port         string
	DatabaseURL  string
	AuthDomain   string
	AuthAudience string
	PaymentsKey  string
}

func Load() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No configuration file found, reading from env: %v", err)
	}

	return &Configuration{
		Port:         viper.GetString("PORT"),
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		AuthDomain:   viper.GetString("AUTH0_DOMAIN"),
		AuthAudience: viper.GetString("AUTH0_AUDIENCE"),
		PaymentsKey:  viper.GetString("STRIPE_SECRET_KEY"),
	}
}