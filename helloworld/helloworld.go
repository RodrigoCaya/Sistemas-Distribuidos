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
	cant_intentos string
}

var seguimientos []Seguimiento


var cont int = 0

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	cont = cont + 1
	result :=""
	codigo :=""
	if message.Tipo == "retail"{
		codigo = "1"+strconv.Itoa(cont)
		result = "Codigo de seguimiento de " + message.Producto + " es: " + codigo
	}else{
		codigo = "2"+strconv.Itoa(cont)
		result = "Codigo de seguimiento de " + message.Producto + " es: " + codigo
	}
	
	
	f, err := os.OpenFile("csv/registro.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		cant_intentos: "",
	}
	seguimientos = append(seguimientos, seguimiento1)
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
		}
	}else{
		return &CodeRequest{Code: result}, nil
	}
}
