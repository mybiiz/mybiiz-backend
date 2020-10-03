package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Hello wolrd!")

	db := DbInit()

	r := mux.NewRouter()
	Route(r, &db)

	// handler := cors.Default().Handler(r)
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"}}).Handler(r)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", handler)

	defer db.Close()

	// r := mux.NewRouter()
}
