package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// GENERIC CRUD FUNCTION

// All generic CRUD function
func All(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db.Find(table)
	json.NewEncoder(w).Encode(table)
}

// Get generic CRUD function
func Get(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["id"]
	db.Where("id = ?", id).First(table)
	json.NewEncoder(w).Encode(table)
}

// Post generic CRUD function
func Post(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(table)
	db.Save(table)
	json.NewEncoder(w).Encode(table)
}

// Delete generic CRUD function
func Delete(db *gorm.DB, table interface{}, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.Delete(table, id)
	json.NewEncoder(w).Encode(table)
}

// PageInfo function
func GetPageInfo(db *gorm.DB, count int64, table interface{}, r *http.Request) PageInfo {
	// var count int64
	var page int
	var perPage int

	pageInt, err := strconv.Atoi(r.FormValue("page"))

	if err != nil {
		page = 0
	} else {
		page = pageInt
	}

	perPageInt, err := strconv.Atoi(r.FormValue("perPage"))

	if err != nil {
		perPage = 10
	} else {
		perPage = perPageInt
	}

	// db.Find(table).Count(&count)
	// db.Scopes(Paginate(r)).Find(table)

	lastPage := func() float64 {
		if int64(perPage) == count {
			return 0
		} else {
			return float64(count) / float64(perPage)
		}
	}()

	return PageInfo{
		TotalElements: count,
		Last:          int64(math.Floor(lastPage)),
		Page:          page,
		PerPage:       perPage,
	}
}

// Paginate Function
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var page int
		var perPage int

		pageInt, err := strconv.Atoi(r.FormValue("page"))

		if err != nil {
			page = 0
		} else {
			// fmt.Println("Page Int", pageInt)
			page = pageInt
		}

		perPageInt, err := strconv.Atoi(r.FormValue("perPage"))

		if err != nil {
			perPage = 10
		} else {
			// fmt.Println("Per page", perPageInt)

			switch {
			case perPageInt > 100:
				perPage = 100
			case perPageInt <= 0:
				perPage = 10
			default:
				perPage = perPageInt
			}
		}

		// fmt.Println("Offset & perpage:", page, perPage)

		offset := page * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

// User

// AllUsers endpoint
// @Summary All Users
// @Tags users
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func AllUsers(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		All(db, &users, w, r)
	}
}

// UsersPaged endpoint
// @Summary All Users Paged
// @Tags users
// @Produce  json
// @Param page query int true "Page no"
// @Param perPage query int true "Items per page"
// @Success 200 {array} User
// @Router /userspaged [get]
func UsersPaged(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var users []User
		var count int64
		// db.Find(&users).Count(&count)

		dbToFind := db.Preload("Partners.Business.ServiceType").
			Preload("Partners.Bank").
			Order("id desc")
			// Scopes(Paginate(r))

		dbToFind.Find(&users).Count(&count).Scopes(Paginate(r)).Find(&users)
		page := Page{
			GetPageInfo(db, count, &users, r),
			users,
		}
		json.NewEncoder(w).Encode(&page)
	}
}

// GetUser endpoint
// @Summary Get User
// @Tags users
// @Produce  json
// @Param   id      path   int     true  "Some ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func GetUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		Get(db, &user, w, r)
	}
}

// DeleteUser endpoint
// @Summary Delete User
// @Tags users
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} User
// @Router /users/{id} [delete]
func DeleteUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		Delete(db, &user, w, r)
	}
}

// PostUser endpoint
// @Summary Post User
// @Tags users
// @Produce  json
// @Param   body      body  User     true "User"
// @Success 200 {object} User
// @Router /users/{id} [post]
func PostUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		Post(db, &user, w, r)
	}

}

// Building

// AllBuildings endpoint
// @Summary All Buildings
// @Tags buildings
// @Produce  json
// @Success 200 {array} Building
// @Router /buildings [get]
func AllBuildings(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var buildings []Building
		All(db, &buildings, w, r)
	}
}

// GetBuilding endpoint
// @Summary Get Building
// @Tags buildings
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Building
// @Router /buildings/{id} [get]
func GetBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Get(db, &building, w, r)
	}

}

// DeleteBuilding endpoint
// @Summary Delete Building
// @Tags buildings
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Building
// @Router /buildings/{id} [delete]
func DeleteBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Delete(db, &building, w, r)
	}
}

