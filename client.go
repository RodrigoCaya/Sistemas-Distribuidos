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
	valor int
	tienda string
	destino string
	propietario int
}

func pymes(){
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
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
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
			if err != nil{
				log.Printf("error al procesar el valor: %v", err)
				continue
			}
			p.valor = i
		}

		if record[5] != ""{
			j, err := strconv.Atoi(record[5])
			if err != nil{
				log.Printf("error al procesar el propietario: %v", err)
				continue
			}
			p.propietario = j
		}

		pyme = append(pyme, p)
	}
	fmt.Println(pyme)
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := helloworld.NewHelloworldServiceClient(conn)
	
	message := helloworld.Message{
		Body: "hola jean",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
	
}
