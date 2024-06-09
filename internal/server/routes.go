package server

import (
	"github.com/gorilla/mux"
)

// SetupRouter sets up the routes and middleware
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Apply middleware
	router.Use(LoggingMiddleware)

	// Setup routes
	router.HandleFunc("/api/validate", ValidateEmailHandler).Methods("POST")

	return router
}
