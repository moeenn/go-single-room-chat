package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"ws/pkg/server"
)

func main() {
	s := server.New()
	http.Handle("/ws", websocket.Handler(s.HandleWS))
	http.ListenAndServe(":5000", nil)
}