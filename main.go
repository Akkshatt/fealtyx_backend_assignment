package main

import (
	"log"
	"net/http"
	"fealtyx_backend_assignment/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id:[0-9]+}", handlers.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id:[0-9]+}", handlers.UpdateStudentByID).Methods("PUT")
	r.HandleFunc("/students/{id:[0-9]+}", handlers.DeleteStudentByID).Methods("DELETE")
	r.HandleFunc("/students/{id:[0-9]+}/summary", handlers.GenerateStudentSummary).Methods("GET")

	
	log.Fatal(http.ListenAndServe(":8080", r))
}
