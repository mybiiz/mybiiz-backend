package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mybiiz/mybiiz-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Route partners
// @Summary Get all partners
// @ID partners
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Success 200 {object} Partner
// @Router /partners/ [get]
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

	r.HandleFunc("/users/{id}/view", ViewUser(db)).Methods("GET")
	r.HandleFunc("/userssave", SaveUser(db)).Methods("POST")

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
	// ShowPartners godoc
	// @Summary show partners
	// @Router /partners [get]
	// @Produce json
	r.HandleFunc("/partners", func(w http.ResponseWriter, r *http.Request) {
		var partners []Partner
		All(db, &partners, w, r)
	}).Methods("GET")

	// GetStringByInt example
	// @Summary Add a new pet to the store
	// @Description get string by ID
	// @ID get-string-by-int
	// @Accept  json
	// @Produce  json
	// @Param   some_id      path   int     true  "Some ID"
	// @Success 200 {string} string	"ok"
	// @Router /testapi/get-string-by-int/{some_id} [get]
	// @Router /accounts/{id} [get]
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
	r.HandleFunc("/partners/{id}/view", PartnerView(db)).Methods("GET")

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

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
