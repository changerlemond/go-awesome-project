package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	r := mux.NewRouter()
	tmpl := template.Must(template.ParseFiles("form.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, nil)
	}).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var details ContactDetails
		_ = json.NewDecoder(r.Body).Decode(&details) // read body, then decode
		_ = json.NewEncoder(w).Encode(details)       // encode, then send to user
	}).Methods("POST")

	_ = http.ListenAndServe(":8080", r)
}
