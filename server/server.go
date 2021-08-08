package server

import "gitlab.com/awesome-life/god/config"

func Init() {
	config := config.GetConfig()
	router := NewRouter()

	router.Run(config.GetString("server.port"))
}
