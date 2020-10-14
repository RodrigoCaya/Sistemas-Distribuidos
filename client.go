package main

import (
	"log"
	//"time"
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)


func pym(conn *grpc.ClientConn){	
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
		//time.Sleep(2 * time.Second)
	}
}

func ret(conn *grpc.ClientConn){	
	f, err := os.Open("retail/retail.csv")
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

	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		c := helloworld.NewHelloworldServiceClient(conn)
		
		message := helloworld.Message{
			Id: record[0],
			Producto: record[1],
			Valor: record[2],
			Tienda: record[3],
			Destino: record[4],
			Propietario: "",
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: %s", response.Id)
		//time.Sleep(2 * time.Second)
	}
}


func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	for{
		fmt.Println("Escoge: ") 
		fmt.Println("(1) Enviar pedidos") 
		fmt.Println("(2) Buscar pedido")
		fmt.Println("-----------------")
		var first string 	  
		fmt.Scanln(&first)
		if first == "1"{
			fmt.Println("Escoge: ") 
			fmt.Println("(1) Pymes") 
			fmt.Println("(2) Retail")
			fmt.Println("-----------------")
			var second string 	  
			fmt.Scanln(&second)
			if second == "1"{
				pym(conn)
				break
			}
			if second == "2"{
				ret(conn)
				break
			}
		}
		if first == "2"{
			fmt.Println("Introduzca el código de seguimiento: ")
			var thrd string 	  
			fmt.Scanln(&thrd)
			break
		}

	} 	
}