// PostBuilding endpoint
// @Summary Post Building
// @Tags buildings
// @Produce  json
// @Param   body      body  Building     true "Building"
// @Success 200 {object} Building
// @Router /buildings [post]
func PostBuilding(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var building Building
		Post(db, &building, w, r)
	}
}

// AllBanks endpoint
// @Summary All Banks
// @Tags banks
// @Produce  json
// @Success 200 {array} Bank
// @Router /banks [get]
func AllBanks(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var banks []Bank
		All(db, &banks, w, r)
	}

}

// Role

// AllRoles endpoint
// @Summary All Roles
// @Tags roles
// @Produce  json
// @Success 200 {array} Role
// @Router /roles [get]
func AllRoles(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var roles []Role
		All(db, &roles, w, r)
	}
}

// GetRole endpoint
// @Summary Get Role
// @Tags roles
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Role
// @Router /roles/{id} [get]
func GetRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var role Role
		Get(db, &role, w, r)
	}
}

// DeleteRole endpoint
// @Summary Delete Role
// @Tags roles
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Role
// @Router /roles/{id} [delete]
func DeleteRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var role Role
		Delete(db, &role, w, r)
	}
}

// PostRole endpoint
// @Summary Post Role
// @Tags roles
// @Produce  json
// @Param   body      body  Role     true "Role"
// @Success 200 {object} Role
// @Router /roles [post]
func PostRole(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var role Role
		Post(db, &role, w, r)
	}
}

// Partners

// AllPartners endpoint
// @Summary All Partners
// @Tags partners
// @Produce  json
// @Success 200 {array} Role
// @Router /partners [get]
func AllPartners(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var partners []Partner
		All(db, &partners, w, r)
	}
}

// PartnersPaged endpoint
// @Summary All Partners Paged
// @Tags partners
// @Produce  json
// @Param page query int true "Page no"
// @Param perPage query int true "Items per page"
// @Success 200 {array} Partner
// @Router /partnerspaged [get]
func PartnersPaged(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var partners []Partner
		var count int64
		// db.Find(&partners).Count(&count)

		serviceTypeId := r.FormValue("serviceTypeId")
		fmt.Println("Service type ID:", serviceTypeId)

		dbToFind := db.
			// Debug().
			Preload("User").
			Preload("Business.ServiceType").
			Preload("Bank").
			Joins("Business").
			Order("id desc")
			// Scopes(Paginate(r))

		if serviceTypeId != "0" {
			dbToFind.Where("Business.service_type_id = ?", serviceTypeId)
		}
		dbToFind.Find(&partners).Count(&count).Scopes(Paginate(r)).Find(&partners)

		page := Page{
			GetPageInfo(db, count, &partners, r),
			partners,
		}
		json.NewEncoder(w).Encode(&page)
	}
}

// PartnersExcel endpoint
// @Summary All Partners Paged excel
// @Tags partners
// @Success 200
// @Router /partnersexcel [get]
func PartnersExcel(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f := excelize.NewFile()
		index := f.NewSheet("Partners")

		f.SetCellValue("Partners", "A1", "No")
		f.SetCellValue("Partners", "B1", "Email")
		f.SetCellValue("Partners", "C1", "Name")
		f.SetCellValue("Partners", "D1", "Business Name")
		f.SetCellValue("Partners", "E1", "Phone")
		f.SetCellValue("Partners", "F1", "Bank")
		f.SetCellValue("Partners", "G1", "Bank Account #")
		f.SetCellValue("Partners", "H1", "Service Type")

		var partners []Partner
		db.
			Preload("Business.ServiceType").
			Preload("Bank").
			Preload("User").
			Find(&partners)

		for i, partner := range partners {
			excelIndex := i + 2

			f.SetCellValue("Partners", fmt.Sprintf("A%d", excelIndex), i+1)
			f.SetCellValue("Partners", fmt.Sprintf("B%d", excelIndex), partner.User.Email)
			f.SetCellValue("Partners", fmt.Sprintf("C%d", excelIndex), fmt.Sprintf("%s %s", partner.FirstName, partner.LastName))
			f.SetCellValue("Partners", fmt.Sprintf("D%d", excelIndex), partner.Business.Name)
			f.SetCellValue("Partners", fmt.Sprintf("E%d", excelIndex), partner.Phone)
			f.SetCellValue("Partners", fmt.Sprintf("F%d", excelIndex), fmt.Sprintf("%s - %s", partner.Bank.Code, partner.Bank.Name))
			f.SetCellValue("Partners", fmt.Sprintf("G%d", excelIndex), partner.BankAccountID)
			f.SetCellValue("Partners", fmt.Sprintf("H%d", excelIndex), partner.Business.ServiceType.Name)
		}

		f.SetActiveSheet(index)

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=partners.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		f.Write(w)
	}
}

