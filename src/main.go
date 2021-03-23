package main

import (
	"chat-server/src/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app.SetupApp()

	port := ":5000"
	fmt.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
