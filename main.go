package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"ws/pkg/logger"
	"ws/pkg/server"
)

const (
	HOST = "0.0.0.0"
	PORT = ":5000"
)

func main() {
	log := logger.New()
	s := server.New(log)

	http.Handle("/ws", websocket.Handler(s.HandleWS))

	log.Info("starting websocket server on port " + PORT)
	http.ListenAndServe(HOST+PORT, nil)
}
