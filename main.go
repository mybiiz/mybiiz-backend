package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

// @title MyBiiz Backend
// @version 1.0
// @description This is the api doc for MyBiiz Backend.

// @BasePath /
func main() {
	fmt.Println("Hello wolrd!")

	db := DbInit()

	r := mux.NewRouter()
	r.Use(AuthMiddleware)
	r.Use(BrotliMiddleware)
	Route(r, &db)

	// handler := cors.Default().Handler(r)
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"}}).Handler(r)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", handler)

	// defer db.Close()
}
