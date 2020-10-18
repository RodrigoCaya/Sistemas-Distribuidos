package main

import (
	"log"
	"time"
	"strconv"
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"context"
		
	"google.golang.org/grpc"

	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

//Recoge los valores de pymes.csv y envia los datos al servidor
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
		//lee los datos
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		c := helloworld.NewHelloworldServiceClient(conn)

		//crea un mensaje
		message := helloworld.Message{}
		if record[5] == "1"{
			message = helloworld.Message{
				Id: record[0],
				Producto: record[1],
				Valor: record[2],
				Tienda: record[3],
				Destino: record[4],
				Estado: "En bodega",
				Prioritario: record[5],
				Tipo: "prioritario",
			}
		}else{
			message = helloworld.Message{
				Id: record[0],
				Producto: record[1],
				Valor: record[2],
				Tienda: record[3],
				Destino: record[4],
				Estado: "En bodega",
				Prioritario: record[5],
				Tipo: "normal",
			}
		}

		//envia los datos a la funcion SayHello del servidor
		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("%s", response.Id)
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
}

//Recoge los valores de retail.csv y envia los datos al servidor
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
		//lee los datos
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
			Estado: "En bodega",
			Tipo: "retail",
		}

		//envia los datos a la funcion SayHello del servidor
		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("%s", response.Id)
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
}

//Manda un codigo de seguimiento al servidor para saber su estado
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

	var first string
	var second string 
	var tiempo1 string
	var tiempo2 string
	var thrd string 
	for{
		fmt.Println("-----------------")
		fmt.Println("Escoge: ") 
		fmt.Println("(1) Enviar pedidos") 
		fmt.Println("(2) Buscar pedido")
		fmt.Println("(0) Salir")
		fmt.Println("-----------------")
		 	  
		fmt.Scanln(&first)
		if first == "1"{
			fmt.Println("Escoge: ") 
			fmt.Println("(1) Pymes") 
			fmt.Println("(2) Retail")
			fmt.Println("-----------------")
				  
			fmt.Scanln(&second)
			if second == "1"{
				fmt.Println("Indique el intervalo de tiempo entre pedidos: ")
				 	  
				fmt.Scanln(&tiempo1)
				i, err := strconv.Atoi(tiempo1)
				if err != nil{
					log.Printf("error al ingresar el valor: %v", err)
					continue
				}
				go pym(conn,i)
				
			}
			if second == "2"{
				fmt.Println("Indique el intervalo de tiempo entre pedidos: ")
				 	  
				fmt.Scanln(&tiempo2)
				i, err := strconv.Atoi(tiempo2)
				if err != nil{
					log.Printf("error al ingresar el valor: %v", err)
					continue
				}
				go ret(conn,i)
				
			}
		}
		if first == "2"{
			fmt.Println("Introduzca el código de seguimiento: ")
				  
			fmt.Scanln(&thrd)
			codigo(conn,thrd)
			
		}
		if first == "0"{
			break
		}
	} 	
}
