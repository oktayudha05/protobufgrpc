package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "protobufgrpc/api"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedBadwordServiceServer
}

func (s *server) SayBadword(ctx context.Context, req *pb.BadwordReq) (*pb.BadwordRes, error) {
	nama := req.Nama
	jmlKata := req.JmlKata
	kalimats := "anjjj luu "
	fmt.Printf("dapet req ke badword: %s\n", req)
	if jmlKata > 1 {
		return &pb.BadwordRes{Kalimat: kalimats + nama}, nil
	}
	return  &pb.BadwordRes{Kalimat: "anjing"}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterBadwordServiceServer(s, &server{})
	fmt.Println("grpc listen on port 50052")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}