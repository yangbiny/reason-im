package service

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"reason-im/internal/utils/logger"
	"strconv"
)

var UserHubs = make(map[int64]*Hub)

type Msg struct {
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Msg        string `json:"msg"`
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

type Hub struct {
	FromUserId int64
	ToUserId   int64
	client     *Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub(toUserId, fromUserId int64) *Hub {
	return &Hub{
		ToUserId:   toUserId,
		FromUserId: fromUserId,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) onConnect(client *Client) {
	UserHubs[h.FromUserId] = h
	h.client = client
}

func (h *Hub) onDisconnect(c *Client) {
	delete(UserHubs, h.FromUserId)
	close(c.send)
}

func (h *Hub) onBroadcast(message []byte) {
	h.client.send <- message
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
	FriendId int64 `form:"friend_id"`
}

func ServeWs(cmd *MessageCmd, c *gin.Context) {
	if cmd.UserId == cmd.FriendId {
		c.AbortWithStatus(500)
		return
	}
	w := c.Writer
	r := c.Request
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

	hub := NewHub(cmd.FriendId, cmd.UserId)
	go hub.Run()

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
		var msg Msg
		err = json.Unmarshal(message, &msg)
		if err != nil {
			logger.Error(nil, "unmarshal has failed", "err", err)
			continue
		}
		toUserId := msg.ToUserId
		hub, exist := UserHubs[toUserId]
		if exist {
			hub.broadcast <- message
		}
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
