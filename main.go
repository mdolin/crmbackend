package main

import (
	handlers "crmbackend/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instantiate a new router
	router := mux.NewRouter()

	// Register handler functions to the same path
	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handlers.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	// Pass the customer router into ListenAndServe
	http.ListenAndServe(":3000", router)
}
