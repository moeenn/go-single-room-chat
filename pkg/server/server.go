package server

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

const BUF_SIZE = 1024

type Server struct {
	conns map[*websocket.Conn]bool
}

func New() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection", ws.RemoteAddr())

	/** TODO: protect conns with a mutex to prevent race conditions */
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := ws.Read(buf)
		if err != nil {

			/** handle situation where the client closes the connection */
			if err == io.EOF {
				break
			}

			/** TODO: do proper error handling */
			fmt.Println("read error", err)

			/** breaking out of the loop will close the websocket */
			continue
		}
		fmt.Printf("message: %s\n", buf[:n])

		/** echo the same message back to the client */
		ws.Write(buf[:n])
	}
}