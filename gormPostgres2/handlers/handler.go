package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"postgres/gormPostgres2/controller"
)

func Handler() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", controller.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{employeeId}", controller.GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{employeeId}", controller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{employeeId}", controller.DeleteEmployeeById).Methods("DELETE")
	http.ListenAndServe(":2345", r)
}
