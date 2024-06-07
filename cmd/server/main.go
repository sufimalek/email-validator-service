package main

import (
	"email-validator-service/config"
	"email-validator-service/internal/server"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	r := server.NewRouter()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
