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
// @Tags users
// @Produce  json
// @Success 200 {array} []User
// @Router /users [get]
func AllUsers(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var users []User
	return All(db, &users)
}

// GetUser endpoint
// @Summary Get User
// @Tags users
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
// @Tags users
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
// @Tags users
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
// @Tags buildings
// @Produce  json
// @Success 200 {array} []Building
// @Router /buildings [get]
func AllBuildings(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var buildings []Building
	return All(db, &buildings)
}

// GetBuilding endpoint
// @Summary Get Building
// @Tags buildings
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
// @Tags buildings
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
// @Tags buildings
// @Produce  json
// @Param   body      body  Building     true "Building"
// @Success 200 {object} Building
// @Router /buildings [post]
func PostBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var building Building
	return Post(db, &building)
}

// AllBanks endpoint
// @Summary All Banks
// @Tags banks
// @Produce  json
// @Success 200 {array} []Bank
// @Router /banks [get]
func AllBanks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var banks []Bank
	return All(db, &banks)
}

// Role

// AllRoles endpoint
// @Summary All Roles
// @Tags roles
// @Produce  json
// @Success 200 {array} []Role
// @Router /roles [get]
func AllRoles(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var roles []Role
	return All(db, &roles)
}

// GetRole endpoint
// @Summary Get Role
// @Tags roles
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Role
// @Router /roles/{id} [get]
func GetRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Get(db, &role)
}

// DeleteRole endpoint
// @Summary Delete Role
// @Tags roles
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Role
// @Router /roles/{id} [delete]
func DeleteRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Delete(db, &role)
}

// PostRole endpoint
// @Summary Post Role
// @Tags roles
// @Produce  json
// @Param   body      body  Role     true "Role"
// @Success 200 {object} Role
// @Router /roles [post]
func PostRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var role Role
	return Post(db, &role)
}

// Partners

// AllPartners endpoint
// @Summary All Partners
// @Tags partners
// @Produce  json
// @Success 200 {array} []Role
// @Router /partners [get]
func AllPartners(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partners []Partner
	return All(db, &partners)
}

// GetPartner endpoint
// @Summary Get Partner
// @Tags partners
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Partner
// @Router /partners/{id} [get]
func GetPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Get(db, &partner)
}

// DeletePartner endpoint
// @Summary Delete Partner
// @Tags partners
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Partner
// @Router /partners/{id} [delete]
func DeletePartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Delete(db, &partner)
}

// PostPartner endpoint
// @Summary Post Partner
// @Tags partners
// @Produce  json
// @Param   body      body  Partner     true "Partner"
// @Success 200 {object} Partner
// @Router /partners [post]
func PostPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	return Post(db, &partner)
}

// Businesses

// AllBusinesses endpoint

// AllBusinesses endpoint
// @Summary All Businesses
// @Tags businesses
// @Produce  json
// @Success 200 {array} []Business
// @Router /businesses [get]
func AllBusinesses(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var businesses []Business
	return All(db, &businesses)
}

// GetBusiness endpoint
// @Summary Get Business
// @Tags businesses
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /businesses/{id} [get]
func GetBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Get(db, &business)
}

// DeleteBusiness endpoint
// @Summary Delete Business
// @Tags businesses
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /businesses/{id} [delete]
func DeleteBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Delete(db, &business)
}

// PostBusiness endpoint
// @Summary Post Business
// @Tags businesses
// @Produce  json
// @Param   body      body  Business     true "Business"
// @Success 200 {object} Business
// @Router /businesses [post]
func PostBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var business Business
	return Post(db, &business)
}

// ServiceType

// AllServiceTypes endpoint
// @Summary All Service Types
// @Tags servicetypes
// @Produce  json
// @Success 200 {array} []ServiceType
// @Router /servicetypes [get]
func AllServiceTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceTypes []ServiceType
	return All(db, &serviceTypes)
}

// GetServiceType endpoint
// @Summary Get Service Types
// @Tags servicetypes
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /servicetypes/{id} [get]
func GetServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Get(db, &serviceType)
}

// DeleteServiceType endpoint
// @Summary Delete Service Type
// @Tags servicetypes
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ServiceType
// @Router /servicetypes/{id} [delete]
func DeleteServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Delete(db, &serviceType)
}

// PostServiceType endpoint
// @Summary Post Service Type
// @Tags servicetypes
// @Produce  json
// @Param   body      body  ServiceType     true "Service Type"
// @Success 200 {object} ServiceType
// @Router /servicetype [post]
func PostServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var serviceType ServiceType
	return Post(db, &serviceType)
}

// Room

// AllRooms endpoint
// @Summary All Rooms
// @Tags rooms
// @Produce  json
// @Success 200 {array} []Room
// @Router /rooms [get]
func AllRooms(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var rooms []Room
	return All(db, &rooms)
}

