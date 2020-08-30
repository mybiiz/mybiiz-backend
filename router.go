package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Rest funcs
func Get(db *gorm.DB, table interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		// var row []Table
		db.Find(&table).Order("id desc")
		json.NewEncoder(w).Encode(table)
	}
}

func Route(r *mux.Router, dbPointer **gorm.DB) {
	db := *dbPointer

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	// User
	// Role

	// Building
	r.HandleFunc("/buildings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var buildings []Building
		db.Find(&buildings).Order("id desc")
		json.NewEncoder(w).Encode(buildings)
	}).Methods("GET")
	// r.HandleFunc("/buildings", Get(db, []Building{}))

	r.HandleFunc("/buildings/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var building Building
		db.First(&building, "id = ?", id)
		json.NewEncoder(w).Encode(building)
	}).Methods("GET")

	r.HandleFunc("/buildings/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		db.Delete(&Building{}, id)
	}).Methods("DELETE")

	r.HandleFunc("/buildings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var building Building
		json.NewDecoder(r.Body).Decode(&building)
		db.Save(&building)
		json.NewEncoder(w).Encode(building)
	}).Methods("POST")

	// Partner
	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var partners []Partner
		db.Find(&partners).Order("id desc")
		json.NewEncoder(w).Encode(partners)
	}).Methods("GET")

	r.HandleFunc("/partners/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var partner Partner
		db.First(&partner, "id = ?", id)
		json.NewEncoder(w).Encode(partner)
	}).Methods("GET")

	r.HandleFunc("/partners/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		db.Delete(&Partner{}, "id = ?", id)
	}).Methods("DELETE")

	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var partner Partner
		json.NewDecoder(r.Body).Decode(&partner)
		db.Save(&partner)
		json.NewEncoder(w).Encode(partner)
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
