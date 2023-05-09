package main

import (
	"EuprvaSaobracajna/config"
	"EuprvaSaobracajna/startup"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(&config)
	server.Start()
}
