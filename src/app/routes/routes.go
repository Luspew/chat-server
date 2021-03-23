package routes

import (
	"chat-server/src/app/controllers"
	"net/http"
)

// App Routes
func Routes() {
	http.HandleFunc("/chat", controllers.ChatServerController)
}
