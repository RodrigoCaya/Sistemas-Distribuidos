package main

import (
	"fmt"
	"log"
	"time"
	"strconv"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

type Camion struct{
	id string
	tipo string
	espacio int
	disponibilidad int
}

var camiones []Camion

// func delivery(){
	//todo
	// cambiar disponibilidad a 1 al volver
	// hacer +1 al espacio cuando entregue
// }

func conectar(i int, c helloworld.HelloworldServiceClient, tiempo int){
	for{
		message := helloworld.PaqueteRequest{
			Idcamion: camiones[i].id,
			Tipo: camiones[i].tipo,
		}
		log.Printf("Camión "+camiones[i].id+" se está preparando")
		response, err := c.EnviarPaquete(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling EnviarPaquete: %s", err)
		}
		if response.Idpaquete != "No hay más paquetes" {
			log.Printf("El camión %s tomó el paquete %s con código %s",camiones[i].id ,response.Idpaquete, response.Seguimiento)
			camiones[i].espacio = camiones[i].espacio - 1
		}else{
			log.Printf("No hay más paquetes para el camión %s",camiones[i].id)
			if camiones[i].espacio == 1{
				log.Printf("El camión %s fue a hacer su entrega",camiones[i].id)
				camiones[i].disponibilidad = 0
				//delivery()
			}else if camiones[i].espacio == 2 {
				log.Printf("El camión %s se ha quedado vacío",camiones[i].id)
			}
		}
		if camiones[i].espacio == 0 {
			log.Printf("El camión %s fue a hacer su entrega",camiones[i].id)
			camiones[i].disponibilidad = 0
			//delivery()
		}
		time.Sleep(time.Duration(tiempo) * time.Second)
	}
}

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := helloworld.NewHelloworldServiceClient(conn)

	camionr1 := Camion{
		id: "r1",
		tipo: "retail",
		espacio: 2,
		disponibilidad: 1,
	}
	
	camionr2 := Camion{
		id: "r2",
		tipo: "retail",
		espacio: 2,
		disponibilidad: 1,
	}

	camion := Camion{
		id: "n1",
		tipo: "normal",
		espacio: 2,
		disponibilidad: 1,
	}

	camiones = append(camiones, camionr1)
	camiones = append(camiones, camionr2)
	camiones = append(camiones, camion)

	fmt.Println("Indique el intervalo de tiempo de espera de los camiones: ") 
	var first string 	  
	fmt.Scanln(&first)
	i, err := strconv.Atoi(first)
	if err != nil{
		log.Printf("error al ingresar el valor: %v", err)
	}
	log.Printf("Los camiones estan listos para partir")

	go conectar(1,c,i)

	go conectar(2,c,i)

	conectar(0,c,i)

	

}

