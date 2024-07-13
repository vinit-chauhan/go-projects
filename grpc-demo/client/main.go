package main

import (
	"log"

	pb "github.com/vinit-chauhan/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Name: []string{"Alice", "Bob", "Carla"},
	}

	// Unary
	// callSayHello(client)

	// Server side streaming
	callSatHelloServerSideStreaming(client, names)
}
