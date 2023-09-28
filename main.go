package main

import (
	"log"
	"websays/config"
)

func main() {
	log.Println("starting server")
	config.GetInstance().Setup("setup/prod.json")
	// httpHandler.GetInstance().Start()
}
