package main

import (
	"backend-qrcode/db"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//init router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database

	db.DB = db.SetupDB()
	// Migrate(db.DB)
	defer db.DB.Close()

	// SOCKET

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
		AllowedMethods:   []string{"PUT", "DELETE", "POST", "GET", "PATCH"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
