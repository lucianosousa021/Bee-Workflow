package main

import (
	"log"
	"meeting/handler"
	"net/http"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	handler := handler.NewHandler()

	http.HandleFunc("/ws", handler.HandleConnections)

	port := ":8080"
	log.Printf("Servidor iniciado em http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
