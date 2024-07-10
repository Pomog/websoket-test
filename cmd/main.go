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

	/* Use index.html and script.js to sent message on websocket
	fs := http.FileServer(http.Dir("."))
	http.Handle("/script.js", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set CSP header to allow inline scripts and WebSocket connections
		w.Header().Set("Content-Security-Policy", "default-src 'self'; connect-src 'self' ws://localhost:8080; script-src 'self'")
		http.ServeFile(w, r, "index.html")
	})
	*/

	log.Println("Server is running on http://localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
