package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow any origin for WebSocket connections
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Define a struct to hold chat messages
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var messages []Message

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	log.Printf("WebSocket connection established from %s", conn.RemoteAddr())

	// Send message history to the new client
	for _, msg := range messages {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error writing WebSocket message: %v", err)
			conn.Close()
			delete(clients, conn)
			return
		}
	}

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			delete(clients, conn)
			break
		}

		// Store the message in memory
		messages = append(messages, msg)

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing WebSocket message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	port := 8000
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Starting WebSocket server on port %d...", port)
	go handleMessages() // Start the message broadcasting goroutine
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
