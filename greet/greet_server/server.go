package main

import (
	"context"
	"desktop/GRPC_PROJECT/Desktop/grpc-project/greet/greetpb"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	fmt.Printf("Greet function is implemented with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	r := "Hello " + firstName
	res := &greetpb.GreetingResponse{
		Result: r,
	}
	return res, nil
}
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes invoked many times.....")
	firstName := req.GetGreeting().GetFirstName()

	for i := 1; i < 10; i++ {
		temp := "Hello " + firstName + " " + " " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: temp,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}
func main() {
	fmt.Println("server")

	// listner.......
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	// creating a new grpc server..
	grpcServer := grpc.NewServer()
	// registering the new service...
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	// binding the port with the grpc......
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
