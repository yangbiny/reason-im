package service

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Msg  string
	Conn *websocket.Conn
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func StartWebsocketService() {
	http.HandleFunc("/ws", service)
}

func service(write http.ResponseWriter, request *http.Request) {
	var upgrade, err = (&websocket.Upgrader{}).Upgrade(write, request, nil)
	if err != nil {
		return
	}
	clients[upgrade] = true
}
