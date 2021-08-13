package server

import (
	"github.com/ankurnarkhede/golang-microservice-boilerplate/config"
	"github.com/ankurnarkhede/golang-microservice-boilerplate/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterV1Routes : registers routes, config and logger in the package.
func RegisterV1Routes(router *gin.RouterGroup, config config.BaseConfig, logger *zap.Logger) {
	router.GET("/health", controllers.Health)

	logger.Info("v1 routes injected successfully")
}
