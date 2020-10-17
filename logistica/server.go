package main 
import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
	"github.com/streadway/amqp"
	"fmt"
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
	conexionca()
	
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

