package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MyServer(httpPort string, router *mux.Router) error {
	if httpPort == "" {
		httpPort = "8080"
	}

	serverUrl := fmt.Sprintf(":%s", httpPort)
	log.Printf("Starting server on http://localhost%s", serverUrl)

	err := http.ListenAndServe(serverUrl, nil)
	if err != nil {
		log.Fatal("failed to start server", err)
	}

	return nil

}
