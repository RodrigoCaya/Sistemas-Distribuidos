package main

import (
	"fmt"
	"log"
	"time"
	"strconv"
	"math/rand"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

type Camion struct{
	id string
	tipo string
	pack []*helloworld.PaqueteRequest
	disponibilidad int
}

var camiones []Camion

func delivery(carga Camion){
	//todo
	// cambiar disponibilidad a 1 al volver
	// hacer +1 al espacio cuando entregue
	// reintentar cuesta 10 dignipesos
	if carga.pack[0].Valor > carga.pack[1].Valor{
		probabilidad := rand.Intn(100)
		log.Printf("%d",probabilidad)
		log.Printf("El camión %s fue a %s a enviar el paquete %s",carga.id,carga.pack[0].Destino,carga.pack[0].Idpaquete)
	}
}

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
			camiones[i].pack = append(camiones[i].pack, response)
			// camiones[i].espacio = camiones[i].espacio - 1
		}else{
			log.Printf("No hay más paquetes para el camión %s",camiones[i].id)
			if len(camiones[i].pack) == 1{
				log.Printf("El camión %s fue a hacer su entrega",camiones[i].id)
				camiones[i].disponibilidad = 0
				// delivery()
			}else if len(camiones[i].pack) == 0 {
				log.Printf("El camión %s se ha quedado vacío",camiones[i].id)
			}
		}
		if len(camiones[i].pack) == 2 {
			log.Printf("El camión %s fue a hacer su entrega",camiones[i].id)
			camiones[i].disponibilidad = 0
			delivery(camiones[i])
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
		pack: nil,
		disponibilidad: 1,
	}
	
	camionr2 := Camion{
		id: "r2",
		tipo: "retail",
		pack: nil,
		disponibilidad: 1,
	}

	camion := Camion{
		id: "n1",
		tipo: "normal",
		pack: nil,
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

