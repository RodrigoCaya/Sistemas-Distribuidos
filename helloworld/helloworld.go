package helloworld

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Printf("Received message body from client: %s", message.Id)
	return &Message{Id: "Hello From the Server!"}, nil
}
