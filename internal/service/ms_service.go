package service

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/utils/logger"
)

var magicGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

var UserHubs = make(map[int64]*Hub)

type Msg struct {
	ToUserId   int64  `json:"to_user_id" required:"true"`
	FromUserId int64  `json:"from_user_id" required:"true"`
	Msg        string `json:"msg" required:"true"`
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
}

type Hub struct {
	UserId     int64
	client     *Client
	register   chan *Client
	unregister chan *Client
	receive    chan []byte
	write      chan []byte
}

func NewHub(userId int64) *Hub {
	return &Hub{
		UserId:     userId,
		receive:    make(chan []byte),
		write:      make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) onConnect(client *Client) {
	UserHubs[h.UserId] = h
	h.client = client
}

func (h *Hub) onDisconnect(c *Client) {
	delete(UserHubs, h.UserId)
	close(c.hub.write)
}

func (h *Hub) onReceive(message []byte) {
	var msg Msg
	err := json.Unmarshal(message, &msg)
	if err != nil {
		logger.Error(nil, "unmarshal has failed", "err", err)
		return
	}
	id := msg.FromUserId
	hub, exist := UserHubs[id]
	if exist {
		hub.write <- message
	}

}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.onConnect(client)
		case client := <-h.unregister:
			h.onDisconnect(client)
		case message := <-h.receive:
			h.onReceive(message)
		case message := <-h.write:
			h.onWriteMsg(message)
		}
	}
}

func (h *Hub) onWriteMsg(message []byte) {
	err := h.client.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return
	}
}

type MessageCmd struct {
	UserId int64 `login_user_id:"user_id"`
}

func ServeWs(c *gin.Context, cmd *MessageCmd) (bool, *apierror.ApiError) {
	w := c.Writer
	r := c.Request
	key := r.Header.Get("Sec-WebSocket-Key")
	fullString := fmt.Sprintf("%s%s", key, magicGUID)
	hash := sha1.Sum([]byte(fullString))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	w.Header().Set("Sec-WebSocket-Accept", encoded)

	conn, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}

	hub := NewHub(cmd.UserId)
	go hub.run()

	client := &Client{hub: hub, conn: conn}

	client.hub.register <- client

	return true, nil
}

func SendMsg(fromUserId, receiverId int64, msg string) {
	hub := UserHubs[receiverId]
	if hub == nil {
		return
	}
	msgStruct := Msg{
		ToUserId:   receiverId,
		FromUserId: fromUserId,
		Msg:        msg,
	}
	marshal, _ := json.Marshal(msgStruct)
	hub.write <- marshal
}
