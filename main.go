package main

import (
	"backend-qrcode/db"
	"log"
	"net/http"
	"os"
)

func main() {

	//init router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database

	db.DB = db.SetupDB()
	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
