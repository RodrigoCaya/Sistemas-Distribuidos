package main

import (
	"log"
	"time"
	"strconv"
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)


func pym(conn *grpc.ClientConn, tiempo int){	
	f, err := os.Open("csv/pymes.csv")
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

		log.Printf("%s", response.Id)
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
}

func ret(conn *grpc.ClientConn, tiempo int){	
	f, err := os.Open("csv/retail.csv")
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
			Estado: "En curso",
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("%s", response.Id)
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
}

func codigo(conn *grpc.ClientConn, codigo string){
	c := helloworld.NewHelloworldServiceClient(conn)
		
		message := helloworld.CodeRequest{
			Code: codigo,
		}

		response, err := c.Buscar(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling Buscar: %s", err)
		}

		log.Printf("%s", response.Code)
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	for{
		fmt.Println("-----------------")
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
				fmt.Println("Indique el intervalo de tiempo entre pedidos: ")
				var tiempo1 string 	  
				fmt.Scanln(&tiempo1)
				i, err := strconv.Atoi(tiempo1)
				if err != nil{
					log.Printf("error al procesar el propietario: %v", err)
					continue
				}
				pym(conn,i)
				break
			}
			if second == "2"{
				fmt.Println("Indique el intervalo de tiempo entre pedidos: ")
				var tiempo2 string 	  
				fmt.Scanln(&tiempo2)
				i, err := strconv.Atoi(tiempo2)
				if err != nil{
					log.Printf("error al procesar el propietario: %v", err)
					continue
				}
				ret(conn,i)
				break
			}
		}
		if first == "2"{
			fmt.Println("Introduzca el código de seguimiento: ")
			var thrd string 	  
			fmt.Scanln(&thrd)
			codigo(conn,thrd)
			break
		}

	} 	
}
