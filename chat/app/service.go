package app

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerChat interface {
	Connect(ctx echo.Context) error
}

type service struct {
}

func CreateInstance() HandlerChat {
	return service{}
}

func (s service) Connect(ctx echo.Context) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	conn.LocalAddr()
	serv := NewServer("/ws")
	go serv.Listen()
	return nil
}
