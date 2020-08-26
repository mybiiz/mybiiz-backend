package main

import (
  "github.com/joho/godotenv"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/jinzhu/gorm"
  "log"
  "os"

  "fmt"
  "github.com/gorilla/mux"
  "net/http"
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

  r := mux.NewRouter()
  Route(r)

  fmt.Println("Listening on port 8080")
  http.ListenAndServe(":8080", r)

  defer db.Close()

  // r := mux.NewRouter()
}
