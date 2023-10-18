package service

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.onConnect(client)
		case client := <-h.unregister:
			h.onDisconnect(client)
		case message := <-h.broadcast:
			h.onBroadcast(message)
		}
	}
}

type MessageCmd struct {
	UserId   int64 `login_user_id:"user_id"`
	FriendId int64 `json:"friend_id"`
}

func (h *Hub) onConnect(c *Client) {
	h.clients[c] = true
	fmt.Println("Client connected")
}

// Handle client disconnect
func (h *Hub) onDisconnect(c *Client) {
	delete(h.clients, c)
	close(c.send)
	fmt.Println("Client disconnected")
}

// Broadcast message to all clients
func (h *Hub) onBroadcast(message []byte) {
	for c := range h.clients {
		c.send <- message
	}
}

func ServeWs(hub *Hub, cmd *MessageCmd, w http.ResponseWriter, r *http.Request) {
	/*	if cmd.UserId == cmd.FriendId {
		return
	}*/
	key := r.Header.Get("Sec-WebSocket-Key")
	guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

	fullString := fmt.Sprintf("%s%s", key, guid)
	hash := sha1.Sum([]byte(fullString))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	w.Header().Set("Sec-WebSocket-Accept", encoded)

	conn, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte)}

	client.hub.register <- client

	go client.read()
	go client.write()
}

func (c *Client) read() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.broadcast <- message
	}
	c.hub.unregister <- c
}

func (c *Client) write() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			msg := strconv.AppendInt(message, int64(rand.Int()), 10)
			c.conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
