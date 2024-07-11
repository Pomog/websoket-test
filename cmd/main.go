package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

const Port = ":8080"

func main() {
	server := NewSever()

	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/order-book", websocket.Handler(server.orderBook))

	/* The handler for the script.js
	fs := http.FileServer(http.Dir("."))
	http.Handle("/script.js", fs)
	*/

	log.Println("Server is running on http://localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
