package main 
import(
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
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

func main(){
	conexioncl()
	conexionca()
	
}
