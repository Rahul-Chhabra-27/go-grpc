package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/greet/greetpb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	// Listener....
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		fmt.Printf("could not connect %v", err)
	}
	// GRPC client server..
	grpcClientServer := greetpb.NewGreetServiceClient(connection);
	doUnary(grpcClientServer);
}

func doUnary(grpcClientServer greetpb.GreetServiceClient) {
	request := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rahul",
			LastName:  "Chhabra",
		},
	}
	res, err := grpcClientServer.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Greet %v", res.Result)
}
