package main

import (
	"log"
	"websays/config"
	"websays/httpHandler"
)

//The main function to start the app
func main() {
	log.Println("starting server")
	config.GetInstance().Setup("setup/prod.json")
	httpHandler.GetInstance().Start()
}
