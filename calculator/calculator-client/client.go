package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/calculator/calculatorpb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	// Listener....
	connection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		fmt.Printf("could not connect %v", err)
	}
	// GRPC client server..
	grpcClientServer := calculatorpb.NewCalculatorClient(connection)
	doUnary(grpcClientServer);
}

func doUnary(grpcClientServer calculatorpb.CalculatorClient) {
	request := &calculatorpb.SumRequest{
		First: 10,
		Second: 10,
	}
	res,err := grpcClientServer.Sum(context.Background(),request);

	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Greet %v", res.Result);
}
