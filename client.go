package main

import (
	"log"
	"time"
	"os"
	"encoding/csv"
	"io"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

type pymes struct {
	id string
	producto string
	valor string
	tienda string
	destino string
	propietario string
}


func pym(conn *grpc.ClientConn){
	//leer pymes
	
	f, err := os.Open("pymes/pymes.csv")
	if err != nil{
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.FieldsPerRecord = 6

	if _, err := r.Read(); err != nil{
		panic(err)
	}

	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		// msj := pymes{
		// 	id: record[0],
		// 	producto: record[1],
		// 	valor: record[2],
		// 	tienda: record[3],
		// 	destino: record[4],
		// 	propietario: record[5],
		// }

		c := helloworld.NewHelloworldServiceClient(conn)
		
		message := helloworld.Message{
			Id: record[0],
			Producto: record[1],
			Valor: record[2],
			Tienda: record[3],
			Destino: record[4],
			Propietario: record[5],
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: %s", response.Id)
		time.Sleep(2 * time.Second)
	}
}


func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	pym(conn)
	
}
