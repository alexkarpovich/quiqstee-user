package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "github.com/alexkarpovich/micro-svc/service"
)

const (
	selfAddress = "0.0.0.0:50052";
	telegramAddress = "telegram:50053"
)

type server struct{}

func (s *server) SnipIt(ctx context.Context, in *pb.SnipRequest) (*pb.SnipResponse, error) {
	log.Printf("Received: %v", in.Url)
	return &pb.SnipResponse{Url: "Hello " + in.Url}, nil
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", selfAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUrlSnipServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGrpcClient() {
	time.Sleep(10 * time.Second)
	conn, err := grpc.Dial(telegramAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUrlSnipServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SnipIt(ctx, &pb.SnipRequest{Url: "go-client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Url)
}

func main() {
	go startGrpcServer();
	go startGrpcClient();

	for {}
}
