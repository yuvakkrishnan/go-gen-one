package server

import (
	handlers "go-gen-one/handler"

	"github.com/gorilla/mux"
)

// NewRouter initializes mux router with all endpoints
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/jwt", handlers.JWTHandler).Methods("GET")
	r.HandleFunc("/hash", handlers.HashHandler).Methods("GET")
	r.HandleFunc("/resty", handlers.RestyHandler).Methods("GET")
	return r
}
