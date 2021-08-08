package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/awesome-life/god/config"
	"gitlab.com/awesome-life/god/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
