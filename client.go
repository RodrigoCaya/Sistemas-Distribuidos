package main

import (
	"log"

	"os"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

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

func pym() []pymes{
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
			tienda: record[3],
			destino: record[4],
		}

		if record[2] != ""{
			i, err := strconv.Atoi(record[2])
			
			p.valor = i
		}

		if record[5] != ""{
			j, err := strconv.Atoi(record[5])
			
			p.propietario = j
		}

		pyme = append(pyme, p)

	return(pyme)
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	//leer pymes
	
	msj := pym()

	c := helloworld.NewHelloworldServiceClient(conn)
	
	message := helloworld.Message{
		id: msj.id,
		producto: msj.producto,
		valor: msj.valor,
		tienda: msj.tienda,
		destino: msj.destino,
		propietario: msj.propietario,
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)

	
}
