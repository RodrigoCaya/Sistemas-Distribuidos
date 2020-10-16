package main

import (
	"fmt"
	"log"
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

func conectar(i int, c helloworld.HelloworldServiceClient){
	message := helloworld.PaqueteRequest{
		Idcamion: camiones[i].id,
		Tipo: camiones[i].tipo,
	}
	log.Printf("Camión "+camiones[i].id+" se está preparando")
		response, err := c.EnviarPaquete(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling EnviarPaquete: %s", err)
		}
		log.Printf("El camión %s tomó el paquete %s",camiones[i].id ,response.Idpaquete)
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

	log.Printf("Los camiones estan listos para partir")

	go conectar(1,c)

	go conectar(2,c)

	conectar(0,c)

	

}

