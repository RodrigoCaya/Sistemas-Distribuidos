package main 
import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
	"github.com/streadway/amqp"
	"fmt"
	"time"
)

func conexioncl(){
	liscliente, err := net.Listen("tcp", ":9000")
	
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	
	s := helloworld.Server{}
	grpcServer := grpc.NewServer()
	helloworld.RegisterHelloworldServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(liscliente); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func conexionca(){
	liscamion, err2 := net.Listen("tcp", ":9001")
	if err2 != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err2)
	}
	s := helloworld.Server{}
	grpcServer := grpc.NewServer()
	helloworld.RegisterHelloworldServiceServer(grpcServer, &s)

	if err2 := grpcServer.Serve(liscamion); err2 != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err2)
	}
}



func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main(){
	go conexioncl()
	go conexionca()

	//conexion
	conn, err := amqp.Dial("amqp://test:test@10.6.40.154:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//canal
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//cola
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

