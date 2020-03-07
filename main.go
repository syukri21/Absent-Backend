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

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
