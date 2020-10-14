package main
import(
        "log"
        "net"
        "google.golang.org/grpc"
        "github.com/RodrigoCaya/Sistemas-Distribuidos/helloworld"
)

func main(){
        lis, err := net.Listen("tcp", ":9001")
        if err != nil {
                log.Fatalf("Failed to listen on port 9001: %v", err)
        }
        s := helloworld.Server{}
        grpcServer := grpc.NewServer()
        helloworld.RegisterHelloworldServiceServer(grpcServer, &s)
        if err := grpcServer.Serve(lis); err != nil {
                log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
        }
}
