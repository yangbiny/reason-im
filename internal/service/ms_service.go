package service

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net/http"
)

type WsService struct {
	conn *websocket.Conn
}

type Message struct {
	Msg string
}

type MSServiceCmd struct {
	UserId   int64 `login_user_id:"userId"`
	FriendId int64 `json:"friend_id"`
}

var clients = make(map[*websocket.Conn]bool)

func MSService(write http.ResponseWriter, request *http.Request, cmd *MSServiceCmd) (bool, error) {
	if cmd.UserId == cmd.FriendId {
		return false, errors.WithStack(errors.New("不能和自己 聊天"))
	}

	// 将Http 请求升级到 websocket 协议
	var conn, err = (&websocket.Upgrader{}).Upgrade(write, request, nil)
	if err != nil {
		return false, errors.WithStack(err)
	}
	clients[conn] = true

	return true, nil
}
