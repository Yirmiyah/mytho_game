package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type client struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *client) readLoop() {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		log.Printf("received message: %s", message)
		forwardMessage(message)
	}
}

func (c *client) writeLoop() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}

var clients = make(map[*client]bool)
var broadcast = make(chan []byte)

func forwardMessage(message []byte) {
	broadcast <- message
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &client{
		conn: conn,
		send: make(chan []byte),
	}

	clients[c] = true

	go c.readLoop()
	go c.writeLoop()
}

func broadcastLoop() {
	for {
		message := <-broadcast

		for c := range clients {
			select {
			case c.send <- message:
			default:
				delete(clients, c)
				close(c.send)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	go broadcastLoop()

	log.Println("Listening on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
