package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Config BaseConfig //nolint

// Init is an exported method that takes the environment starts the
// viper (external lib) and returns the configuration struct.
// Preference: ENV VARIABLES > VARIABLES DEFINED IN FILE > DEFAULTS.
func Init() (BaseConfig, error) {
	var cfg *viper.Viper

	// Load environment variables from .env if available
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: No .env file found")
	}

	cfg = viper.New()

	// set default configs
	setDefaultConfig(cfg)

	// default viper configs
	// Setting `_` as the separater for config keys containing `.` in its name
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.SetConfigType("json")
	cfg.SetConfigName(cfg.GetString("env"))
	cfg.AddConfigPath("../config/")
	cfg.AddConfigPath("config/")

	err = cfg.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

	err = cfg.Unmarshal(&Config)
	if err != nil {
		log.Fatal("Error on unmarshalling configuration")
	}

	return Config, nil
}

// GetConfig : To be used across the app to get config.
func GetConfig() BaseConfig {
	return Config
}
