package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mybiiz/mybiiz-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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
	r.HandleFunc("/generate", GenerateJwt(db)).Methods("GET")

	// User
	r.HandleFunc("/users", AllUsers(db)).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser(db)).Methods("GET")
	r.HandleFunc("/users/{id}", DeleteUser(db)).Methods("DELETE")
	r.HandleFunc("/users", PostUser(db)).Methods("POST")
	r.HandleFunc("/users/{id}/view", ViewUser(db)).Methods("GET")
	r.HandleFunc("/userssave", SaveUser(db)).Methods("POST")

	// Role
	r.HandleFunc("/roles", AllRoles(db)).Methods("GET")
	r.HandleFunc("/roles/{id}", GetRole(db)).Methods("GET")
	r.HandleFunc("/roles/{id}", DeleteRole(db)).Methods("DELETE")
	r.HandleFunc("/roles", PostUser(db)).Methods("POST")

	// Building
	r.HandleFunc("/buildings", AllBuildings(db)).Methods("GET")
	r.HandleFunc("/buildings/{id}", GetBuilding(db)).Methods("GET")
	r.HandleFunc("/buildings/{id}", DeleteBuilding(db)).Methods("DELETE")
	r.HandleFunc("/buildings", PostBuilding(db)).Methods("POST")

	// Partner
	r.HandleFunc("/partners", AllPartners(db)).Methods("GET")
	r.HandleFunc("/partners/{id}", GetPartner(db)).Methods("GET")
	r.HandleFunc("/partners/{id}", DeletePartner(db)).Methods("DELETE")
	r.HandleFunc("/partners", PostPartner(db)).Methods("POST")
	r.HandleFunc("/partnersregister", PartnerRegisterHandler(db)).Methods("POST")
	r.HandleFunc("/partnersview", PartnersView(db)).Methods("GET")
	r.HandleFunc("/partners/{id}/view", PartnerView(db)).Methods("GET")

	// Businesses
	r.HandleFunc("/businesses", AllBusinesses(db)).Methods("GET")
	r.HandleFunc("/businesses/{id}", GetBusiness(db)).Methods("GET")
	r.HandleFunc("/businesses/{id}", DeleteBusiness(db)).Methods("DELETE")
	r.HandleFunc("/businesses", PostBusiness(db)).Methods("POST")

	// Service Types
	r.HandleFunc("/servicetypes", AllServiceTypes(db)).Methods("GET")
	r.HandleFunc("/servicetypes/{id}", GetServiceType(db)).Methods("GET")
	r.HandleFunc("/servicetypes/{id}", DeleteServiceType(db)).Methods("DELETE")
	r.HandleFunc("/servicetypes", PostServiceType(db)).Methods("POST")

	// Room
	r.HandleFunc("/rooms", AllRooms(db)).Methods("GET")
	r.HandleFunc("/rooms/{id}", GetRoom(db)).Methods("GET")
	r.HandleFunc("/rooms/{id}", DeleteRoom(db)).Methods("DELETE")
	r.HandleFunc("/rooms", PostRoom(db)).Methods("POST")

	// ComingSoonEmail
	r.HandleFunc("/coming-soon-emails", AllComingSoonEmails(db)).Methods("GET")
	r.HandleFunc("/coming-soon-emails/{id}", GetComingSoonEmail(db)).Methods("GET")
	r.HandleFunc("/coming-soon-emails/{id}", DeleteComingSoonEmail(db)).Methods("DELETE")
	r.HandleFunc("/coming-soon-emails", PostComingSoonEmail(db)).Methods("POST")

	// Bank (read-only)
	r.HandleFunc("/banks", AllBanks(db)).Methods("GET")

	// Room types (read-only)
	r.HandleFunc("/roomtypes", AllRoomTypes(db)).Methods("GET")

	// Food Accomodation (read-only)
	r.HandleFunc("/foodaccomodations", AllFoodAccomodations(db)).Methods("GET")

	// Cancellation Fee (read-only)
	r.HandleFunc("/cancellationfees", AllCancellationFees(db)).Methods("GET")

	// Guest Type (read-only)
	r.HandleFunc("/guesttypes", AllGuestTypes(db)).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
