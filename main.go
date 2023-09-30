package main

import (
	"log"
	"websays/config"
	"websays/httpHandler"
)

//The main function to start the app
func main() {
	log.Println("starting server...")
	//Read the config file
	config.GetInstance().Setup("setup/prod.json")
	//Start the mux server
	httpHandler.GetInstance().Start()
}
