package main

import (
	"fmt"
	"os"
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
	restriccion string
}

var camiones []Camion

func aux_estado_camiones(num int){
	l01 := "Carga 1: Vacía"
	l02 := "Producto: Vacío"
	l03 := "Intentos: 0"
	l04 := "Carga 2: Vacía"
	l05 := "Producto: Vacío"
	l06 := "Intentos: 0"
	if len(camiones[num].pack) == 1 {
		l01 = "Carga 1: "+camiones[num].pack[0].Estado
		l02 = "Producto: "+camiones[num].pack[0].Producto
		l03 = "Intentos: "+strconv.Itoa(int(camiones[num].pack[0].Intentos))
		}else if len(camiones[num].pack) == 2 {
			l01 = "Carga 1: "+camiones[num].pack[0].Estado
			l02 = "Producto: "+camiones[num].pack[0].Producto
			l03 = "Intentos: "+strconv.Itoa(int(camiones[num].pack[0].Intentos))
			l04 = "Carga 2: "+camiones[num].pack[1].Estado
			l05 = "Producto: "+camiones[num].pack[1].Producto
			l06 = "Intentos: "+strconv.Itoa(int(camiones[num].pack[1].Intentos))
		}
	fmt.Println(l01)
	fmt.Println(l02)
	fmt.Println(l03)
	fmt.Println(l04)
	fmt.Println(l05)
	fmt.Println(l06)
}

func estado_camiones(){
	inicio_final := "*******************************"
	separacion := "-------------------------------"
	l0 := "CAMION R1"
	l1 := "CAMION R2"
	l2 := "CAMION N1"
	fmt.Println(inicio_final)
	fmt.Println(l0)
	aux_estado_camiones(0)
	fmt.Println(separacion)
	fmt.Println(l1)
	aux_estado_camiones(1)
	fmt.Println(separacion)
	fmt.Println(l2)
	aux_estado_camiones(2)
	fmt.Println(inicio_final)
}

func reporte(carga Camion){ //falta agregar la hora de entrega
	intent := ""
	nombre := "registro"+carga.id+".csv"
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	i := 0
	v := len(carga.pack)
	for{
		if i >= v {
			break
		}
		intent = strconv.Itoa(int(carga.pack[i].Intentos))
		
		_, err = f.Write([]byte(carga.pack[i].Idpaquete+","+carga.pack[i].Tipo+","+carga.pack[i].Valor+","+carga.pack[i].Origen+","+carga.pack[i].Destino+","+intent+","+carga.pack[i].Tiempo+"\n"))
		if err != nil {
			log.Fatal(err)
		}
		i = i+1
	}
	f.Close()
}

func reportarse(carga Camion, c helloworld.HelloworldServiceClient)(nueva_carga Camion){
	reporte(carga)
	if carga.pack[0].Tipo == "retail" {
		carga.restriccion = "0"
	}else if len(carga.pack) == 2 {
		if carga.pack[1].Tipo == "retail"{
			carga.restriccion = "0"
		}
	}
	for{
		if len(carga.pack) == 0 {
			break
		}
		message := helloworld.PaqueteRequest{
			Idpaquete: carga.pack[0].Idpaquete,
			Idcamion: carga.pack[0].Idcamion,
			Seguimiento: carga.pack[0].Seguimiento,
			Tipo: carga.pack[0].Tipo,
			Valor: carga.pack[0].Valor,
			Intentos: carga.pack[0].Intentos,
			Estado: carga.pack[0].Estado,
			Producto: carga.pack[0].Producto,
			Origen: carga.pack[0].Origen,
			Destino: carga.pack[0].Destino,
			Tiempo: carga.pack[0].Tiempo,
		}
		response, err := c.EnviarDatos(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling EnviarPaquete: %s", err)
		}
		if response.Code == "ok"{
		}
		_, carga.pack = carga.pack[0], carga.pack[1:] //pop
	}
	nueva_carga = carga
	return
}

func vale_la_pena(pakete *helloworld.PaqueteRequest)(res int){
	if pakete.Tipo == "retail" {
		res = 1
		return
	}else{
		valor, err := strconv.Atoi(pakete.Valor)
		if err != nil{
			log.Printf("El valor del paquete no es un número: %v", err)
		}
		if int32(valor) > (pakete.Intentos+1)*10{
			res = 1
			return
		}else{
			res = 0
			return
		}
	}
}

func tirar_dados(carga Camion, i int)(nueva_carga Camion){
	//estado_camiones()
	//log.Printf("El camión %s fue a %s a enviar el paquete %s",carga.id,carga.pack[i].Destino,carga.pack[i].Idpaquete)
	carga.pack[i].Intentos = carga.pack[i].Intentos + 1
	probabilidad := rand.Intn(100)
	if probabilidad < 20 {
		//log.Printf("No se encontraba nadie en el domicilio %s :(",carga.pack[i].Destino)
		if vale_la_pena(carga.pack[i]) == 1{
			if carga.pack[i].Intentos == 3{
				//log.Printf("El paquete %s se cambia a No recibido",carga.pack[i].Idpaquete)
				carga.pack[i].Estado = "No Recibido"
				carga.pack[i].Tiempo = "0"
			}
		}else{
			//log.Printf("El paquete %s se cambia a No recibido",carga.pack[i].Idpaquete)
			carga.pack[i].Estado = "No Recibido"
			carga.pack[i].Tiempo = "0"
		}
		nueva_carga = carga
		return
	}else{
		//log.Printf("Se entregó el paquete %s en el domicilio %s :D",carga.pack[i].Idpaquete,carga.pack[i].Destino)
		carga.pack[i].Estado = "Recibido"
		tiempo := time.Now()
		carga.pack[i].Tiempo = tiempo.Format("2006-01-02 15:04:05")
		nueva_carga = carga
		return
	}
	nueva_carga = carga
	return
}

