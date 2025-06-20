package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "protobufgrpc/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error){
	text := "hallo "
	return &pb.HelloResponse{Message: text + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	fmt.Println("grpc listen on port 50051")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}