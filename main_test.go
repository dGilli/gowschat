package main

import (
	"net/http"
	"testing"

	"golang.org/x/net/websocket"
)

func TestWS(t *testing.T) {
    go func() {
        server := NewServer()
        http.Handle("/ws", websocket.Handler(server.handleWS))
        http.ListenAndServe(":8000", nil)
    }()

    origin := "http://localhost/"
	url := "ws://localhost:8000/ws"

    // Connect two clients to the server
    client1, err := websocket.Dial(url, "", origin)
    if err != nil {
        t.Fatalf("error dialing WebSocket server for client 1: %v", err)
    }
    defer client1.Close()

    client2, err := websocket.Dial(url, "", origin)
    if err != nil {
        t.Fatalf("error dialing WebSocket server for client 2: %v", err)
    }
    defer client2.Close()

    // Send a message from client1
    message1 := []byte("Hello from client 1")
    if err := websocket.Message.Send(client1, string(message1)); err != nil {
        t.Fatalf("error sending message from client 1: %v", err)
    }

    // Read the message from client2
    var message2 string
    if err := websocket.Message.Receive(client2, &message2); err != nil {
        t.Fatalf("error reading message from client 2: %v", err)
    }

    // Check if client2 received the correct message
    if message2 != string(message1) {
        t.Errorf("expected message '%s', got '%s'", string(message1), message2)
    }
}
