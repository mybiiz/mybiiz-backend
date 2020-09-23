package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func All(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db.Find(table)
	json.NewEncoder(w).Encode(table)
}

func Get(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	db.Where("id = ?", id).First(table)
	json.NewEncoder(w).Encode(table)
}

func Post(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(table)
	db.Save(table)
	json.NewEncoder(w).Encode(table)
}

func Delete(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.Delete(table, id)
	json.NewEncoder(w).Encode(table)
}

func SaveUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userPostBody UserPostBody
		json.NewDecoder(r.Body).Decode(&userPostBody)
		hash, err := bcrypt.GenerateFromPassword([]byte(userPostBody.Password), bcrypt.DefaultCost)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		db.Save(&User{
			Name:     userPostBody.Name,
			Username: userPostBody.Username,
			Email:    userPostBody.Email,
			Password: string(hash)})
	}
}
