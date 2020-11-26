package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BankData struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func Populate(db *gorm.DB) {
	// Populate Service Types
	fmt.Println("Checking service types...")
	serviceTypes := []string{
		"Hotel",
		"Apartemen",
		"Kost",
		"Laundry",
		"Rental Mobil"}

	for _, serviceTypeName := range serviceTypes {
		var serviceType ServiceType
		if res := db.Where("name = ?", serviceTypeName).First(&serviceType); res.Error != nil {
			fmt.Println(serviceTypeName, "not found! Creating...")
			db.Save(&ServiceType{Name: serviceTypeName})
		}
	}

	// Populate admin login
	fmt.Println("Checking administrator user...")
	var user User
	if res := db.Where("email = ?", "admin@admin.com").First(&user); res.Error != nil {
		fmt.Println("Admin user not found! Creating...")
		regularPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
		db.Save(&User{
			Email:    "admin@admin.com",
			Password: string(regularPassword)})
	}

	// Populate roomTypes
	fmt.Println("Populating room types...")
	roomTypes := []RoomType{
		RoomType{Name: "Single"},
		RoomType{Name: "Double"},
		RoomType{Name: "Double-Deluxe"},
		RoomType{Name: "Twin"},
		RoomType{Name: "Suite Family"},
		RoomType{Name: "King/Superior"}}

	for _, roomType := range roomTypes {
		var foundRoomType RoomType
		if res := db.Where("name = ?", roomType.Name).First(&foundRoomType); res.Error != nil {
			fmt.Println(roomType.Name, "not found! Creating...")
			db.Save(&roomType)
		}
	}

	// Populate FoodAccomodation
	fmt.Println("Populating food accomodations...")
	foodAccomodations := []FoodAccomodation{
		FoodAccomodation{Name: "Gratis Sarapan"},
		FoodAccomodation{Name: "Gratis Dinner"},
		FoodAccomodation{Name: "Gratis Sarapan & Dinner"},
		FoodAccomodation{Name: "Gratis Jamuan"},
		FoodAccomodation{Name: "Tidak ada makanan"}}

	for _, foodAccomodation := range foodAccomodations {
		var foundFoodAccomodation FoodAccomodation
		if res := db.Where("name = ?", foodAccomodation.Name).First(&foundFoodAccomodation); res.Error != nil {
			fmt.Println(foodAccomodation.Name, "not found! Creating...")
			db.Save(&foodAccomodation)
		}
	}

	// Populate CancellationFee
	fmt.Println("Populating cancellation fee...")
	cancellationFees := []CancellationFee{
		CancellationFee{Name: "Gratis"},
		CancellationFee{Name: "10% sebelum 24 jam"},
		CancellationFee{Name: "Tidak Bisa Cancel"}}

	for _, cancellationFee := range cancellationFees {
		var foundCancellationFee CancellationFee
		if res := db.Where("name = ?", cancellationFee.Name).First(&foundCancellationFee); res.Error != nil {
			fmt.Println(cancellationFee.Name, "not found! Creating...")
			db.Save(&cancellationFee)
		}
	}

	// Populate GuestType
	fmt.Println("Populating guest type...")
	guestTypes := []GuestType{
		GuestType{Name: "Perempuan"},
		GuestType{Name: "Laki-Laki"},
		GuestType{Name: "Campur"},
		GuestType{Name: "Tidak Ada Ketentuan"}}

	for _, guestType := range guestTypes {
		var foundGuestType GuestType
		if res := db.Where("name = ?", guestType.Name).First(&foundGuestType); res.Error != nil {
			fmt.Println(guestType.Name, "not found! Creating...")
			db.Save(&guestType)
		}
	}

	// Populate banks
	fmt.Println("Populating banks...")

	bankFile, err := os.Open("data/bank.json")
	defer bankFile.Close()

	if err != nil {
		panic(err)
	}

	bankJSON, err := ioutil.ReadAll(bankFile)

	if err != nil {
		panic(err)
	}

	var banks []Bank
	json.Unmarshal(bankJSON, &banks)

	for _, bankData := range banks {
		var foundBank Bank
		if res := db.Where("name = ?", bankData.Name).First(&foundBank); res.Error != nil {
			fmt.Println(bankData.Name, "not found! Creating...")
			db.Save(&bankData)
		}
	}

	// Populate Cities
	fmt.Println("Populating cities...")

	cities := []City{
		{Name: "Surabaya", Lat: -7.2754438, Lon: 112.6426429},
		{Name: "Jakarta", Lat: -6.2293867, Lon: 106.6894301},
		{Name: "Malang", Lat: -7.9784695, Lon: 112.5617421},
		{Name: "Bali", Lat: -8.4543385, Lon: 114.5110407},
	}

	for _, city := range cities {
		var foundCity City
		if res := db.Where("name = ?", city.Name).First(&foundCity); res.Error != nil {
			fmt.Println("City", city.Name, "not found! Creating...")
			db.Save(&city)
		} else {
			fmt.Println("City", city.Name, "found!")
		}
	}

	fmt.Println("Population successful!")
}
