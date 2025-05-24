package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "test/proto" // test is package name in go.mod, proto is folder

	"google.golang.org/grpc"
)

// server is used to implement the Greeter service.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements Greeter.SayHello
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
