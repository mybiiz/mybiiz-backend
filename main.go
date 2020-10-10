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
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host mybiiz.swagger.io
// @BasePath /v2
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
}
