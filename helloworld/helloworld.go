package helloworld

import (
	"os"
	"log"
	"strconv"
	"golang.org/x/net/context"
)

type Server struct{
}

var cont int = 0

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	f, err := os.OpenFile("csv/registro.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
	}
	_, err = f.Write([]byte(message.Id+","+message.Producto+","+message.Valor+","+message.Tienda+","+message.Destino+","+message.Propietario+","+message.Estado))
    if err != nil {
        log.Fatal(err)
	}
	f.Close()

	log.Printf("%s %s %s %s %s %s %s", message.Id, message.Producto, message.Valor, message.Tienda, message.Destino, message.Propietario, message.Estado)
	cont = cont + 1
	result := message.Producto + " es: " + strconv.Itoa(cont)
	return &Message{Id: result}, nil
}