// GetPartner endpoint
// @Summary Get Partner
// @Tags partners
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Partner
// @Router /partners/{id} [get]
func GetPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Get(db, &partner, w, r)
	}
}

// DeletePartner endpoint
// @Summary Delete Partner
// @Tags partners
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Partner
// @Router /partners/{id} [delete]
func DeletePartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Delete(db, &partner, w, r)
	}
}

// PostPartner endpoint
// @Summary Post Partner
// @Tags partners
// @Produce  json
// @Param   body      body  Partner     true "Partner"
// @Success 200 {object} Partner
// @Router /partners [post]
func PostPartner(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		Post(db, &partner, w, r)
	}
}

// Businesses

// AllBusinesses endpoint

// AllBusinesses endpoint
// @Summary All Businesses
// @Tags businesses
// @Produce  json
// @Success 200 {array} Business
// @Router /businesses [get]
func AllBusinesses(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var businesses []Business
		All(db, &businesses, w, r)
	}
}

// GetBusiness endpoint
// @Summary Get Business
// @Tags businesses
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /businesses/{id} [get]
func GetBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var business Business
		Get(db, &business, w, r)
	}
}

// DeleteBusiness endpoint
// @Summary Delete Business
// @Tags businesses
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /businesses/{id} [delete]
func DeleteBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var business Business
		Delete(db, &business, w, r)
	}
}

// PostBusiness endpoint
// @Summary Post Business
// @Tags businesses
// @Produce  json
// @Param   body      body  Business     true "Business"
// @Success 200 {object} Business
// @Router /businesses [post]
func PostBusiness(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var business Business
		Post(db, &business, w, r)
	}
}

// ServiceType

// AllServiceTypes endpoint
// @Summary All Service Types
// @Tags servicetypes
// @Produce  json
// @Success 200 {array} ServiceType
// @Router /servicetypes [get]
func AllServiceTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var serviceTypes []ServiceType
		All(db, &serviceTypes, w, r)
	}
}

// GetServiceType endpoint
// @Summary Get Service Types
// @Tags servicetypes
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Business
// @Router /servicetypes/{id} [get]
func GetServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var serviceType ServiceType
		Get(db, &serviceType, w, r)
	}
}

// DeleteServiceType endpoint
// @Summary Delete Service Type
// @Tags servicetypes
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ServiceType
// @Router /servicetypes/{id} [delete]
func DeleteServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var serviceType ServiceType
		Delete(db, &serviceType, w, r)
	}
}

// PostServiceType endpoint
// @Summary Post Service Type
// @Tags servicetypes
// @Produce  json
// @Param   body      body  ServiceType     true "Service Type"
// @Success 200 {object} ServiceType
// @Router /servicetype [post]
func PostServiceType(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var serviceType ServiceType
		Post(db, &serviceType, w, r)
	}
}

// Room

// AllRooms endpoint
// @Summary All Rooms
// @Tags rooms
// @Produce  json
// @Success 200 {array} Room
// @Router /rooms [get]
func AllRooms(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var rooms []Room
		All(db, &rooms, w, r)
	}
}

// GetRoom endpoint
// @Summary Get Room
// @Tags rooms
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Room
// @Router /rooms/{id} [get]
func GetRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var room Room
		Get(db, &room, w, r)
	}
}

// DeleteRoom endpoint
// @Summary Delete Room
// @Tags rooms
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} Room
// @Router /rooms/{id} [delete]
func DeleteRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var room Room
		Delete(db, &room, w, r)
	}
}

// PostRoom endpoint
// @Summary Post Room
// @Tags rooms
// @Produce  json
// @Param   body      body  Room     true "Room"
// @Success 200 {object} Room
// @Router /rooms [post]
func PostRoom(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var room Room
		Post(db, &room, w, r)
	}
}

// RoomsView endpoint
// @Summary All Rooms view with preload
// @Tags rooms
// @Produce  json
// @Success 200 {array} Room
// @Router /rooms [get]
func RoomsView(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id := mux.Vars(r)["id"]

		var rooms []Room

		builder := db.Preload("RoomType").
			Preload("FoodAccomodation").
			Preload("CancellationFee").
			Preload("GuestType")

		builder.Where("business_id = ?", id).Find(&rooms)

		json.NewEncoder(w).Encode(&rooms)
	}
}