func delivery(carga Camion, t_envio int)(nueva_carga Camion){
	time.Sleep(time.Duration(t_envio) * time.Second)
	for{
		valor0, err := strconv.Atoi(carga.pack[0].Valor)
		if err != nil{
			log.Printf("El valor del paquete no es un número: %v", err)
		}
		valor1, err := strconv.Atoi(carga.pack[1].Valor)
		if err != nil{
			log.Printf("El valor del paquete no es un número: %v", err)
		}
		if carga.pack[0].Estado == "En camino" && carga.pack[1].Estado == "En camino"{
			if carga.pack[0].Intentos == carga.pack[1].Intentos{
				if valor0 > valor1{
					carga = tirar_dados(carga,0)
				}else{
					carga = tirar_dados(carga,1)
				}
			}else{
				if carga.pack[0].Intentos < carga.pack[1].Intentos{
					carga = tirar_dados(carga,0)
				}else{
					carga = tirar_dados(carga,1)
				}
			}
		}else{ //si ya se entregó minimo 1
			if carga.pack[0].Estado != "En camino" && carga.pack[1].Estado != "En camino"{
				break
			}else{
				if carga.pack[0].Estado == "En camino"{
					carga = tirar_dados(carga,0)
				}else{
					carga = tirar_dados(carga,1)
				}
			}
		}
	}
	nueva_carga = carga
	return
}

func delivery1(carga Camion, t_envio int)(nueva_carga Camion){
	time.Sleep(time.Duration(t_envio) * time.Second)
	for{
		//log.Printf("El camión %s fue a %s a enviar el paquete %s",carga.id,carga.pack[0].Destino,carga.pack[0].Idpaquete)
		carga.pack[0].Intentos = carga.pack[0].Intentos + 1
		probabilidad := rand.Intn(100)
		if probabilidad < 20 {
			//log.Printf("No se encontraba nadie en el domicilio %s :(",carga.pack[0].Destino)
			if vale_la_pena(carga.pack[0]) == 1{
				if carga.pack[0].Intentos == 3 {
					//log.Printf("El paquete %s se cambia a No recibido",carga.pack[0].Idpaquete)
					carga.pack[0].Estado = "No Recibido"
					carga.pack[0].Tiempo = "0"
					nueva_carga = carga
					return
				}
			}else{
				//log.Printf("El paquete %s se cambia a No recibido",carga.pack[0].Idpaquete)
				carga.pack[0].Estado = "No Recibido"
				carga.pack[0].Tiempo = "0"
				nueva_carga = carga
				return
			}
		}else{
			//log.Printf("Se entregó el paquete %s en el domicilio %s :D",carga.pack[0].Idpaquete,carga.pack[0].Destino)
			carga.pack[0].Estado = "Recibido"
			tiempo := time.Now()
			carga.pack[0].Tiempo = tiempo.Format("2006-01-02 15:04:05")
			nueva_carga = carga
			return
		}

	}
	nueva_carga = carga
	return
}

func conectar(i int, c helloworld.HelloworldServiceClient, tiempo int, t_envio int){
	for{
		message := helloworld.PaqueteRequest{
			Idcamion: camiones[i].id,
			Idpaquete: camiones[i].restriccion,
			Tipo: camiones[i].tipo,
		}
		response, err := c.EnviarPaquete(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling EnviarPaquete: %s", err)
		}
		camiones[i].restriccion = response.Tiempo
		if err != nil{
			log.Printf("El valor del paquete no es un número: %v", err)
		}
		if len(camiones[i].pack) == 1{
			estado_camiones()
			time.Sleep(time.Duration(tiempo) * time.Second)
		}
		if response.Idpaquete != "No hay más paquetes" {
			camiones[i].pack = append(camiones[i].pack, response)
		}else{
			if len(camiones[i].pack) == 1{
				//salir a hacer delivery
				estado_camiones()
				camiones[i] = delivery1(camiones[i],t_envio)
				estado_camiones()
				camiones[i] = reportarse(camiones[i],c)
			}
		}
		if len(camiones[i].pack) == 2 {
			//salir a hacer delivery
			estado_camiones()
			camiones[i] = delivery(camiones[i],t_envio)
			estado_camiones()
			camiones[i] = reportarse(camiones[i],c)
		}
		
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
		restriccion: "1",
	}
	
	camionr2 := Camion{
		id: "r2",
		tipo: "retail",
		pack: nil,
		restriccion: "1",
	}

	camion := Camion{
		id: "n1",
		tipo: "normal",
		pack: nil,
		restriccion: "0",
	}

	camiones = append(camiones, camionr1)
	camiones = append(camiones, camionr2)
	camiones = append(camiones, camion)

	fmt.Println("Indique el intervalo de tiempo de espera de los camiones: ") 
	var first string 	  
	fmt.Scanln(&first)
	espera_msj, err := strconv.Atoi(first)
	if err != nil{
		log.Printf("error al ingresar el valor: %v", err)
	}
	
	fmt.Println("Indique el tiempo que demora un envío: ") 
	var sec string
	fmt.Scanln(&sec)
	espera_envio, err := strconv.Atoi(sec)
	if err != nil{
		log.Printf("error al ingresar el valor: %v", err)
	}
	estado_camiones()
	go conectar(1,c,espera_msj,espera_envio)

	go conectar(2,c,espera_msj,espera_envio)

	conectar(0,c,espera_msj,espera_envio)
}

