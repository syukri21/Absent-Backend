package main

import (
	"backend-qrcode/db"
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p), messageType)

	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)

}

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

	router.HandleFunc("/ws", handleWS)

	// SOCKET

	log.Fatal(http.ListenAndServe(":"+port, router))
}
