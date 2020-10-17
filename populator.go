package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

	fmt.Println("Population successful!")
}
