package config

import (
	"log"

)

var Config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.SetConfigName(env)
	Config.AddConfigPath("../config/")
	Config.AddConfigPath("config/")
	err = Config.ReadInConfig()

	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return Config
}
