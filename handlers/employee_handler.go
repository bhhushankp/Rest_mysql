package handlers

import (
	"encoding/json"
	"net/http"
	"rest_mysql/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	emp := models.Employee{}
	json.NewDecoder(r.Body).Decode(&emp)
	employee, err := models.CreateEmployee(&emp)
	if err != nil {
		http.Error(w, "Could not create Employee", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, employee)

}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", 400)
		return
	}
	employee, err := models.GetEmployee(ID)
	if err != nil {
		http.Error(w, "Employee not found.", http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, employee)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	emp := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid employee payload", 400)
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", 400)
		return
	}
	employee, err := models.UpdateEmployee(ID, &emp)
	if err != nil {
		http.Error(w, "Couldn't update the employee", 500)
		return
	}
	respondWithJSON(w, http.StatusOK, employee)

}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := models.GetEmployees()
	if err != nil {
		http.Error(w, "Could not get employees from DB", http.StatusInternalServerError)
	}
	respondWithJSON(w, http.StatusOK, employees)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", 400)
		return
	}
	err = models.DeleteEmployee(ID)
	if err != nil {
		http.Error(w, "Could not delete the employee", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Employee Deleted Succesfully!!"})

}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)

}
