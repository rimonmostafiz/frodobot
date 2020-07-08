// Package cfg contains configuration related things
package cfg

import (
	"github.com/spf13/viper"
	"log"
)

// Viper instance for configuration
var Viper = viper.New()

// InitViper initializes viper config file
func InitViper(cfg string) {
	Viper.SetConfigFile(cfg)
}

// ReadFromEnv reads a string value for given key from env file
func ReadFromEnv(key string) string {
	checkConfiguration()

	value, ok := Viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

// ReadStringMapFromEnv reads a map[string]string for given key from env file
func ReadStringMapFromEnv(key string) map[string]string {
	checkConfiguration()

	value := Viper.GetStringMapString(key)
	return value
}

// checkConfiguration will discover and load the configuration file from disk
// If not exists it will throw error
func checkConfiguration() {
	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}
