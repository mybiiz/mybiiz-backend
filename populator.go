package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Populate(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		PopulateServiceType(db)
		PopulateUser(db)
	}
}

func PopulateServiceType(db *gorm.DB) {
	serviceTypes := []ServiceType{
		ServiceType{Name: "Hotel"},
		ServiceType{Name: "Apartment"},
		ServiceType{Name: "Boarding House"},
		ServiceType{Name: "Laundry"},
		ServiceType{Name: "Car Rent"}}

	for _, serviceType := range serviceTypes {
		db.Save(&serviceType)
	}
}

func PopulateUser(db *gorm.DB) {
	generatedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	users := []User{
		User{Name: "Valian Masdani", Username: "valian", Email: "valianmasdani@gmail.com", Password: string(generatedPassword)},
		User{Name: "Mira Ayu", Username: "mira", Email: "miraayu@gmail.com", Password: string(generatedPassword)},
		User{Name: "Administrator", Username: "admin", Email: "admin@admin.com", Password: string(generatedPassword)}}

	for _, user := range users {
		db.Save(&user)
	}
}
