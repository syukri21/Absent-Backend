package main

import (
	"backend-qrcode/db"

	socketIo "backend-qrcode/socket.io"

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

	socket := socketIo.GetSocketIO()
	socket.Run()
	go socket.Server.Serve()
	defer socket.Server.Close()

	router.Handle("/socket.io/", socket.Server)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3021", "chrome-extension://eajaahbjpnhghjcdaclbkeamlkepinbl"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"PUT", "DELETE", "POST", "GET", "PATCH"},
		// Enable Debugging for testing, consider disabling in production
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
