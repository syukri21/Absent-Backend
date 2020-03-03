package main

import (
	"backend-qrcode/db"
	"backend-qrcode/socket"

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
	// Migrate(db.DB)
	defer db.DB.Close()

	// SOCKET
	webSocket, err := socket.NewSocket()

	if err != nil {
		println("ewrrr")
		log.Fatal(err)
	}

	webSocket.Listen()
	go webSocket.Serve()
	defer webSocket.Close()

	//create http server
	http.Handle("/socket.io/", webSocket.Server)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
