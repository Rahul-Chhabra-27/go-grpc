package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/calculator/calculatorpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (*server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.Calculator_PrimeNumberDecompositionServer) error {
	fmt.Println("Recieved PrimeNumberDecomposition %v\n", request)
	//number := request.GetNumber();

	return nil
}

func (*server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	firstNumber := request.GetFirst()
	lastNumber := request.GetSecond()

	result := firstNumber + lastNumber

	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}
func main() {
	fmt.Println("calculator server")

	// listner.......
	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	// creating a new grpc server..
	grpcServer := grpc.NewServer()
	// registering the new service...

	calculatorpb.RegisterCalculatorServer(grpcServer, &server{})
	// binding the port with the grpc......
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
