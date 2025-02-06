package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string
	DatabaseURL  string
	AuthDomain   string
	AuthAudience string
	PaymentsKey  string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, reading from env: %v", err)
	}

	return &Config{
		Port: viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		AuthDomain: viper.GetString("AUTH_DOMAIN"),
		PaymentsKey: viper.GetString("PAYMENTS_KEY"),
	}
}