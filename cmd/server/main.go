package main

import (
	"email-validator-service/config"
	"email-validator-service/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	router := server.SetupRouter()

	log.Printf("Starting server on port %d...", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
