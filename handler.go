package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// GENERIC CRUD FUNCTION

// All generic CRUD function
func All(db *gorm.DB, table interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		db.Find(table)
		json.NewEncoder(w).Encode(table)
	}
}

// Get generic CRUD function
func Get(db *gorm.DB, table interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		db.Where("id = ?", id).First(table)
		json.NewEncoder(w).Encode(table)
	}
}

// Post generic CRUD function
func Post(db *gorm.DB, table interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewDecoder(r.Body).Decode(table)
		db.Save(table)
		json.NewEncoder(w).Encode(table)
	}
}

// Delete generic CRUD function
func Delete(db *gorm.DB, table interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		db.Delete(table, id)
		json.NewEncoder(w).Encode(table)
	}
}

// User

// AllUsers endpoint
// @Summary All Users
// @Produce  json
// @Success 200 {array} []User
// @Router /users [get]
func AllUsers(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var users []User
	return All(db, &users)
}

// GetUser endpoint
// @Summary Get User
// @Produce  json
// @Param   id      path   int     true  "Some ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func GetUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var user User
	return Get(db, &user)
}

// DeleteUser endpoint
// @Summary Delete User
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} User
// @Router /users/{id} [delete]
func DeleteUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var user User
	return Delete(db, &user)
}

// PostUser endpoint
// @Summary Post User
// @Produce  json
// @Param   body      body  User     true "User"
// @Success 200 {object} User
// @Router /users/{id} [post]
func PostUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var user User
	return Get(db, &user)
}

// Building

// AllBuildings endpoint
// @Summary All Buildings
// @Produce  json
// @Success 200 {array} []Building
// @Router /buildings [get]
func AllBuildings(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var buildings []Building
	return All(db, &buildings)
}

// GetBuilding endpoint
// @Summary Get Building
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Building
// @Router /buildings/{id} [get]
func GetBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var building Building
	return Get(db, &building)
}

// DeleteBuilding endpoint
// @Summary Delete Building
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Building
// @Router /buildings/{id} [delete]
func DeleteBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var building Building
	return Delete(db, &building)
}

// PostBuilding endpoint
// @Summary Post Building
// @Produce  json
// @Param   body      body  Building     true "Building"
// @Success 200 {object} Building
// @Router /buildings [post]
func PostBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var building Building
	return Post(db, &building)
}

// Role
func AllRoles(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var roles []Role
	return All(db, &roles)
}
func GetRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Get(db, &role)
}
func DeleteRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Delete(db, &role)
}
func PostRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Post(db, &role)
}

// Partners

// AllPartners endpoint
func AllPartners(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partners []Partner
	return All(db, &partners)
}
func GetPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Get(db, &partner)
}
func DeletePartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Delete(db, &partner)
}
func PostPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Post(db, &partner)
}

// Businesses

// AllBusinesses endpoint
func AllBusinesses(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var businesses []Business
	return All(db, &businesses)
}
func GetBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Get(db, &business)
}
func DeleteBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Delete(db, &business)
}
func PostBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Post(db, &business)
}

// ServiceType

// AllServiceTypes endpoint
func AllServiceTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceTypes []ServiceType
	return All(db, &serviceTypes)
}
func GetServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Get(db, &serviceType)
}
func DeleteServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Delete(db, &serviceType)
}
func PostServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Post(db, &serviceType)
}

// Room

// AllRooms endpoint
func AllRooms(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var rooms []Room
	return All(db, &rooms)
}
func GetRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Get(db, &room)
}
func DeleteRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Delete(db, &room)
}
func PostRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Post(db, &room)
}

// ComingSoonEmails

// AllComingSoonEmails endpoint
func AllComingSoonEmails(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmails []ComingSoonEmail
	return All(db, &comingSoonEmails)
}
func GetComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Get(db, &comingSoonEmail)
}
func DeleteComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Delete(db, &comingSoonEmail)
}
func PostComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Post(db, &comingSoonEmail)
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

func GenerateJwt(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

		b := make([]rune, 64)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}

		fmt.Fprintf(w, "%s", string(b))
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
			// registerPostBody.Partner.UserID = registerPostBody.User.ID

			// db.Save(&registerPostBody.Partner)

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

func PartnerView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var partner Partner
		db.Where("id = ?", id).Preload("Business.ServiceType").Preload("User").First(&partner)
		json.NewEncoder(w).Encode(&partner)
	}
}

func ViewUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		w.Header().Set("content-type", "application/json")
		var user User
		db.Where("id = ?", id).Preload("Partners.Business.ServiceType").Find(&user)
		json.NewEncoder(w).Encode(&user)
	}
}
