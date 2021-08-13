package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ankurnarkhede/golang-microservice-boilerplate/config"
	"go.uber.org/zap"
)

func NewRouter(config config.BaseConfig, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Add CORS config
	// @todo: Refactor CORS to allow specific sources. Not working currently
	router.Use(cors.Default())

	v1Routes := router.Group("/api/v1")

	RegisterV1Routes(v1Routes, config, logger)

	return router
}
