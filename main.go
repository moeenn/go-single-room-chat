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
	lg := logger.New()
	s := server.New(lg)

	http.Handle("/ws", websocket.Handler(s.HandleWS))

	lg.Info("starting websocket server on port " + PORT)
	http.ListenAndServe(HOST+PORT, nil)
}
