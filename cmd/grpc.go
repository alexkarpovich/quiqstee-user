package main

import (
	"os"
	"context"
	"log"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "github.com/alexkarpovich/quiqstee-user/service"
)

type server struct{}

func (s *server) SnipIt(ctx context.Context, in *pb.SnipRequest) (*pb.SnipResponse, error) {
	log.Printf("Received: %v", in.Url)
	return &pb.SnipResponse{Url: "Hello " + in.Url}, nil
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUrlSnipServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
