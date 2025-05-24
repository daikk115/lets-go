package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "test/proto"
)

type HelloServiceImpl struct {
	pb.UnsafeHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Server received: " + args.GetValue()}
	return reply, nil
}

// grpc stream
func (p *HelloServiceImpl) Channel(stream pb.HelloService_ChannelServer) error {

	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// Define a message with type String
		fmt.Println("Channel received:", args.GetValue())
		reply := &pb.String{Value: "Channel sentback: update-" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, new(HelloServiceImpl))
	fmt.Printf("Stream server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
