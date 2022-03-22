package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Customer struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "root:pass@tcp(127.0.0.1:3306)/mydb?parseTime=true"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect into DB.")
	}
	DB.AutoMigrate(&Customer{})
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []Customer
	if dbc := DB.Find(&customers); dbc.Error != nil {
		fmt.Println(err.Error())
		panic("Cannot Retrieve customers.")
		return
	}
	json.NewEncoder(w).Encode(customers)
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var customer Customer
	if err := DB.First(&customer, params["id"]).Error; err != nil {
		fmt.Println(err.Error())
		panic("Cannot Retrieve customers " + params["id"])
		return
	}
	json.NewEncoder(w).Encode(customer)
}
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	DB.Create(&customer)
	json.NewEncoder(w).Encode(customer)
}
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var customer Customer
	DB.First(&customer, params["id"])
	json.NewDecoder(r.Body).Decode(&customer)
	DB.Save(&customer)
	json.NewEncoder(w).Encode(customer)
}
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var customer Customer
	if err := DB.Delete(&customer, params["id"]).Error; err != nil {
		fmt.Println(err.Error())
		panic("Cannot Delete customers " + params["id"])
		return
	}
	json.NewEncoder(w).Encode("Customer has been deleted.")
}
