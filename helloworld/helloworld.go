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
	log.Printf("%s %s %s %s %s %s", message.Id, message.Producto, message.Valor, message.Tienda, message.Destino, message.Propietario)
	cont = cont + 1
	result := message.Producto + "es" + strconv.Itoa(cont)
	return &Message{Id: result}, nil
}
