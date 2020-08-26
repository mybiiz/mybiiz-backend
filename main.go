package main

import (
  "github.com/joho/godotenv"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/jinzhu/gorm"
  "log"
  "os"

  "fmt"
  // "github.com/gorilla/mux"
)

func main() {
  fmt.Println("Hello wolrd!")

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
    log.Fatal(err)
  }

  defer db.Close()

  // r := mux.NewRouter()
}
