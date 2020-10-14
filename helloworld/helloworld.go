package helloworld

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Printf("Received message body from client: %s %s %s %s %s %s", message.Id, message.Producto, message.Valor, message.Tienda, message.Destino, message.Propietario)
	return &Message{Id: "Hello From the Server!"}, nil
}
