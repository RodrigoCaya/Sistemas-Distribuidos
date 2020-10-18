package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"strconv"
	"encoding/json"
	"encoding/csv"
	"github.com/streadway/amqp"
)

//Estructura para el marshalling/unmarshalling de json
type Finan struct {
	Estado string `json:"estado"`
	Intentos  string `json:"intento"`
	Valor string `json:"valor"`
	Tipo   string `json:"tipo"`
	Id string `json:"id"`
}

//Muestra el error en pantalla si ocurre un fallo en la conexión con rabbitMQ

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//Calcula el balance total según el registro recibido por parte de logística

func balancetotal(){
	csvfile, err := os.Open("csv/registro.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	total := float64(0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		balance, err := strconv.ParseFloat(record[3],64)
		if err != nil {
			log.Fatal(err)
		}
		total = total + balance
	}
	fmt.Printf("El balance total es %f \n",total)
}

//Agrega datos al csv y lo crea si no existe
func registrocsv(valor1 string, valor2 string, valor3 string, valor4 string){
	path := "csv/registro.csv"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(valor1+","+valor2+","+valor3+","+valor4+"\n"))
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

//Setea la conexión con el emisor a través de RabbitMQ

func conection(){
	conn, err := amqp.Dial("amqp://test:test@10.6.40.154:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//Hace unmarshalling del json recibido y calcula el balance
	go func() {
		for d := range msgs {
			var respuesta Finan

			err = json.Unmarshal([]byte(d.Body), &respuesta)

			recibidos := 0
			intentos := 0
			no_recibidos := 0
			var balance float64
			
			valor, err := strconv.Atoi(respuesta.Valor)
			if err != nil {
				log.Fatal(err)
			}
			if respuesta.Estado == "Recibido"{
				recibidos = 1
			}else{
				no_recibidos = no_recibidos + 1
			}
			intentos, err = strconv.Atoi(respuesta.Intentos)
			if respuesta.Tipo == "prioritario" {
				balance = float64(valor*recibidos - no_recibidos*valor - 10*(intentos-1)) + (0.3)*float64(valor) //optiposting
			}else if respuesta.Tipo == "normal"{
				balance = float64(valor*recibidos - no_recibidos*valor - 10*(intentos-1))
			}else{
				balance = float64(valor)
			}
			s := fmt.Sprintf("%f", balance)
			registrocsv(respuesta.Id,respuesta.Estado,respuesta.Intentos,s)
		}
	}()
}

func main() {
	go conection()
	log.Printf("Presiona 0 para ver el balance total: ")
	var first string 	  
	fmt.Scanln(&first)
	balancetotal()
	
}