package config

import "github.com/ankurnarkhede/golang-microservice-boilerplate/utils/logger"

// BaseConfig the application's configuration.
//nolint:maligned,unused // need same structure for better readability
type BaseConfig struct {
	Port      string
	Env       string
	LogConfig logger.LoggingConfig
}
