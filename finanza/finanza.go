package main

import (
	"fmt"
	"log"
	"os"
	"io"
	// "strings"
	"strconv"
	"encoding/json"
	"encoding/csv"
	"github.com/streadway/amqp"
)

type Finan struct {
	Estado string `json:"estado"`
	Intentos  string `json:"intento"`
	Valor string `json:"valor"`
	Tipo   string `json:"tipo"`
	Id string `json:"id"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func balancetotal(){
	// Open the file
	csvfile, err := os.Open("registro.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	total := float64(0)
	for {
		// Read each record from csv
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
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1], record[2], record[3])
	}
	fmt.Printf("El balance total es %f \n",total)
}

func registrocsv(valor1 string, valor2 string, valor3 string, valor4 string){
	path := "registro.csv"
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

	// forever := make(chan bool)

	go func() {
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			var respuesta Finan

			err = json.Unmarshal([]byte(d.Body), &respuesta)

			// log.Printf("ID:%s Intentos:%s",respuesta.Id, respuesta.Intentos)

			recibidos := 0
			intentos := 0
			no_recibidos := 0
			var balance float64
			// var intentos int
			
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

			// auxbalance := strconv.FormatFloat(balance, 'E', -1, 64)
			s := fmt.Sprintf("%f", balance)
			registrocsv(respuesta.Id,respuesta.Estado,respuesta.Intentos,s)
		}
	}()
}

func main() {
	//csv
	//30%
	//input para qe termine y muestre el valor final

	go conection()
	log.Printf("Presiona 0 para ver el balance total: ")
	var first string 	  
	fmt.Scanln(&first)
	balancetotal()
	//-----------------------------------------------------

	


	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
	
}