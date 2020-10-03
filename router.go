package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Route(r *mux.Router, dbPointer **gorm.DB) {
	db := *dbPointer

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	// Login & Register
	r.HandleFunc("/login", Login(db)).Methods("POST")
	r.HandleFunc("/register", Register(db)).Methods("POST")

	// Generate secure JWT secret
	r.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

		b := make([]rune, 64)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}

		fmt.Fprintf(w, "%s", string(b))
	}).Methods("GET")

	// User
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		All(db, &users, w, r)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var user User
		Get(db, &user, w, r)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var user User
		Delete(db, &user, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var user User
		Post(db, &user, w, r)
	}).Methods("POST")

	r.HandleFunc("/userssave", SaveUser(db))

	// Role

	// Building
	r.HandleFunc("/buildings", func(w http.ResponseWriter, r *http.Request) {
		var buildings []Building
		All(db, &buildings, w, r)
	}).Methods("GET")

	r.HandleFunc("/buildings/{id}", func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Get(db, &building, w, r)
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
		Get(db, &partner, w, r)
	}).Methods("GET")

	r.HandleFunc("/partners/{id}", func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Delete(db, &partner, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Post(db, &partner, w, r)
	}).Methods("POST")

	r.HandleFunc("/partnersregister", PartnerRegisterHandler(db))
	r.HandleFunc("/partnersview", PartnersView(db)).Methods("GET")

	// Businesses
	r.HandleFunc("/businesses", func(w http.ResponseWriter, r *http.Request) {
		var businesses []Business
		All(db, &businesses, w, r)
	}).Methods("GET")

	// Service Types
	r.HandleFunc("/servicetypes", func(w http.ResponseWriter, r *http.Request) {
		var serviceTypes []ServiceType
		All(db, &serviceTypes, w, r)
	}).Methods("GET")

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

	// Populator
	r.HandleFunc("/populate", Populate(db)).Methods("GET")

	r.HandleFunc("/populate/servicetypes", func(w http.ResponseWriter, r *http.Request) {
		PopulateServiceType(db)
	}).Methods("GET")

	r.HandleFunc("/populate/users", func(w http.ResponseWriter, r *http.Request) {
		PopulateUser(db)
	}).Methods("GET")
}
