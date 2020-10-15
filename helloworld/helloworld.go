package helloworld

import (
	"log"
	"strconv"
	"golang.org/x/net/context"
)

type Server struct{
}

var cont int = 0

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Printf("Received message body from client: %s %s %s %s %s %s", message.Id, message.Producto, message.Valor, message.Tienda, message.Destino, message.Propietario)
	cont = cont + 1
	return &Message{Id: strconv.Itoa(cont)}, nil
}
