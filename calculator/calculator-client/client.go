package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/calculator/calculatorpb"
	"fmt"
	"io"
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
	//doUnary(grpcClientServer);
	PrimeNumberDecompositionClientServerStreaming(grpcClientServer);
}
func PrimeNumberDecompositionClientServerStreaming(grpcClientServer calculatorpb.CalculatorClient) {
	fmt.Println("PrimeNumberDecomposition Stream");
	request := &calculatorpb.PrimeNumberDecompositionRequest {
		Number : 3238742343247324631,
	}
	resStream, err := grpcClientServer.PrimeNumberDecomposition(context.Background(),request);

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC %v", err)
	}
	for {
		divisor,err := resStream.Recv();

		if err == io.EOF {
			// we have reached to the end of the file...
			break;
		}
		if err != nil {
			log.Fatalf("Unable to get the divisor");
		}
		
		fmt.Printf("The divisors are %v \n", divisor.PrimeFactor);
	}
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
