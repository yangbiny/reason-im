package service

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net/http"
)

type Message struct {
	Msg  string
	Conn *websocket.Conn
}

type MSServiceCmd struct {
	UserId   int64 `login_user_id:"userId"`
	FriendId int64 `json:"friend_id"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func MSService(write http.ResponseWriter, request *http.Request, cmd *MSServiceCmd) (bool, error) {
	if cmd.UserId == cmd.FriendId {
		return false, errors.WithStack(errors.New("不能和自己 聊天"))
	}

	var upgrade, err = (&websocket.Upgrader{}).Upgrade(write, request, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	clients[upgrade] = true
	return true, nil
}
