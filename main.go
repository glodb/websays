package main

import (
	"log"
	"websays/config"
	"websays/httpHandler"
)

func main() {
	log.Println("starting server")
	config.GetInstance().Setup("setup/prod.json")
	httpHandler.GetInstance().Start()
}
