package server

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/validate", ValidateEmail).Methods("POST")
	return r
}
