package main

import (
	"log"

	"os"
	"encoding/csv"

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

func pym() pymes{
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
	var pyme []pymes
	//for
		record, err := r.Read()

		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		p := pymes{
			id: record[0],
			producto: record[1],
			valor: record[2],
			tienda: record[3],
			destino: record[4],
			propietario: record[5],
		}

		pyme = append(pyme, p)

	return(p)
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	//leer pymes
	
	for{
		msj := pym()

		c := helloworld.NewHelloworldServiceClient(conn)
		
		message := helloworld.Message{
			Id: msj.id,
			Producto: msj.producto,
			Valor: msj.valor,
			Tienda: msj.tienda,
			Destino: msj.destino,
			Propietario: msj.propietario,
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: %s", response.Id)
	}
}
