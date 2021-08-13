package main

import (
	"github.com/ankurnarkhede/golang-microservice-boilerplate/config"
	"github.com/ankurnarkhede/golang-microservice-boilerplate/server"
	logger "github.com/ankurnarkhede/golang-microservice-boilerplate/utils/logger"
)

func main() {
	cfg, _ := config.Init()
	log := logger.InitialiseLogger(cfg.LogConfig)
	server.Init(cfg, log)
}
