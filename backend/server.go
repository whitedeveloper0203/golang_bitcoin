package main

import (
	"log"
	"net/http"

	"bitbucket.com/cobinix/backend/handlers/retrieve"
	"bitbucket.com/cobinix/backend/handlers/signup"
	"bitbucket.com/cobinix/backend/handlers/update"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartServer handles http & calling appropriate funcs on POST/GET
func StartServer() {
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/userinfo", signup.CreateUser).Methods("POST")
	r.HandleFunc("/api/userinfo/{auth0uuid}", retrieve.GetUser).Methods("GET")
	r.HandleFunc("/api/userinfo/{auth0uuid}", update.DemoBalance).Methods("POST")

	// Wrap the router with cors middleware
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
