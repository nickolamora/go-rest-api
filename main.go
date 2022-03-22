package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/customer", GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", GetCustomer).Methods("GET")
	router.HandleFunc("/customer", CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", DeleteCustomer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	InitialMigration()
	initializeRouter()
}
