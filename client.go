package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func startClient() {
	// Connect to the WebSocket server
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send a message to the server
	message := []byte("Hello, server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal(err)
	}

	// Listen for messages from the server
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Client received message: %s\n", p)

		time.Sleep(time.Second)

		// Send another message to the server
		message = []byte("Thanks for echoing back my message!")
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Fatal(err)
		}
	}
}
