package main

import (
	"desktop/GRPC_PROJECT/Desktop/grpc-project/greet/greetpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("server")

	// listner.......
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	// cresating a new grpc server..
	grpcServer := grpc.NewServer()
	// registering the new service...
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	// binding the port with the grpc......
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
