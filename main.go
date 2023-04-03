package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"ws/pkg/log"
	"ws/pkg/server"
)

const (
	HOST = "0.0.0.0"
	PORT = ":5000"
)

func main() {
	logger := log.New()
	s := server.New(logger)

	http.Handle("/ws", websocket.Handler(s.HandleWS))

	logger.Log(log.INFO, "starting websocket server on port "+PORT)
	http.ListenAndServe(HOST+PORT, nil)
}