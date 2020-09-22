package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Route(r *mux.Router, dbPointer **gorm.DB) {
	db := *dbPointer

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	// User
	// Role

	// Building
	r.HandleFunc("/buildings", func(w http.ResponseWriter, r *http.Request) {
		var buildings []Building
		All(db, &buildings, w, r)
	}).Methods("GET")

	r.HandleFunc("/buildings/{id}", func(w http.ResponseWriter, r *http.Request) {
		var building Building
		All(db, &building, w, r)
	}).Methods("GET")

	r.HandleFunc("/buildings/{id}", func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Delete(db, &building, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/buildings", func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Post(db, &building, w, r)
	}).Methods("POST")

	// Partner
	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		var partners []Partner
		All(db, &partners, w, r)
	}).Methods("GET")

	r.HandleFunc("/partners/{id}", func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		All(db, &partner, w, r)
	}).Methods("GET")

	r.HandleFunc("/partners/{id}", func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Delete(db, &partner, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Post(db, &partner, w, r)
	}).Methods("POST")

	// Room

	// r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	var users []User
	// 	db.Find(&users)

	// 	json.NewEncoder(w).Encode(users)
	// })

	// r.HandleFunc("/testadd", func(w http.ResponseWriter, r *http.Request) {
	// 	user := User{
	// 		Name:     "Valian",
	// 		Username: "valian",
	// 		Password: "mypassword"}

	// 	db.Save(&user)
	// })

	// ComingSoonEmail
	r.HandleFunc("/coming-soon-emails", func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmails []ComingSoonEmail
		All(db, &comingSoonEmails, w, r)
	}).Methods("GET")

	r.HandleFunc("/coming-soon-emails", func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmail ComingSoonEmail
		Post(db, &comingSoonEmail, w, r)
	}).Methods("POST")

	r.HandleFunc("/populate", func(w http.ResponseWriter, r *http.Request) {
		var buildings []Building
		buildings = []Building{
			Building{
				Name:    "Kost A",
				Address: "Kost A address",
				Phone:   "62xxx-xxxx-xxxx",
				Lat:     "testLat",
				Lon:     "testLon"},
			Building{
				Name:    "Kost B",
				Address: "Kost B address",
				Phone:   "62xxx-xxxx-xxxx",
				Lat:     "testLat",
				Lon:     "testLon"},
			Building{
				Name:    "Kost C",
				Address: "Kost C address",
				Phone:   "62xxx-xxxx-xxxx",
				Lat:     "testLat",
				Lon:     "testLon"},
			Building{
				Name:    "Kost D",
				Address: "Kost D address",
				Phone:   "62xxx-xxxx-xxxx",
				Lat:     "testLat",
				Lon:     "testLon"}}

		for _, building := range buildings {
			db.Save(&building)
		}
	}).Methods("GET")
}
