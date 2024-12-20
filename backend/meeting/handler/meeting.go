package handler

import (
	"fmt"
	"log"
	"meeting/model"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)
	mutex   = sync.RWMutex{}
)

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Printf("Nova conexão recebida de: %s", r.RemoteAddr)

	ws, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro no upgrade: %v", err)
		return
	}
	defer ws.Close()

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	log.Printf("Cliente conectado: %s", ws.RemoteAddr())

	for {
		log.Printf("Aguardando mensagem do cliente...")
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}

		log.Printf("Mensagem recebida - Tipo: %d, Tamanho: %d bytes", messageType, len(message))

		if messageType == websocket.BinaryMessage {
			// log.Printf("Tipo da mensagem recebida: %d", messageType)
			// log.Printf("Recebido frame de vídeo de tamanho: %d bytes", len(message))
			// log.Printf("Primeiros 50 bytes do frame: %v", message[:min(50, len(message))])

			response := model.VideoMessage{
				Type: "processing_status",
				Data: fmt.Sprintf("Frame processado com sucesso. Tamanho: %d bytes", len(message)),
			}

			if err := ws.WriteJSON(response); err != nil {
				log.Printf("Erro ao enviar confirmação: %v", err)
			}
		} else {
			log.Printf("Mensagem recebida de tipo diferente: %d", messageType)
		}
	}
}
