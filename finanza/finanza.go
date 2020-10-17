package main

import (
	"fmt"
	"log"
	"time"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	for{
		time.Sleep(time.Duration(10) * time.Second)
		conn, err := amqp.Dial("amqp://test:test@10.6.40.154:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()
	
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()
	
		q, err := ch.QueueDeclare(
			"hello-queue", // name
			false,         // durable
			false,         // delete when unused
			false,         // exclusive
			false,         // no-wait
			nil,           // arguments
		)
		failOnError(err, "Failed to declare a queue")
		body := "{name:arvind, message:hello}"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}

}