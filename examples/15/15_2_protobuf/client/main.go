package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "test/proto" // test is package name in go.mod, proto is folder

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "Anonymous"
)

func main() {
	// Set up a connection to the server using NewClient (replaces deprecated Dial)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	c := pb.NewGreeterClient(conn)

	// Get the args[1] as client name instead of defaultName
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the RPC call
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
