package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	Upgrader websocket.Upgrader
}

func NewHandler() *Handler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Handler{
		Upgrader: upgrader,
	}
}
