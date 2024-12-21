package handler

import (
	"log"
	"net/http"
)

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade para WebSocket
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Identificação do cliente
	clientID := r.URL.Query().Get("id")
	roomID := r.URL.Query().Get("room")

	client := &Client{
		ID:   clientID,
		Conn: conn,
		Room: roomID,
	}
	h.Clients[clientID] = client

	log.Printf("Client %s connected to room %s", clientID, roomID)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Client %s disconnected", clientID)
			delete(h.Clients, clientID)
			break
		}
		// Encaminha a mensagem para o canal broadcast
		h.Broadcast <- msg
	}
}

func (h *Handler) HandleMessages() {
	for {
		msg := <-h.Broadcast
		// Enviar mensagem para todos os clientes na mesma sala
		for _, client := range h.Clients {
			if client.Room == msg.Room && client.ID != msg.Sender {
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("Erro ao enviar mensagem para %s: %v", client.ID, err)
					client.Conn.Close()
					delete(h.Clients, client.ID)
				}
			}
		}
	}
}
