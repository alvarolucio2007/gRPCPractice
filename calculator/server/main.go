package main

import (
	"log"
	"net"

	"github.com/alvarolucio2007/gRPCPractice/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	proto.RegisterSumServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