// GetRoom endpoint
// @Summary Get Room
// @Tags rooms
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Room
// @Router /rooms/{id} [get]
func GetRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Get(db, &room)
}

// DeleteRoom endpoint
// @Summary Delete Room
// @Tags rooms
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Room
// @Router /rooms/{id} [delete]
func DeleteRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Delete(db, &room)
}

// PostRoom endpoint
// @Summary Post Room
// @Tags rooms
// @Produce  json
// @Param   body      body  Room     true "Room"
// @Success 200 {object} Room
// @Router /rooms [post]
func PostRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var room Room
	return Post(db, &room)
}

// ComingSoonEmails

// AllComingSoonEmails endpoint
// @Summary All Coming Soon Emails
// @Tags comingsoonemails
// @Produce  json
// @Success 200 {array} []ComingSoonEmail
// @Router /coming-soon-emails [get]
func AllComingSoonEmails(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmails []ComingSoonEmail
	return All(db, &comingSoonEmails)
}

// GetComingSoonEmail endpoint
// @Summary Get Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails/{id} [get]
func GetComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Get(db, &comingSoonEmail)
}

// DeleteComingSoonEmail endpoint
// @Summary Delete Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails/{id} [delete]
func DeleteComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Delete(db, &comingSoonEmail)
}

// PostComingSoonEmail endpoint
// @Summary Post Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   body      body  ComingSoonEmail     true "Room"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails [post]
func PostComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var comingSoonEmail ComingSoonEmail
	return Post(db, &comingSoonEmail)
}

// Login endpoint
// @Summary Login
// @Tags auth
// @Param   body      body  LoginBody     true "Login body"
// @Success 200 {string} string "jwt"
// @Router /login [post]
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

// GenerateJwt endpoint
// @Summary Generate secure JWT key
// @Tags auth
// @Success 200 {string} string	"ok"
// @Router /generate [get]
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

// SaveUser endpoint
// @Summary Save User
// @Tags users
// @Param   body      body  UserPostBody     true "UserPostBody"
// @Success 200 {object} string "ok"
// @Router /userssave [post]
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

// Register endpoint
// @Summary Register
// @Tags auth
// @Param   body      body  RegisterPostBody     true "RegisterPostBody"
// @Success 200 {string} string "ok"
// @Router /register [post]
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

// PartnerRegisterHandler endpoint
// @Summary PartnerRegisterHandler
// @Tags partners
// @Param   body      body PartnerRegister     true "PartnerRegister body"
// @Success 200 {string} string "ok"
// @Router /login [post]
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

// PartnersView endpoint
// @Summary All Roles
// @Tags partners
// @Produce  json
// @Success 200 {array} []Role
// @Router /partnersview [get]
func PartnersView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var partners []Partner
		db.Preload("Business.ServiceType").Preload("User").Find(&partners)
		json.NewEncoder(w).Encode(&partners)
	}
}

// PartnerView endpoint
// @Summary Get Partner View
// @Produce  json
// @Tags partners
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Partner
// @Router /partners/{id}/view [get]
func PartnerView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]
		var partner Partner
		db.Where("id = ?", id).Preload("Business.ServiceType").Preload("User").First(&partner)
		json.NewEncoder(w).Encode(&partner)
	}
}

// ViewUser endpoint
// @Summary Get User view
// @Tags users
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} User
// @Router /users/{id}/view [get]
func ViewUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		w.Header().Set("content-type", "application/json")
		var user User
		db.Where("id = ?", id).Preload("Partners.Business.ServiceType").Preload("Partners.Business.Rooms").Preload("Partners.Bank").Find(&user)
		json.NewEncoder(w).Encode(&user)
	}
}

// AllRoomTypes endpoint
// @Summary AllRoomTypes
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} []RoomType
// @Router /roomtypes [get]
func AllRoomTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var roomTypes []RoomType
	return All(db, &roomTypes)
}

// AllFoodAccomodations endpoint
// @Summary AllRoomTypes
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} []FoodAccomodation
// @Router /foodaccomodations [get]
func AllFoodAccomodations(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var foodAccomodations []FoodAccomodation
	return All(db, &foodAccomodations)
}

// AllCancellationFees endpoint
// @Summary AllCancellationFees
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} []CancellationFee
// @Router /cancellationfees [get]
func AllCancellationFees(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var cancellationFees []CancellationFee
	return All(db, &cancellationFees)
}

// AllGuestTypes endpoint
// @Summary AllGuestTypes
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} []GuestType
// @Router /guesttypes [get]
func AllGuestTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	var guestTypes []GuestType
	return All(db, &guestTypes)
}
