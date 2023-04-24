package handlers

import (
	model "crmbackend/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Static HTML welcome page /
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "./static/index.html")
}

// Getting all the customers through a /customers route
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Customers)
}

// Get customer by id through a /customers/{id} route
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := model.Customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.Customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer not Found!")
	}
}

// Add customer through a /customers route
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer model.Customer

	data, _ := io.ReadAll(r.Body)
	json.Unmarshal(data, &newCustomer)

	newCustomer.ID = uuid.New().String()
	model.Customers[newCustomer.ID] = newCustomer
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCustomer)
}

// Update customer through a /customers/{id} route
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var receivedData map[string]string
	var updateCustomer model.Customer

	data, _ := io.ReadAll(r.Body)
	json.Unmarshal(data, &receivedData)
	json.Unmarshal(data, &updateCustomer)

	if customer, ok := model.Customers[id]; ok {
		for field := range receivedData {
			switch field {
			case "name":
				customer.Name = updateCustomer.Name
			case "role":
				customer.Role = updateCustomer.Role
			case "email":
				customer.Email = updateCustomer.Email
			case "phone":
				customer.Phone = updateCustomer.Phone
			case "contacted":
				customer.Contacted = updateCustomer.Contacted
			}
		}
		model.Customers[id] = customer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

// Delete customer through a /customers/{id} route
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := model.Customers[id]; ok {
		delete(model.Customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Customer deleted!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed!")
	}
}
