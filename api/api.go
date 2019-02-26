package api

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message is %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
