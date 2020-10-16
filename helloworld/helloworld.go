package helloworld

import (
	"os"
	"log"
	"strconv"
	"time"
	"golang.org/x/net/context"
)

type Server struct{
}

type Seguimiento struct{
	id_paquete string
	estado_paquete string
	id_camion string
	id_seguimiento string
	cant_intentos int32
}

var seguimientos []Seguimiento

type Paquete struct{
	id_paquete string
	id_seguimiento string
	tipo string
	valor string
	intentos int32
	estado string
	origen string
	destino string
	producto string
}

var retail []Paquete
var prioritario []Paquete
var no_prioritario []Paquete

var cont int = 0

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	cont = cont + 1
	result :=""
	codigo :=""
	if message.Tipo == "retail"{
		//codigo = "0"
		codigo = "1"+strconv.Itoa(cont)
		result = "Codigo de seguimiento de " + message.Producto + " es: " + codigo
	}else{
		codigo = "2"+strconv.Itoa(cont)
		result = "Codigo de seguimiento de " + message.Producto + " es: " + codigo
	}
	
	
	f, err := os.OpenFile("../csv/registro.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
	}
	tiempo := time.Now()
	tiempo.Format("2006-01-02 15:04:05")

	_, err = f.Write([]byte(tiempo.Format("2006-01-02 15:04:05")+","+strconv.Itoa(cont)+","+message.Tipo+","+message.Producto+","+message.Valor+","+message.Tienda+","+message.Destino+","+codigo+"\n"))
    if err != nil {
        log.Fatal(err)
	}
	f.Close()
	seguimiento1 := Seguimiento{
		id_paquete: strconv.Itoa(cont),
		estado_paquete: message.Estado,
		id_camion: "",
		id_seguimiento: codigo,
		cant_intentos: 0,
	}
	seguimientos = append(seguimientos, seguimiento1)

	//agregar paquetes a colas
	paquete1 := Paquete{
		id_paquete: strconv.Itoa(cont),
		id_seguimiento: codigo,
		tipo: message.Tipo,
		valor: message.Valor,
		intentos: 0,
		estado: message.Estado,
		origen: message.Tienda,
		destino: message.Destino,
		producto: message.Producto,
	}
	
	if paquete1.tipo == "retail"{
		retail = append(retail, paquete1)
	}else if paquete1.tipo == "normal"{
				no_prioritario = append(no_prioritario, paquete1)
	}else{
		prioritario = append(prioritario, paquete1)
	}

	return &Message{Id: result}, nil
}

func (s *Server) Buscar(ctx context.Context, message *CodeRequest) (*CodeRequest, error) {
	i := 0
	result := "No se encontró el producto"
	if seguimientos != nil{
		for{
			if seguimientos[i].id_seguimiento == message.Code{
				result = "El estado de su producto es: "+seguimientos[i].estado_paquete
				return &CodeRequest{Code: result}, nil
			}
			i = i+1
		}
	}
	return &CodeRequest{Code: result}, nil
}

func (s *Server) EnviarPaquete(ctx context.Context, message *PaqueteRequest) (*PaqueteRequest, error) {
	p := Paquete{}
	i := 0
	vacio := 0
	if message.Tipo == "retail"{ //si es camion retail
		if len(retail)!=0{
			p, retail = retail[0], retail[1:] //pop
			p.estado = "En camino"
		}else if len(prioritario)!=0{
			p, prioritario = prioritario[0], prioritario[1:] //pop
			p.estado = "En camino"
		}
		
	}else{ //si es camion normal
		if len(prioritario)!=0{ 
			p, prioritario = prioritario[0], prioritario[1:] //pop
			p.estado = "En camino"
		}else if len(no_prioritario)!=0{ 
			p, no_prioritario = no_prioritario[0], no_prioritario[1:] //pop
			p.estado = "En camino"
		}else{//si estan las 3 colas vacias
			vacio = 1
		}
	}
	if vacio == 0{
		for{
			if seguimientos[i].id_paquete == p.id_paquete{
				seguimientos[i].estado_paquete = "En camino"
				seguimientos[i].id_camion = message.Idcamion
				seguimientos[i].cant_intentos = 1
				break
			}
			i = i+1
		}
		return &PaqueteRequest{Idpaquete: p.id_paquete,Idcamion: message.Idcamion,Seguimiento: p.id_seguimiento,Tipo: p.tipo,Valor: p.valor,Intentos: p.intentos,Estado: p.estado,Producto: p.producto,Origen: p.origen,Destino: p.destino}, nil
	}else{
		return &PaqueteRequest{Idpaquete: "No hay más paquetes"}, nil
	}
	//Se puede asignar un paquete prioritario a los camiones de retail tras volver de una entrega con paquetes de retail.
}