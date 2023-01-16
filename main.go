package main

import (
	"log"

	"mateuszgua/to-do-list/server"
	"mateuszgua/to-do-list/server/router"
	config "mateuszgua/to-do-list/utils"
)

func main() {
	config := config.LoadConfig()
	router, err := router.MyRouter()
	if err != nil {
		log.Fatal("failed to add router", err)
	}

	err = server.MyServer(config.HttpPort, router)
	if err != nil {
		log.Fatal("failed connect with server", err)
	}
}
