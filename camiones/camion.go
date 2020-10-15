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

	i := 0
	for{
		if i == 3{
			break
		}
		message := helloworld.PaqueteRequest{
			Idcamion: camiones[i].id,
			Tipo: camiones[i].tipo,
		}

		log.Printf("Camión "+camiones[i].id+" se está preparando")
		response, err := c.EnviarPaquete(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling EnviarPaquete: %s", err)
		}
		log.Printf("%s", response.Idpaquete)
		i = i+1
	}

	// string idpaquete = 1;
	// string idcamion = 2;
	// string seguimiento = 3;
	// string tipo = 4;
	// string valor = 5;
	// string intentos = 6;
	// string estado = 7;
	// string producto = 8;

	

}

