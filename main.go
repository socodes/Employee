package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Surname  string  `json:"surname"`
	Speciality string `json:"speciality"`
	Department *Department `json:"department"`
}

// Author struct
type Department struct {
	DepartmentName string `json:"departmentName"`
	Manager  string `json:"manager"`
}

// Init books var as a slice Book struct
var employees []Employee

// Get all books
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// Get single book
func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

// Add new book
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

// Update book
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = params["id"]
			employees = append(employees, employee)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
}

// Delete book
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	employees = append(employees, Employee{ID: "1", Name: "John", Surname: "Break", Speciality:"GoLang", Department: &Department{DepartmentName: "Information Technologies", Manager: "Elon Musk"}})
	employees = append(employees, Employee{ID: "2", Name: "Walter", Surname: "Bishop", Department: &Department{DepartmentName: "Science", Manager: "Heisenberg White"}})

	// Route handles & endpoints
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", createEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	// Start server
	log.Fatal(http.ListenAndServe("https://mdidin-employee.herokuapp.com/", r))
}