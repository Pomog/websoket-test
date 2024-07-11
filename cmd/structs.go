package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"time"
)

type Server struct {
	connections map[*websocket.Conn]bool
}

func NewSever() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection form: ", ws.RemoteAddr())
	s.connections[ws] = true

	defer func() {
		delete(s.connections, ws)
		err := ws.Close()
		if err != nil {
			return
		}
		fmt.Println("Connection closed: ", ws.RemoteAddr())
	}()

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected: ", ws.RemoteAddr())
				break
			}
			fmt.Println("read error: ", err)
			continue
		}
		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.connections {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}

func (s *Server) orderBook(ws *websocket.Conn) {
	fmt.Println("New connection to orderBook form: ", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderBook data ->  %s\n", time.Now().UTC())
		ws.Write([]byte(payload))
		time.Sleep(time.Second)
	}
}
