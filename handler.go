package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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
	// w.WriteHeader(http.StatusCreated)
}

func Delete(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.Delete(table, id)
	json.NewEncoder(w).Encode(table)
}

func Login(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
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

func Register(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var registerPostBody RegisterPostBody
		json.NewDecoder(r.Body).Decode(&registerPostBody)

		// fmt.Println(registerPostBody)

		// Find existing user
		var user User
		if res := db.Where("email = ?", registerPostBody.User.Email).First(&user); res.Error != nil { // no users found
			generatedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerPostBody.User.Password), bcrypt.DefaultCost)
			registerPostBody.User.Password = string(generatedPassword)

			db.Save(&registerPostBody.User)
			registerPostBody.Partner.UserID = registerPostBody.User.ID

			db.Save(&registerPostBody.Partner)

		} else { // existing user found
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("User dengan email ini sudah terdaftar."))
		}
	}
}

func PartnerRegisterHandler(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var partnerRegister PartnerRegister
		json.NewDecoder(r.Body).Decode(&partnerRegister)
		// fmt.Println(partnerRegister)

		// Remove data:image/jpeg;base64,
		b64Split := strings.Split(partnerRegister.CitizenIDPhoto, "base64,")

		// fmt.Println(b64Split)

		imgDec, err := base64.StdEncoding.DecodeString(b64Split[1])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed converting citizen ID Photo.")
			return
		}

		f, err := os.Create(fmt.Sprintf("img/user_%d.jpg", partnerRegister.Partner.UserID))
		defer f.Close()

		if err != nil {
			fmt.Println("Error opening file.")
		}

		f.Write(imgDec)
		f.Sync()

		// Save business first
		if res := db.Save(&partnerRegister.Partner.Business); res.Error != nil {
			fmt.Println("Failed saving business")
			return
		}

		partnerRegister.Partner.BusinessID = partnerRegister.Partner.Business.ID
		fmt.Println("Saved business ID:", partnerRegister.Partner.Business.ID)
		fmt.Println("User ID:", partnerRegister.Partner.UserID)

		// Modify user registration completed
		var user User
		if res := db.Where("id = ?", partnerRegister.Partner.UserID).First(&user); res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", res.Error)
			return
		}

		user.RegistrationCompleted = true
		db.Save(&user)

		db.Save(&partnerRegister.Partner)
	}
}

func PartnersView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var partners []Partner
		db.Preload("Business.ServiceType").Preload("User").Find(&partners)
		json.NewEncoder(w).Encode(&partners)
	}
}
