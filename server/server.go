package server

import (
	"github.com/ankurnarkhede/golang-microservice-boilerplate/config"
	"go.uber.org/zap"
)

// Init : Initialize the server.
func Init(config config.BaseConfig, logger *zap.Logger) {
	router := NewRouter(config, logger)

	router.Run(":" + config.Port)
}
