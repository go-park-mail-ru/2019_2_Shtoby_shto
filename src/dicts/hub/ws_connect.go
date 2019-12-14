package hub

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type handler struct {
	hub *Hub
}

func NewWsHandler(
	e *echo.Echo,
	hub *Hub,
) {
	handler := handler{
		hub: hub,
	}
	e.GET("/cards/ws", handler.ws)
}

func (h *handler) ws(ctx echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	client := &Client{Hub: h.hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	client.ReadPump()
	return nil
}