// ComingSoonEmails

// AllComingSoonEmails endpoint
// @Summary All Coming Soon Emails
// @Tags comingsoonemails
// @Produce  json
// @Success 200 {array} ComingSoonEmail
// @Router /coming-soon-emails [get]
func AllComingSoonEmails(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmails []ComingSoonEmail
		All(db, &comingSoonEmails, w, r)
	}
}

// GetComingSoonEmail endpoint
// @Summary Get Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails/{id} [get]
func GetComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmail ComingSoonEmail
		Get(db, &comingSoonEmail, w, r)
	}
}

// DeleteComingSoonEmail endpoint
// @Summary Delete Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails/{id} [delete]
func DeleteComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmail ComingSoonEmail
		Delete(db, &comingSoonEmail, w, r)
	}
}

// PostComingSoonEmail endpoint
// @Summary Post Coming Soon Email
// @Tags comingsoonemails
// @Produce  json
// @Param   body      body  ComingSoonEmail     true "Room"
// @Success 200 {object} ComingSoonEmail
// @Router /coming-soon-emails [post]
func PostComingSoonEmail(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comingSoonEmail ComingSoonEmail
		Post(db, &comingSoonEmail, w, r)
	}
}

// Cities

// AllCities endpoint
// @Summary All Cities
// @Tags cities
// @Produce  json
// @Success 200 {array} City
// @Router /cities [get]
func AllCities(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cities []City
		All(db, &cities, w, r)
	}
}

// GetCity endpoint
// @Summary Get City
// @Tags cities
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} City
// @Router /cities/{id} [get]
func GetCity(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var city City
		Get(db, &city, w, r)
	}
}

// DeleteCity endpoint
// @Summary Delete City
// @Tags cities
// @Produce  json
// @Param   id      path   int     true  "ID"
// @Success 200 {object} City
// @Router /cities/{id} [delete]
func DeleteCity(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var city City
		Delete(db, &city, w, r)
	}
}

// PostCity endpoint
// @Summary Post City
// @Tags cities
// @Produce  json
// @Param   body      body  City     true "City"
// @Success 200 {object} City
// @Router /cities [post]
func PostCity(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var city City
		Post(db, &city, w, r)
	}
}

// CitiesPaged endpoint
// @Summary Get Cities Paged
// @Tags cities
// @Produce  json
// @Success 200 {object} City
// @Router /citiespaged [get]
func CitiesPaged(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var cities []City
		var count int64

		dbToFind := db.Order("id desc")

		dbToFind.Find(&cities).Count(&count).Scopes(Paginate(r)).Find(&cities)

		page := Page{
			GetPageInfo(db, count, &cities, r),
			cities,
		}

		json.NewEncoder(w).Encode(&page)
	}
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
			fmt.Fprintf(w, "User not found!")
			return
		}

		// fmt.Printf("JWT secret %s\n", jwtSecret)
		// fmt.Println("Found user")

		if notMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password)); notMatch != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Password does not match!")
			return
		}

		// fmt.Println("Passwrd match")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID})

		tokenString, err := token.SignedString([]byte(jwtSecret))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error generating token!")
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
// @Router /partnersregister [post]
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
// @Success 200 {array} Role
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
// @Success 200 {array} RoomType
// @Router /roomtypes [get]
func AllRoomTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var roomTypes []RoomType
		All(db, &roomTypes, w, r)
	}
}

// AllFoodAccomodations endpoint
// @Summary AllRoomTypes
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} FoodAccomodation
// @Router /foodaccomodations [get]
func AllFoodAccomodations(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var foodAccomodations []FoodAccomodation
		All(db, &foodAccomodations, w, r)
	}
}

// AllCancellationFees endpoint
// @Summary AllCancellationFees
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} CancellationFee
// @Router /cancellationfees [get]
func AllCancellationFees(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cancellationFees []CancellationFee
		All(db, &cancellationFees, w, r)
	}
}

// AllGuestTypes endpoint
// @Summary AllGuestTypes
// @Tags rooms_additions
// @Produce  json
// @Success 200 {array} GuestType
// @Router /guesttypes [get]
func AllGuestTypes(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var guestTypes []GuestType
		All(db, &guestTypes, w, r)
	}
}
