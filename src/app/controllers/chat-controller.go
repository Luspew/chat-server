package controllers

import (
	"chat-server/src/server"
	"net/http"
)

func ChatServerController(w http.ResponseWriter, r *http.Request) {
	server.ChatServer(w, r)
}
