package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Route(r *mux.Router, dbPointer **gorm.DB) {
	db := *dbPointer

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	// Login
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Load secret key
		err := godotenv.Load()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		var loginBody LoginBody
		json.NewDecoder(r.Body).Decode(&loginBody)

		// fmt.Println("Decode body success")

		var user User
		if res := db.Where("email = ?", loginBody.Email).First(&user); res.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// fmt.Printf("JWT secret %s\n", jwtSecret)
		// fmt.Println("Found user")

		if notMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password)); notMatch != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// fmt.Println("Passwrd match")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID})

		tokenString, err := token.SignedString([]byte(jwtSecret))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// fmt.Println("Token generated")
		// fmt.Println(tokenString)

		fmt.Fprintf(w, "%s", tokenString)
	}).Methods("POST")

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
