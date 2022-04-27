package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ProductApp/app"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", app.CreateProduct).Methods("POST")
	s.HandleFunc("/getAllUsers", app.GetAllProduct).Methods("GET")
	s.HandleFunc("/getUserProfile", app.GetProduct).Methods("POST")
	s.HandleFunc("/updateProfile", app.UpdateProduct).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", app.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
