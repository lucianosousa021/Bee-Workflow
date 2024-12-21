package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Room string
}

type Message struct {
	Type    string `json:"type"`    // Tipo: "offer", "answer", "candidate"
	Room    string `json:"room"`    // Sala
	Sender  string `json:"sender"`  // ID do remetente
	Payload string `json:"payload"` // Conteúdo (SDP ou ICE)
}

type Handler struct {
	Upgrader  websocket.Upgrader
	Client    *Client
	Clients   map[string]*Client
	Broadcast chan Message
}

func NewHandler() *Handler {
	clients := make(map[string]*Client) // Map de clientes conectados
	broadcast := make(chan Message)     // Canal para mensagens
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Handler{
		Upgrader:  upgrader,
		Clients:   clients,
		Broadcast: broadcast,
	}
}
