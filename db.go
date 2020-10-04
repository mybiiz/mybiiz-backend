package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func DbInit() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName))

	if err != nil {
		log.Fatal("Error openign database")
	}

	tables := []interface{}{
		&Partner{},
		&Role{},
		&Business{},
		&Building{},
		&Room{},
		&GuestType{},
		&CancellationFee{},
		&RoomType{},
		&FoodAccomodation{},
		&ComingSoonEmail{},
		&Bank{},
		&Car{},
		&CarType{},
		&FuelType{},
		&Laundry{},
		&LaundryType{},
		&ServiceType{},
		&User{}}

	for _, table := range tables {
		db.AutoMigrate(table)
	}

	// db.AutoMigrate(&Partner{})
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Role{})
	// db.AutoMigrate(&Building{})
	// db.AutoMigrate(&Room{})
	// db.AutoMigrate(&ComingSoonEmail{})
	// db.AutoMigrate(&Bank{})

	// db.Save(&Partner{
	// 	Name: "My Partner"})
	checkAndPopulate(db)

	return db
}

func checkAndPopulate(db *gorm.DB) {
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

	fmt.Println("Population successful!")
}
