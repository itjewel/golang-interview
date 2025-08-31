package main

import (
	"fmt"
	"golang-interview/database"
	"golang-interview/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded, using system environment variables")
	}

	database.Connect() // DB connect
	defer database.DB.Close()

	mux := http.NewServeMux()
	routes.RestaurantRoutes(mux)

	fmt.Println("Server running at http://localhost:8000")
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}

}
