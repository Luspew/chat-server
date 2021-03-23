package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Check request origin
func originChecker(r *http.Request) bool {
	// Allow all origins
	return true
}

func requestReader(conn *websocket.Conn) {
	for {
		msgType, buf, erro := conn.ReadMessage()
		if erro != nil {
			log.Println(erro)
		}

		log.Println(string(buf))

		if erro := conn.WriteMessage(msgType, buf); erro != nil {
			log.Fatal(erro)
		}
	}
}

func ChatServer(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = originChecker

	ws, erro := upgrader.Upgrade(w, r, nil)
	if erro != nil {
		log.Fatal(erro)
	}

	log.Println("Client connected on chat server")

	requestReader(ws)
}
