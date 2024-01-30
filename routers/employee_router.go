package routers

import (
	"rest_mysql/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/createemp", handlers.CreateEmployee).Methods("POST")
	r.HandleFunc("/getemps", handlers.GetAllEmployee).Methods("GET")
	r.HandleFunc("/getbyid/{id}", handlers.GetEmployee).Methods("GET")
	r.HandleFunc("/updateemp/{id}", handlers.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/deleteemp/{id}", handlers.DeleteEmployee).Methods("DELETE")
	return r
}
