package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"sync"
	"ws/pkg/logger"
)

const BUF_SIZE = 1024

type Server struct {
	logger *logger.Logger
	conns  map[*websocket.Conn]bool
	mu     sync.Mutex
}

func New(logger *logger.Logger) *Server {
	return &Server{
		logger: logger,
		conns:  make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	s.logger.Info(fmt.Sprintf("new incoming connection: %s", ws.RemoteAddr()))

	/** mutex is used to prevent race conditions */
	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := ws.Read(buf)
		if err != nil {

			/** handle situation where the client closes the connection */
			if err == io.EOF {
				s.logger.Info("client disconnected socket")
				delete(s.conns, ws)
				break
			}

			/** TODO: do proper error handling */
			s.logger.Error(fmt.Sprintf("read error: %v+", err))

			/** breaking out of the loop will close the websocket */
			continue
		}

		/** echo the incoming message to all connected websockets */
		msgBytes := buf[:n]
		s.Broadcast(msgBytes)
	}
}

func (s *Server) Broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				s.logger.Warn("failed to broadcast message to websocket")
			}
		}(ws)
	}
}