package main

import (
	"desktop/GRPC_PROJECT/Desktop/grpc-project/greet/greetpb"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	// Listener....
	connection,err := grpc.Dial("localhost:50051",grpc.WithInsecure());
	defer connection.Close();

	if err != nil {
		fmt.Printf("could not connect %v", err);
	}
	// GRPC client server..
	grpcClientServer := greetpb.NewGreetServiceClient(connection);

	fmt.Printf("Created client: ", grpcClientServer);
}
