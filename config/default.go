package config

import (
	"github.com/spf13/viper"
)

func setDefaultConfig(config *viper.Viper) {
	config.SetDefault("port", "3000")
	config.SetDefault("env", "development")
	config.SetDefault("LogConfig.EnableConsole", true)
	config.SetDefault("LogConfig.ConsoleJSONFormat", true)
	config.SetDefault("LogConfig.ConsoleLevel", "info")
	config.SetDefault("LogConfig.EnableFile", true)
	config.SetDefault("LogConfig.FileJSONFormat", true)
	config.SetDefault("LogConfig.FileLevel", "info")
	config.SetDefault("LogConfig.FileLocation", "./god.log")
}
