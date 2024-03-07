package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/greet/greetpb"
	"fmt"
	"io"
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
	//doUnary(grpcClientServer);
	doServerStreaming(grpcClientServer);
}
func doServerStreaming(grpcClientServer greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a stream API");
	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rahul",
			LastName: "Chhabra",
		},
	};
	resStream,err := grpcClientServer.GreetManyTimes(context.Background(),request);
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}

	for {
		message, err := resStream.Recv();
		if err == io.EOF {
			// we have reached the end of the stream...
			break;
		}
		if err != nil {
			log.Fatalf("error while reading the stream %v",err);
		}
		
		log.Printf("Response from GreetManyTimes %v", message.GetResult());
	}
}
func doUnary(grpcClientServer greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary API");
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
