package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      false,       // Disable color
	// 	},
	// )

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
		&User{},
		&City{},
	}

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
	Populate(db)

	return db
}
