package main

import (
	"chat-server/src/app"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	app.SetupApp()

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please, set a port number.")
	}
	port := ":" + arguments[1]
	fmt.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
