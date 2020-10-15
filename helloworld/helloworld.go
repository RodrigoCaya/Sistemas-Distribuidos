package helloworld

import (
	"os"
	"log"
	"strconv"
	"encoding/csv"
	"io"
	"golang.org/x/net/context"
)

type Server struct{
}

var cont int = 0

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	log.Printf("%s %s %s %s %s %s %s", message.Id, message.Producto, message.Valor, message.Tienda, message.Destino, message.Propietario, message.Estado)
	cont = cont + 1
	result := "Codigo de seguimiento de " + message.Producto + " es: " + strconv.Itoa(cont)
	
	f, err := os.OpenFile("csv/registro.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
	}
	_, err = f.Write([]byte(strconv.Itoa(cont)+","+message.Id+","+message.Producto+","+message.Valor+","+message.Tienda+","+message.Destino+","+message.Propietario+","+message.Estado+"\n"))
    if err != nil {
        log.Fatal(err)
	}
	f.Close()
	return &Message{Id: result}, nil
}

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	f, err := os.Open("csv/registro.csv")
	if err != nil{
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 5

	if _, err := r.Read(); err != nil{
		panic(err)
	}
	result := "No se encontró el producto"
	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}
		if record[0] == message.Code{
			result = "El estado de su producto es: "+record[7]
			break
		}
	}
	return &CodeRequest{Code: result}, nil
}
