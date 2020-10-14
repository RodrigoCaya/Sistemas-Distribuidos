package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist14:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := helloworld.NewHelloworldServiceClient(conn)
	
	message := helloworld.Message{
		Body: "hola jean",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
	
}
