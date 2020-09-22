package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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

	return db
}
