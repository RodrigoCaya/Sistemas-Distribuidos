package helloworld

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) SayAnswer(ctx context.Context, respuesta *Respuesta) (*Respuesta, error){
	log.Printf("Received message body from client: %s", respuesta.Body)
	return &Respuesta{Body: "Hello From the Server!"}, nil
}
