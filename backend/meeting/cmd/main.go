package main

import (
	"log"
	"meeting/handler"
	"net/http"
)

func main() {
	handler := handler.NewHandler()
	http.HandleFunc("/ws", handler.HandleConnections)
	go handler.HandleMessages()

	log.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
