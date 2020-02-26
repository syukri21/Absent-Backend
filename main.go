package main

import (
	"backend-qrcode/db"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	Migrate(db.DB)
	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
