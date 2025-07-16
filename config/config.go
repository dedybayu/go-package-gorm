package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

var ENV *Config

func LoadConfig() {
	log.Println("Loading configuration...")
	
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {	
		panic("Error reading config file: " + err.Error())
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	// log.Println("Configuration loaded successfully")
	// log.Printf("Server will run on port: %s", ENV.PORT)
	// log.Printf("Database Name: %s", ENV.DB_NAME)
	// log.Printf("Database Username: %s", ENV.DB_USERNAME)
}
